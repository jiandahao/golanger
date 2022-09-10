package gingen

import (
	"testing"

	tplValidator "github.com/jiandahao/golanger/pkg/template/validator"
)

func TestDefaultClientTempl(t *testing.T) {
	_, errs := tplValidator.Validate(defaultClientTempl)

	if len(errs) > 0 {
		tplValidator.PrintErrorDetails(defaultClientTempl, errs)
		t.Fatalf("invalid template")
	}
}
