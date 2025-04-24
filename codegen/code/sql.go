package code

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

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

func ModelInfoFromTableInfo(tableInfo []TableInfo) ModelInfo {
	fields := make([]Field, len(tableInfo))
	for i, v := range tableInfo {
		fields[i].FieldName.DbName = v.ColumnName
		fields[i].FieldName.Camel = CamelCase(v.ColumnName)
		fields[i].FieldName.Pascal = PascalCase(v.ColumnName)
		fields[i].FieldName.Snake = SnakeCase(v.ColumnName)
		dtype := strings.ToUpper(v.DataType)
		fidleType := FieldTypeInfo{
			DataType: dtype,
			Type:     toRealType(dtype),
			Len:      GetLen(v.Columntype),
			Comment:  v.ColumnComment,
		}
		fields[i].FieldType = fidleType
	}
	var modelInfo ModelInfo
	modelInfo.Fields = fields
	modelInfo.ModelName.DbName = tableInfo[0].TableName
	modelInfo.ModelName.Camel = CamelCase(tableInfo[0].TableName)
	modelInfo.ModelName.Pascal = PascalCase(tableInfo[0].TableName)
	modelInfo.ModelName.Snake = SnakeCase(tableInfo[0].TableName)
	return modelInfo
}

func GetLen(s string) int {
	if s == "" {
		return 0
	}
	start := strings.Index(s, "(")
	end := strings.Index(s, ")")
	if start == -1 || end == -1 || start >= end {
		return 0
	}
	lenstr := s[start+1 : end]
	len, err := strconv.Atoi(lenstr)
	if err != nil {
		return 0
	}
	return len
}

func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	var result strings.Builder
	var upperNext bool

	for i, r := range s {
		if r == '_' {
			upperNext = true
			continue
		}

		if upperNext {
			result.WriteRune(unicode.ToUpper(r))
			upperNext = false
		} else {
			if i == 0 {
				result.WriteRune(unicode.ToLower(r))
			} else {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}

func PascalCase(s string) string {
	if s == "" {
		return ""
	}
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			// 将每个部分的首字母转换为大写
			runes := []rune(part)
			runes[0] = unicode.ToUpper(runes[0])
			parts[i] = string(runes)
		}
	}

	// 拼接所有部分为一个完整的字符串
	return strings.Join(parts, "")
}

func SnakeCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				runes[i] = '_' + unicode.ToLower(r)
			} else {
				runes[i] = unicode.ToLower(r)
			}
		}
	}
	return string(runes)
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

func FetchModelInfo(db *sqlx.DB) ModelInfo {
	var id []TableInfo
	err := db.Select(&id, `SELECT TABLE_NAME,COLUMN_NAME,IS_NULLABLE,DATA_TYPE,COLUMN_TYPE,COLUMN_COMMENT
						FROM information_schema.COLUMNS
						WHERE table_schema = 'db_ems_monitor'
						AND table_name='df_hotel'
						ORDER BY  ORDINAL_POSITION`)
	if err != nil {
		panic(err)
	}
	f := ModelInfoFromTableInfo(id)
	fmt.Println(f)
	return f
}
