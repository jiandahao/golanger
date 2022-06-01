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
func GenerateFromDDL(ddlFilePath string, moduleName string, outpDir string, withCache bool) error {
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

	if err := generetor.StartFromDDL(ddlFilePath, withCache, moduleName); err != nil {
		return err
	}

	return nil
}

func initTemplates() {
	template.Field = `{{.name}} {{.type}} {{.tag}} {{if .hasComment}}// {{.comment}}{{end}}`
	template.Model = `// Code generated.
	
	package {{.pkg}}
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
				{{if .withCache}} cachedConn cache.CachedConn {{end}}
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
	func New{{.upperStartCamelObject}}Model(conn *gorm.DB {{if .withCache}}, cacheConn cache.CachedConn{{end}}) {{.upperStartCamelObject}}Model {
		return &default{{.upperStartCamelObject}}Model{
			dbConn: conn,
			{{- if .withCache}} 
			cachedConn: cacheConn, 
			{{- end }}
		}
	}
	`

	template.Error = `package {{.pkg}}`

	template.InsertMethod = `Insert(ctx context.Context, data *{{.upperStartCamelObject}}) error`
	template.Insert = `
	// Insert insert one record into user_tab.
	func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) error {
		err := dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.WithContext(ctx).Create(&data).Error
		})

		if err != nil {
			return err
		}

		return nil
	}
	`

	template.DeleteMethod = `Delete(ctx context.Context,{{.lowerStartCamelPrimaryKey}} {{.dataType}}) error`
	template.Delete = `
	// Delete deletes by primary key.
	func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
		{{if .withCache}}
			{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
			if err!=nil{
				return err
			}{{end}}
	
			{{.keys}}

			keys := []string{ {{.keyValues}} }
			return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
				return tx.WithContext(ctx).Exec(fmt.Sprintf("DELETE FROM %s WHERE {{.originalPrimaryKey}} = ? LIMIT 1", {{.upperStartCamelObject}}{}.TableName()), {{.lowerStartCamelPrimaryKey}}).Error
			}, func()  {
				m.cachedConn.DelCache(keys...)
			})
		{{- else -}}
		return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.WithContext(ctx).Exec(fmt.Sprintf("DELETE FROM %s WHERE {{.originalPrimaryKey}} = ? LIMIT 1", {{.upperStartCamelObject}}{}.TableName()), {{.lowerStartCamelPrimaryKey}}).Error
		})
		{{- end}}
	}`

	template.UpdateMethod = `Update(ctx context.Context, data *{{.upperStartCamelObject}}) error `
	template.Update = `
	// Update update a record.
	func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, data *{{.upperStartCamelObject}}) error {
		{{if .withCache}}{{- .keys}}
		keys := []string{ {{.keyValues}}}

		return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.WithContext(ctx).Updates(data).Error
		}, func()  {
			 m.cachedConn.DelCache(keys...)
		})
		{{- else -}}
		return dbutils.Transaction(ctx, m.dbConn, func(ctx context.Context, tx *gorm.DB) error {
			return tx.WithContext(ctx).Updates(data).Error
		})
		{{- end}}
	}`

	template.FindOneMethod = `FindOne(ctx context.Context,{{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error)`
	template.FindOne = `
	// FindOne find records by primary key.
	func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
		var resp {{.upperStartCamelObject}}
		{{if .withCache}}{{.cacheKey}}
		err := m.cachedConn.QueryRow(&resp, func(v interface{}) error {
			return m.dbConn.WithContext(ctx).Where("{{.originalPrimaryKey}}  = ?", id).Limit(1).Take(v).Error
		}, {{.cacheKeyVariable}})
		{{else}}
		err := m.dbConn.WithContext(ctx).Where("{{.originalPrimaryKey}} = ?", id).Limit(1).Take(&resp).Error
		{{end}}
		if err != nil {
			return nil, err
		}

		return &resp, nil
	}`

	template.FindOneByFieldMethod = `FindOneBy{{.upperField}}(ctx context.Context,{{.in}}) (*{{.upperStartCamelObject}}, error)`
	template.FindOneByField = `
	// FindOneBy{{.upperField}} finds records by {{.upperField}}.
	func (m *default{{.upperStartCamelObject}}Model) FindOneBy{{.upperField}}(ctx context.Context, {{.in}}) (*{{.upperStartCamelObject}}, error) {
		var resp {{.upperStartCamelObject}}
		{{if .withCache}}{{.cacheKey}}
		err := m.cachedConn.QueryRow(&resp, func(v interface{}) error {
			return m.dbConn.WithContext(ctx).Where("{{.originalField}}", {{.lowerStartCamelField}}).Limit(1).Take(v).Error
		}, {{.cacheKeyVariable}})
		{{else}}
		err := m.dbConn.Where("{{.originalField}}", {{.lowerStartCamelField}}).Limit(1).Take(&resp).Error
		{{end}}
		if err != nil {
			return nil, err
		}

		return &resp, nil
	}`

	template.Imports = `
	import (
		"context"
		"fmt"
		"database/sql"
	
		"github.com/jiandahao/golanger/pkg/storage/cache"
		"github.com/jiandahao/golanger/pkg/storage/dbutils"
		"gorm.io/gorm"
	)
	
	var _ = sql.NullString{}
	`

	template.ImportsNoCache = `
	import (
		"context"
		"fmt"
		"database/sql"
	
		"github.com/jiandahao/golanger/pkg/storage/dbutils"
		"gorm.io/gorm"
	)
	var _ = sql.NullString{}
	`

	template.FindOneByFieldExtraMethod = ``

	template.Vars = `
	{{if .withCache}}
	var (
		{{.cacheKeys}}
	)
	{{end}}`
}
