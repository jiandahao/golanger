// Based on: https://github.com/apexskier/go-template-validation.git

package validator

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	textTemplate "text/template"
)

var (
	maxFixes                    = 10
	templateErrorRegex          = regexp.MustCompile(`template: (.*?):((\d+):)?(\d+): (.*)`)
	findTokenRegex              = regexp.MustCompile(`['"](.+)['"]`)
	findExprRegex               = regexp.MustCompile(`<(\..+?)>`)
	functionNotFoundRegex       = regexp.MustCompile(`function "(.+)" not defined`)
	missingValueForCommandRegex = regexp.MustCompile(`missing value for command`)
	firstEmptyCommandRegex      = regexp.MustCompile(`{{((-?\s*?)|(\s*?-?))}}`)
)

// ErrorLevel is the type of error found
type ErrorLevel string

const (
	misunderstoodError ErrorLevel = "misunderstood"
	parseErrorLevel    ErrorLevel = "parse"
	execErrorLevel     ErrorLevel = "exec"
)

// FGRed red color of text
const FGRed = 31

// BGRed red color of background
const BGRed = 41

// TemplateError describes the details of template error.
type TemplateError struct {
	Line        int
	Char        int
	Description string
	Level       ErrorLevel
}

// CountDigits returns the number of digits in a decimal number
func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

// SplitLines splits a string into lines
// Supports windows or unix line segments
func SplitLines(str string) []string {
	return strings.Split(strings.Replace(str, "\r\n", "\n", -1), "\n")
}

// CreateTemplateError parse and return a template error.
func CreateTemplateError(err error, level ErrorLevel) TemplateError {
	matches := templateErrorRegex.FindStringSubmatch(err.Error())
	if len(matches) == 6 {
		// tplName := matches[1]

		// 2 is line + : group if char is found
		// line is in pos 4, unless a char is found in which case it's 3 and char is 4

		lineIndex := 4
		char := -1
		if matches[3] != "" {
			lineIndex = 3
			char, err = strconv.Atoi(matches[4])
			if err != nil {
				char = -1
			}
		}

		line, err := strconv.Atoi(matches[lineIndex])
		if err != nil {
			line = -1
		} else {
			line = line - 1
		}

		description := matches[5]

		return TemplateError{
			Line:        line,
			Char:        char,
			Description: description,
			Level:       level,
		}
	}
	return TemplateError{
		Line:        -1,
		Char:        -1,
		Description: err.Error(),
		Level:       misunderstoodError,
	}
}

// Validate validates a template
func Validate(text string) (*textTemplate.Template, []TemplateError) {
	baseTpl := textTemplate.New("base template")
	return parseInternal(text, baseTpl, 0)
}

func parseInternal(text string, baseTpl *textTemplate.Template, depth int) (t *textTemplate.Template, tplErrs []TemplateError) {
	lines := SplitLines(text)

	if depth >= maxFixes {
		return baseTpl, tplErrs
	}

	t, err := baseTpl.Parse(text)
	if err == nil {
		return t, tplErrs
	}

	tplErrs = append(tplErrs, CreateTemplateError(err, parseErrorLevel))
	// make this mutable
	tplErr := &tplErrs[len(tplErrs)-1]

	if tplErr.Level == misunderstoodError {
		return t, tplErrs
	}

	if tplErr.Char == -1 {
		// try to find a character to line up with
		tokenLoc := findTokenRegex.FindStringIndex(tplErr.Description)
		if tokenLoc != nil {
			token := string(tplErr.Description[tokenLoc[0]+1 : tokenLoc[1]-1])
			lastChar := strings.LastIndex(lines[tplErr.Line], token)
			firstChar := strings.Index(lines[tplErr.Line], token)
			// if it's not the only match, we don't know which character is the one the error occured on
			if lastChar == firstChar {
				tplErr.Char = firstChar
			}
		}
	}

	badFunctionMatch := functionNotFoundRegex.FindStringSubmatch(tplErr.Description)
	if badFunctionMatch != nil {
		token := badFunctionMatch[1]
		t, parseTplErrs := parseInternal(text, baseTpl.Funcs(textTemplate.FuncMap{
			token: func() error {
				return nil
			},
		}), depth+1)
		return t, append(tplErrs, parseTplErrs...)
	}

	if missingValueForCommandRegex.MatchString(tplErr.Description) {
		matches := firstEmptyCommandRegex.FindStringSubmatch(text)
		if matches != nil {
			line := SplitLines(text)[tplErr.Line]
			indexes := firstEmptyCommandRegex.FindStringIndex(line)
			if indexes != nil {
				tplErr.Char = indexes[0]
			}
			replacement := fmt.Sprintf(fmt.Sprintf("%%%ds", len(matches[0])), "")
			t, parseTplErrs := parseInternal(
				strings.Replace(text, matches[0], replacement, 1),
				baseTpl,
				depth+1,
			)
			return t, append(tplErrs, parseTplErrs...)
		}
	}

	return baseTpl, tplErrs
}

// Exec applies a parsed template to the specified data object
func Exec(t *textTemplate.Template, data interface{}, buf *bytes.Buffer) []TemplateError {
	return execInternal(t, data, buf, 0)
}

func execInternal(t *textTemplate.Template, data interface{}, buf *bytes.Buffer, depth int) []TemplateError {
	tplErrs := make([]TemplateError, 0)
	err := t.Execute(buf, data)
	if err != nil {
		if err.Error() == fmt.Sprintf(`template: %s: "%s" is an incomplete or empty template`, t.Name(), t.Name()) {
			return tplErrs
		}
		tplErr := CreateTemplateError(err, execErrorLevel)
		tplErrs = append(tplErrs, tplErr)

		matches := findExprRegex.FindStringSubmatch(tplErr.Description)
		if len(matches) == 2 {
			fmt.Println(matches[1])
		}
	}
	return tplErrs
}

// PrintErrorDetails prints error details
func PrintErrorDetails(text string, errs []TemplateError) {
	fmt.Println(ErrorDetails(text, errs))
}

// ErrorDetails constructs errors as a single string which shows where validation errors are happening.
func ErrorDetails(text string, errs []TemplateError) string {
	defs := strings.Split(text, "\n")

	var details []string = make([]string, len(defs))

	copy(details, defs)

	for _, err := range errs {
		if err.Line < 0 {
			fmt.Printf("Line: %v, Char: %v, Description: %s", err.Line, err.Char, err.Description)
			continue
		}

		charAt := err.Char
		if charAt <= 0 {
			charAt = 1
		}

		var spaceCounter int
		for _, c := range details[err.Line] {
			if c != ' ' && c != '\t' {
				break
			}
			if c == '\t' {
				spaceCounter = spaceCounter + 8
			} else {
				spaceCounter++
			}
		}

		details[err.Line] = details[err.Line] + "\n" + strings.Repeat(" ", spaceCounter+charAt) + fmt.Sprintf("\033[%vm↑ %v\033[0m", FGRed, err.Description)
	}

	return strings.Join(details, "\n")
}
