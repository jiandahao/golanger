package dbgen

import (
	"bytes"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"unicode"

	"github.com/zeromicro/ddl-parser/console"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
)

var consoleLog = console.NewColorConsole()
var modelTemplateParser *template.Template
var cacheConnTemplateParser *template.Template

func init() {
	var err error
	modelTemplateParser, err = template.New("model template").Parse(dbModelTemplate)
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateFromDDL generate model CURD code from ddl.
func GenerateFromDDL(filename string, database string, outputDir string, withCache bool) error {
	if err := generateFromDDL(filename, database, outputDir, withCache); err != nil {
		consoleLog.ErrorF("failed to generate from ddl: %v", err)
		return err
	}

	consoleLog.Info("Done.")
	return nil
}

func generateFromDDL(filename string, database string, outputDir string, withCache bool) error {
	ddlFilePath, err := filepath.Abs(filename)
	if err != nil {
		return err
	}

	tables, err := parser.Parse(ddlFilePath, database)
	if err != nil {
		return err
	}

	outputDir, err = filepath.Abs(outputDir)
	if err != nil {
		return err
	}

	pkgName := filepath.Base(outputDir)
	imports := []string{
		"context",
		"fmt",
		"gorm.io/gorm",
		"github.com/pkg/errors",
		"github.com/jiandahao/golanger/pkg/storage/options",
		"github.com/jiandahao/golanger/pkg/storage/filter",
	}

	var output = make(map[string]string)
	for _, table := range tables {
		t := &Table{
			Table:   table,
			PkgName: pkgName,
			Imports: imports,
		}

		if err := t.Validate(); err != nil {
			return fmt.Errorf("table %s: %v", t.Name.Source(), err)
		}

		var buf bytes.Buffer
		if err := modelTemplateParser.Execute(&buf, t); err != nil {
			return fmt.Errorf("execute tempalte for table %s error : %v", t.Name.Source(), err)
		}

		//	fmt.Println(buf.String())
		output[fmt.Sprintf("%s_model.go", t.Name.Source())] = buf.String()

		if withCache {
			var buf bytes.Buffer
		}
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	for tab, out := range output {
		outputFilename := filepath.Join(outputDir, tab)
		_, err := os.Stat(outputFilename)
		if err == nil || os.IsExist(err) {
			consoleLog.Warning(fmt.Sprintf("%s already exists, ignored.", fmt.Sprintf("%s_model.go", tab)))
			continue
		}

		if err := WriteFormatedGoFile(outputFilename, []byte(out), os.ModePerm); err != nil {
			return err
		}

	}

	utilsOutput := filepath.Join(outputDir, "utils.go")
	_, err = os.Stat(utilsOutput)
	if os.IsExist(err) {
		return nil
	}

	if err := WriteFormatedGoFile(utilsOutput, []byte(`
		package `+pkgName+`

		import (
			"errors"

			"gorm.io/gorm"
		)

		var ErrInvalidArgument = errors.New("invalid arguments")

		// IsRecordNotFound returns true if err represents a RecordNotFound error.
		func IsRecordNotFound(err error) bool {
			return errors.Is(err, gorm.ErrRecordNotFound)
		}

		// IsInvalidArgumentError returns true if err represents a ErrInvalidArgument error.
		func IsInvalidArgumentError(err error) bool {
			return errors.Is(err, ErrInvalidArgument)
		}`), os.ModePerm); err != nil {
		return err
	}

	return nil
}

func WriteFormatedGoFile(name string, data []byte, perm fs.FileMode) error {
	formatedOuptput, err := format.Source(data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(name, formatedOuptput, os.ModePerm); err != nil {
		return err
	}

	return nil
}

type Table struct {
	*parser.Table
	PkgName string
	Imports []string
}

func (t *Table) Validate() error {
	if t.PrimaryKey.Name.Source() == "" {
		return fmt.Errorf("missing primary key")
	}

	// field validation, keywords or reserved words [https://dev.mysql.com/doc/refman/8.0/en/keywords.html]

	return nil
}

func (t *Table) TableCamleName() string {
	return t.Table.Name.ToCamel()
}

func (t *Table) TableUntitleCamleName() string {
	source := t.Table.Name.ToCamel()
	r := rune(source[0])
	return string(unicode.ToLower(r)) + source[1:]
}

var cacheConnTemplate = `
{{$tableCamelName := .TableCamleName}}
package {{.PkgName}}

import (
	"github.com/jiandahao/golanger/pkg/storage/cache"
)

type {{$tableCamelName}}CacheConn struct {
	conn {{$tableCamelName}}Model
	expire time.Duration
}

func New{{$tableCamelName}}CacheConn(m {{$tableCamelName}}Model) *{{$tableCamelName}}CacheConn {
	return &{{$tableCamelName}}CacheConn{conn: m}
}

func (c *{{$tableCamelName}}CacheConn ) Expire(t time.Duration) *{{$tableCamelName}}CacheConn {
	
}

func (c *{{$tableCamelName}}CacheConn ) Insert(ctx context.Context, data *{{$tableCamelName}}) error {

}

func (c *{{$tableCamelName}}CacheConn ) Delete(ctx context.Context, {{.PrimaryKey.Name.Source}} {{.PrimaryKey.DataType}}) error {

}

func (c *{{$tableCamelName}}CacheConn ) Update(ctx context.Context, data *{{$tableCamelName}}) error {

}

func (c *{{$tableCamelName}}CacheConn ) FindOne(ctx context.Context, {{.PrimaryKey.Name.Source}} {{.PrimaryKey.DataType}}) (*{{$tableCamelName}}, error) {

}

{{- range $key, $value := .UniqueIndex -}}
{{$indexField := index $value 0}}
func (c *{{$tableCamelName}}CacheConn ) FindOneBy{{$indexField.Name.ToCamel}}(ctx context.Context, {{$indexField.NameOriginal}} {{$indexField.DataType}}) (*{{$tableCamelName}}, error) {

}
{{- end}}
func (c *{{$tableCamelName}}CacheConn ) Query(ctx context.Context, filters string, opts ...options.QueryOption) ([]*{{$tableCamelName}}, int64, error) {

}

func (c *{{$tableCamelName}}CacheConn ) WithDB(db *gorm.DB) {{$tableCamelName}}Model {

}
`

var dbModelTemplate = `
{{$tableCamelName := .TableCamleName}}
{{$tableUntitlCameleName := .TableUntitleCamleName}}
package {{.PkgName}}

import (
	{{range .Imports -}}
	"{{.}}"
	{{end}}
)

type (
	// {{$tableCamelName}}Model is an interface that wraps the CURD method.
	{{$tableCamelName}}Model interface{
		Insert(ctx context.Context, data *{{$tableCamelName}}) error
		Delete(ctx context.Context, {{.PrimaryKey.Name.Source}} {{.PrimaryKey.DataType}}) error
		Update(ctx context.Context, data *{{$tableCamelName}}) error
		FindOne(ctx context.Context, {{.PrimaryKey.Name.Source}} {{.PrimaryKey.DataType}}) (*{{$tableCamelName}}, error)
		{{- range $key, $value := .UniqueIndex -}}
		{{$indexField := index $value 0}}
		FindOneBy{{$indexField.Name.ToCamel}}(ctx context.Context, {{$indexField.NameOriginal}} {{$indexField.DataType}}) (*{{$tableCamelName}}, error)
		{{- end}}
		Query(ctx context.Context, filters string, opts ...options.QueryOption) ([]*{{$tableCamelName}}, int64, error)
		WithDB(db *gorm.DB) {{$tableCamelName}}Model
	}

	default{{$tableCamelName}}Model struct {
		dbConn *gorm.DB
	}

	// {{$tableCamelName}} describes the table schema structure.
	{{$tableCamelName}} struct {
		{{range .Fields -}}
		{{.Name.ToCamel}} {{.DataType}} ` + "`" + `gorm:"column:{{.NameOriginal}}"` + "`" + ` {{if .Comment}} // {{.Comment}} {{end}}
		{{end}}
	}
)

// TableName returns the table name.
func ({{$tableCamelName}}) TableName() string {
	return "` + "`{{.Name.Source}}`" + `"
}

// New{{$tableCamelName}}Model creates a default{{$tableCamelName}}Model.
func New{{$tableCamelName}}Model(conn *gorm.DB) {{$tableCamelName}}Model {
	return &default{{$tableCamelName}}Model{
		dbConn: conn,
	}
}

func (m *default{{$tableCamelName}}Model) WithDB(dbConn *gorm.DB) {{$tableCamelName}}Model {
	return &default{{$tableCamelName}}Model{
		dbConn: dbConn,
	}
}

// Insert insert one record.
func (m *default{{$tableCamelName}}Model) Insert(ctx context.Context, data *{{$tableCamelName}}) error {
	err := m.dbConn.WithContext(ctx).Create(data).Error
	if err != nil {
		return errors.Wrap(err, "Insert error")
	}

	return nil
}

// Delete delete a record by primary key.
func (m *default{{$tableCamelName}}Model) Delete(ctx context.Context, {{.PrimaryKey.Name.Source}} {{.PrimaryKey.DataType}}) error {
	if err := m.dbConn.WithContext(ctx).Exec(fmt.Sprintf("DELETE FROM %s WHERE ` + "`{{.PrimaryKey.Name.Source}}`" + ` = ? LIMIT 1", {{$tableCamelName}}{}.TableName()), {{.PrimaryKey.Name.Source}}).Error; err != nil {
		return errors.Wrap(err, "Delete error")
	}

	return nil
}

// Update update a record.
func (m *default{{$tableCamelName}}Model) Update(ctx context.Context, data *{{$tableCamelName}}) error {
	err := m.dbConn.WithContext(ctx).Where("` + "`{{.PrimaryKey.Name.Source}}`" + `  = ?", data.{{.PrimaryKey.Name.ToCamel}}).Updates(data).Error
	if err != nil {
		return errors.Wrap(err, "Update error")
	}
	return nil
}

// FindOne find records by primary key.
func (m *default{{$tableCamelName}}Model) FindOne(ctx context.Context, {{.PrimaryKey.Name.Source}} {{.PrimaryKey.DataType}}) (*{{$tableCamelName}}, error) {
	var resp {{$tableCamelName}}
	err := m.dbConn.WithContext(ctx).Where("` + "`{{.PrimaryKey.Name.Source}}`" + `  = ?", {{.PrimaryKey.Name.Source}}).Limit(1).Take(&resp).Error
	if err != nil {
		return nil, errors.Wrap(err, "FindOne error")
	}

	return &resp, nil
}


{{range $key, $value := .UniqueIndex -}}
{{$indexField := index $value 0}}
// FindOneBy{{$indexField.Name.ToCamel}} find one record by unique key {{$indexField.NameOriginal}}.
func (m *default{{$tableCamelName}}Model) FindOneBy{{$indexField.Name.ToCamel}}(ctx context.Context, {{$indexField.NameOriginal}} {{$indexField.DataType}}) (*{{$tableCamelName}}, error) {
	var resp {{$tableCamelName}}
	err := m.dbConn.WithContext(ctx).Where("` + "`{{$indexField.NameOriginal}}`" + `  = ?", {{$indexField.NameOriginal}}).Limit(1).Take(&resp).Error
	if err != nil {
		return nil, errors.Wrap(err, "FindOneBy{{$indexField.Name.ToCamel}} error")
	}

	return &resp, nil
}
{{end}}


const (
	{{range .Fields -}}
	{{$tableCamelName}}Column_{{.Name.ToCamel}} string = "{{.NameOriginal}}"
	{{end}}
)

// define all permitted query conditions
var {{$tableUntitlCameleName}}QueryFilter = filter.NewParser(map[filter.FieldNameType][]filter.Operator{
	"{{.PrimaryKey.Name.Source}}": {filter.Equal, filter.In},
	{{range $key, $value := .UniqueIndex -}}
	{{$indexField := index $value 0}}"{{$indexField.NameOriginal}}": {filter.Equal, filter.In},
	{{end}}
})

// Query query records by filters.
func (m *default{{$tableCamelName}}Model) Query(ctx context.Context, filters string, opts ...options.QueryOption) ([]*{{$tableCamelName}}, int64, error) {
	dbClient := m.dbConn.WithContext(ctx)
	if filters != "" {
		conds, args, err := {{$tableUntitlCameleName}}QueryFilter.Parse(filters)
		if err != nil {
			return nil, 0, errors.Wrap(ErrInvalidArgument, err.Error())
		}

		dbClient = dbClient.Where(conds, args...)
	}

	var counter int64
	if err := dbClient.Model(&{{$tableCamelName}}{}).Count(&counter).Error; err != nil {
		return nil, 0, errors.Wrap(err, "Query:Count error")
	}

	for _, opt := range opts {
		dbClient = opt(dbClient)
	}

	var records []*{{$tableCamelName}}
	if err := dbClient.Find(&records).Error; err != nil {
		return nil, 0, errors.Wrap(err, "Query:Find error")
	}

	return records, counter, nil
}
`
