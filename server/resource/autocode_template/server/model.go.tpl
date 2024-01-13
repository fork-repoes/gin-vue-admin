// 自动生成模板{{.StructName}}
package {{.Package}}

import (
	"github.com/pkg/errors"
	"github.com/wordpress-plus/server-core/global"
	"github.com/wordpress-plus/server-core/utils"
	"gorm.io/gorm"
	{{ if .HasTimer }}"time"{{ end }}
	{{ if .NeedJSON }}"gorm.io/datatypes"{{ end }}
)

// {{.Description}} 结构体  {{.StructName}}
type {{.StructName}} struct {
      global.GVA_MODEL {{- range .Fields}}
            {{- if eq .FieldType "enum" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};type:enum({{.DataTypeLong}});comment:{{.Comment}};" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "picture" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "video" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
             {{- else if eq .FieldType "file" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "pictures" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "richtext" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}type:text;" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- end }}  {{ if .FieldDesc }}//{{.FieldDesc}} {{ end }} {{- end }}
      {{- if .AutoCreateResource }}
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
      {{- end}}
}

{{ if .TableName }}
// TableName {{.Description}} {{.StructName}}自定义表名 {{.TableName}}
func (*{{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}


// Create 创建{{.StructName}}表记录(存在主键则更新否则创建)
func (c *{{.StructName}}) Create() (err error) {
	return global.DB.Create(c).Error
}

// Updates 更新{{.StructName}}记录(必须存在主键否则失败): 零值不更新
func (c *{{.StructName}}) Updates() (err error) {

	return global.DB.Updates(c).Error
}

// Delete 根据PK删除{{.StructName}}表记录
func (c *{{.StructName}}) Delete() (err error) {

	return global.DB.Delete(c).Error
}

// Deletes 根据PK删除{{.StructName}}表记录(存在主键或只有一条则删除否则失败)
func (c *{{.StructName}}) Deletes() error {

	if utils.ExistUniqueCon(*c) {
		return global.DB.Where(c).Delete(&{{.StructName}}{}).Error
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		count, err := c.Count()
		if err != nil {
			return err
		}

		if count > 1 {
			return errors.New("存在多条记录，请使用唯一标识删除")
		}

		return global.DB.Where(c).Delete(&{{.StructName}}{}).Error
	})
}

// DeleteAll 根据条件删除所有符合{{.StructName}}表记录
func (c *{{.StructName}}) DeleteAll() (err error) {

	return global.DB.Where(c).Delete(&{{.StructName}}{}).Error
}

// Query 根据id获取{{.StructName}}表记录
func (c *{{.StructName}}) Query() (ps []*{{.StructName}}, err error) {
	result := global.DB.Where(c).Find(&ps)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return ps, nil
}

// FirstById 根据id获取{{.StructName}}表记录
func (c *{{.StructName}}) FirstById(id uint) (err error) {
	// @Model({{.StructName}})
	return global.DB.Where("id = ?", id).First(&c).Error
}

// Deprecated: Use First instead.
// 根据id获取{{.StructName}}表记录
func (c *{{.StructName}}) FirstV0() (p *{{.StructName}}, err error) {
	result := global.DB.Where(c).First(&p)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

// First 根据id获取{{.StructName}}表记录
func (c *{{.StructName}}) First() (err error) {
	result := global.DB.Where(c).First(&c)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	return nil
	//}

	return result.Error
}

// Exist 根据id获取{{.StructName}}表记录
func (c *{{.StructName}}) Exist() (bool, error) {

	count, err := c.Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Count 根据id获取{{.StructName}}表记录
func (c *{{.StructName}}) Count() (count int64, err error) {

	if err := global.DB.Model(&{{.StructName}}{}).Where(c).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}



// source code
import (
	"context"
	"github.com/pkg/errors"
	"github.com/wordpress-plus/server-core/service/system"
	"gorm.io/gorm"
)

const initOrder{{.StructName}} = 5000

type init{{.StructName}} struct{}

// auto run
func init() {
	system.RegisterInit(initOrder{{.StructName}}, &init{{.StructName}}{})
}

func (i *init{{.StructName}}) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&model.{{.StructName}}{})
}

func (i *init{{.StructName}}) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&model.{{.StructName}}{})
}

func (i *init{{.StructName}}) InitializerName() string {
	return (&model.{{.StructName}}{}).TableName() + constant.
}

func (i *init{{.StructName}}) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []model.{{.StructName}}{}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", (&model.{{.StructName}}{}).TableName())
	}

	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *init{{.StructName}}) DataInserted(ctx context.Context) bool {
	_, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}

	return true
}
