package code

import (
	"codegen/utils"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TableInfo struct {
	TableName     string `db:"TABLE_NAME"`
	ColumnName    string `db:"COLUMN_NAME"`
	IsNullable    string `db:"IS_NULLABLE"`
	DataType      string `db:"DATA_TYPE"`
	Columntype    string `db:"COLUMN_TYPE"`
	ColumnComment string `db:"COLUMN_COMMENT"`
	ColumnKey     string `db:"COLUMN_KEY"`
}

type ModelInfo struct {
	Package   string
	ModelName NameStyle
	Fields    []Field
}

type Field struct {
	FieldName NameStyle
	FieldType FieldTypeInfo
}

type FieldTypeInfo struct {
	DataType   string
	Type       string
	Len        int
	Comment    string
	IsNullable bool
	IsPrimary  bool
}

type NameStyle struct {
	DbName    string
	Camel     string
	Pascal    string
	Snake     string
	KebabCase string
}

func ModelInfoFromTableInfo(tableInfo []TableInfo, params CustomParameters) ModelInfo {
	fields := make([]Field, len(tableInfo))
	for i, v := range tableInfo {
		fields[i].FieldName = NameStyle{
			DbName:    v.ColumnName,
			Camel:     utils.CamelCase(v.ColumnName),
			Pascal:    utils.PascalCase(v.ColumnName),
			Snake:     utils.SnakeCase(v.ColumnName),
			KebabCase: utils.KebabCase(v.ColumnName),
		}
		dtype := strings.ToUpper(v.DataType)
		fields[i].FieldType = FieldTypeInfo{
			DataType:   dtype,
			Type:       toRealType(params.GenType, dtype),
			Len:        utils.GetLen(v.Columntype),
			Comment:    v.ColumnComment,
			IsNullable: v.IsNullable == "YES",
			IsPrimary:  v.ColumnKey == "PRI",
		}
	}
	var modelInfo ModelInfo
	modelInfo.Package = params.PackageName
	modelInfo.Fields = fields
	modelInfo.ModelName = NameStyle{
		DbName:    tableInfo[0].TableName,
		Camel:     utils.CamelCase(params.ModelName),
		Pascal:    utils.PascalCase(params.ModelName),
		Snake:     utils.SnakeCase(params.ModelName),
		KebabCase: utils.KebabCase(params.ModelName),
	}

	return modelInfo
}

var javaMapping = map[string][]string{
	"Boolean":       {"BIT"},
	"Integer":       {"TINYINT", "SMALLINT", "MEDIUMINT", "INT"},
	"Long":          {"BIGINT"},
	"Float":         {"FLOAT"},
	"Double":        {"DOUBLE"},
	"BigDecimal":    {"DECIMAL"},
	"String":        {"CHAR", "VARCHAR", "JSON", "TEXT"},
	"LocalDateTime": {"DATETIME"},
	"LocalDate":     {"DATE"},
}

var tsMapping = map[string][]string{
	"number":  {"TINYINT", "SMALLINT", "MEDIUMINT", "INT", "BIGINT", "FLOAT", "DOUBLE", "DECIMAL"},
	"string":  {"CHAR", "VARCHAR", "JSON", "TEXT", "DATETIME", "DATE"},
	"boolean": {"BIT"},
}

func toRealType(genType string, datatype string) string {
	reverseMap := make(map[string]string)
	var mapping map[string][]string
	if genType == "java" {
		mapping = javaMapping
	} else {
		mapping = tsMapping
	}
	for key, values := range mapping {
		for _, value := range values {
			reverseMap[value] = key
		}
	}
	return reverseMap[datatype]
}

type CustomParameters struct {
	Schema      string
	TableName   string
	ModelName   string
	PackageName string
	GenType     string
}

func FetchModelInfo(db *sqlx.DB, params CustomParameters) ModelInfo {
	var id []TableInfo
	sql := fmt.Sprintf(`SELECT TABLE_NAME,COLUMN_NAME,IS_NULLABLE,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT,COLUMN_KEY
						FROM information_schema.COLUMNS
						WHERE table_schema = '%s'
						AND table_name='%s'
						ORDER BY  ORDINAL_POSITION`, params.Schema, params.TableName)
	err := db.Select(&id, sql)
	if err != nil {
		panic(err)
	}
	f := ModelInfoFromTableInfo(id, params)
	return f
}
