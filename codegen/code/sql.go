package code

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:1@tcp(192.168.1.94:3306)/2")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
}

type TableInfo struct {
	TableName     string `db:"table_name"`
	ColumnName    string `db:"column_name"`
	IsNullable    string `db:"is_nullable"`
	DataType      string `db:"data_type"`
	Columntype    string `db:"column_type"`
	ColumnComment string `db:"column_comment"`
}

type ModelInfo struct {
	Package   string
	TableName NameStyle
	Columns   []ColumnInfo
}

type ColumnInfo struct {
	ColumnName NameStyle
	DataType   DataTypeWithLen
}

type DataTypeWithLen struct {
	DataType string
	Len      int
}

type NameStyle struct {
	DbName string
	Camel  string
	Pascal string
	Snake  string
}

func ModelInfoFromTableInfo(tableInfo []TableInfo) ModelInfo {
	columns := make([]ColumnInfo, len(tableInfo))
	for i, v := range tableInfo {
		columns[i].ColumnName.DbName = v.ColumnName
		columns[i].ColumnName.Camel = CamelCase(v.ColumnName)
		columns[i].ColumnName.Pascal = PascalCase(v.ColumnName)
		columns[i].ColumnName.Snake = SnakeCase(v.ColumnName)
		dataType := DataTypeWithLen{
			DataType: v.DataType,
			Len:      GetLen(v.Columntype),
		}
		columns[i].DataType = dataType
	}
	var modelInfo ModelInfo
	modelInfo.Columns = columns
	modelInfo.TableName.DbName = tableInfo[0].TableName
	modelInfo.TableName.Camel = CamelCase(tableInfo[0].TableName)
	modelInfo.TableName.Pascal = PascalCase(tableInfo[0].TableName)
	modelInfo.TableName.Snake = SnakeCase(tableInfo[0].TableName)
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

func FetchModelInfo() ModelInfo {
	var id []TableInfo
	err := Db.Select(&id, `SELECT table_name,column_name,is_nullable,data_type,column_type,column_comment
						FROM information_schema.COLUMNS
						WHERE table_schema = 'db_ems_monitor'
						AND table_name='df_station'`)
	defer Db.Close()
	if err != nil {
		panic(err)
	}
	f := ModelInfoFromTableInfo(id)
	fmt.Println(f)
	return f
}
