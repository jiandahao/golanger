package dbgen

import (
	"path/filepath"
	"sync"

	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/gen"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
)

var once sync.Once

// GenerateFromDDL generate database CURD code from DDL.
func GenerateFromDDL(ddlFilePath string, outpDir string) error {
	once.Do(func() {
		initTemplates()
	})

	var err error
	ddlFilePath, err = filepath.Abs(ddlFilePath)
	if err != nil {
		return err
	}

	generetor, err := gen.NewDefaultGenerator(outpDir, &config.Config{NamingFormat: "go_zero"})
	if err != nil {
		return err
	}

	if err := generetor.StartFromDDL(ddlFilePath, false, "database"); err != nil {
		return err
	}

	return nil
}

func initTemplates() {
	template.Field = `{{.name}} {{.type}} {{.tag}} {{if .hasComment}}// {{.comment}}{{end}}`
	template.Model = `package {{.pkg}}
		{{.imports}}
		{{.vars}}
		{{.types}}
		{{.new}}
		{{.insert}}
		{{.find}}
		{{.update}}
		{{.delete}}
		{{.extraMethod}}`

	template.Tag = "`gorm:\"column:{{.field}}\"`"

	// Types defines a template for types in model
	template.Types = `
		type (
			// {{.upperStartCamelObject}}Model is an interface that wraps the CURD method.
			{{.upperStartCamelObject}}Model interface{
				{{.method}}
			}

			default{{.upperStartCamelObject}}Model struct {
				dbConn *gorm.DB
			}

			// {{.upperStartCamelObject}} describes the table schema structure.
			{{.upperStartCamelObject}} struct {
				{{.fields}}
			}
		)`

	template.New = `
	// TableName returns the table name.
	func ({{.upperStartCamelObject}}) TableName() string {
		return {{.table}}
	}

	// New{{.upperStartCamelObject}}Model creates a default{{.upperStartCamelObject}}Model.
	func New{{.upperStartCamelObject}}Model(conn *gorm.DB) {{.upperStartCamelObject}}Model {
		return &default{{.upperStartCamelObject}}Model{
			dbConn: conn,
			// table:      {{.table}},
		}
	}
	`

	template.Error = `package {{.pkg}}`

	template.InsertMethod = `Insert(ctx context.Context, data *{{.upperStartCamelObject}}) error`
	template.Insert = `
	// Insert insert one record into user_tab.
	func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) error {
		err := dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.Create(&data).Error
		})

		if err != nil {
			return err
		}

		return nil
	}
	`

	template.DeleteMethod = ` Delete(ctx context.Context,{{.lowerStartCamelPrimaryKey}} {{.dataType}}) error`
	template.Delete = `
	// Delete deletes by primary key.
	func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
		return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.Exec(fmt.Sprintf("DELETE FROM %s WHERE {{.originalPrimaryKey}} = ? LIMIT 1", {{.upperStartCamelObject}}{}.TableName()), {{.lowerStartCamelPrimaryKey}}).Error
		})
	}`

	template.UpdateMethod = `Update(ctx context.Context, data *{{.upperStartCamelObject}}) error `
	template.Update = `
	// Update update a record.
	func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, data *{{.upperStartCamelObject}}) error {
		return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.Updates(data).Error
		})
	}`

	template.FindOneMethod = `FindOne(ctx context.Context,{{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)`
	template.FindOne = `
	// FindOne find records by primary key.
	func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
		var resp {{.upperStartCamelObject}}
		err := m.dbConn.Where("{{.originalPrimaryKey}} = ?", id).Limit(1).Find(&resp).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	
		if err != nil {
			return nil, err
		}
	
		return &resp, nil
	}`

	template.FindOneByFieldMethod = `FindOneBy{{.upperField}}(ctx context.Context,{{.in}}) (*{{.upperStartCamelObject}}, error)`
	template.FindOneByField = `
	// FindOneBy find records by {{.upperField}}.
	func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}(ctx context.Context, {{.in}}) (*{{.upperStartCamelObject}}, error) {
		var resp {{.upperStartCamelObject}}
		err := m.dbConn.Where("{{.originalField}}", {{.lowerStartCamelField}}).Limit(1).Take(&resp).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		if err != nil {
			return nil, err
		}
	
		return &resp, nil
	}`

	template.Imports = ``
	template.ImportsNoCache = `
	import (
		"context"
		"errors"
		"fmt"
		"database/sql"
	
		dbutils "github.com/jiandahao/golanger/pkg/storage/db"
		"gorm.io/gorm"
	)`

	template.FindOneByFieldExtraMethod = ``

	template.Vars = ``
}
