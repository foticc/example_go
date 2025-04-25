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
	DataType string
	Type     string
	Len      int
	Comment  string
}

type NameStyle struct {
	DbName string
	Camel  string
	Pascal string
	Snake  string
}

func ModelInfoFromTableInfo(tableInfo []TableInfo, params CustomParameters) ModelInfo {
	fields := make([]Field, len(tableInfo))
	for i, v := range tableInfo {
		fields[i].FieldName.DbName = v.ColumnName
		fields[i].FieldName.Camel = utils.CamelCase(v.ColumnName)
		fields[i].FieldName.Pascal = utils.PascalCase(v.ColumnName)
		fields[i].FieldName.Snake = utils.SnakeCase(v.ColumnName)
		dtype := strings.ToUpper(v.DataType)
		fidleType := FieldTypeInfo{
			DataType: dtype,
			Type:     toRealType(dtype),
			Len:      utils.GetLen(v.Columntype),
			Comment:  v.ColumnComment,
		}
		fields[i].FieldType = fidleType
	}
	var modelInfo ModelInfo
	modelInfo.Package = params.PackageName
	modelInfo.Fields = fields
	modelInfo.ModelName.DbName = tableInfo[0].TableName
	modelInfo.ModelName.Camel = utils.CamelCase(params.ModelName)
	modelInfo.ModelName.Pascal = utils.PascalCase(params.ModelName)
	modelInfo.ModelName.Snake = utils.SnakeCase(params.ModelName)
	return modelInfo
}

var mapping = map[string][]string{
	"Boolean":       {"BIT"},
	"Integer":       {"TINYINT", "SMALLINT", "MEDIUMINT"},
	"Long":          {"BIGINT"},
	"Float":         {"FLOAT"},
	"Double":        {"DOUBLE"},
	"BigDecimal":    {"DECIMAL"},
	"String":        {"CHAR", "VARCHAR"},
	"LocalDateTime": {"DATETIME"},
	"LocalDate":     {"DATE"},
}

func toRealType(datatype string) string {
	reverseMap := make(map[string]string)
	for key, values := range mapping {
		for _, value := range values {
			reverseMap[value] = key
		}
	}
	return reverseMap[datatype]
}

type CustomParameters struct {
	TableName   string
	ModelName   string
	PackageName string
}

func FetchModelInfo(db *sqlx.DB, params CustomParameters) ModelInfo {
	var id []TableInfo
	sql := fmt.Sprintf(`SELECT TABLE_NAME,COLUMN_NAME,IS_NULLABLE,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT
						FROM information_schema.COLUMNS
						WHERE table_schema = 'db_ems_monitor'
						AND table_name='%s'
						ORDER BY  ORDINAL_POSITION`, params.TableName)
	err := db.Select(&id, sql)
	if err != nil {
		panic(err)
	}
	f := ModelInfoFromTableInfo(id, params)
	return f
}
