package main

import (
	"codegen/code"
	"codegen/utils"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func initDataBase(param Params) {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s", param.User, param.Pwd, param.Host, param.Schema)
	fmt.Println("connect to database, uri:", uri)
	database, err := sqlx.Open("mysql", uri)
	if err != nil {
		fmt.Println("connect database failed,", err)
		return
	}

	Db = database
}

type Params struct {
	Host        string
	User        string
	Pwd         string
	Schema      string
	Table       string
	Pkg         string
	Model       string
	GenType     string
	TemplateDir string
	OutputDir   string
}

var supportGenType = []string{"java", "ts"}

func contains(arr []string, target string) bool {
	for _, value := range arr {
		if value == target {
			return true
		}
	}
	return false
}

var param = Params{}

func InitParams() error {
	flag.StringVar(&param.Host, "host", "", "host of database")
	flag.StringVar(&param.User, "user", "", "user of database")
	flag.StringVar(&param.Pwd, "pwd", "", "password of database")
	flag.StringVar(&param.Schema, "schema", "", "schema of database")
	flag.StringVar(&param.Table, "table", "", "table of schema")
	flag.StringVar(&param.Pkg, "pkg", "com.example", "package name")
	flag.StringVar(&param.Model, "model", "Example", "model name")
	flag.StringVar(&param.GenType, "lang", "java", "generate type name (java, ts)")
	flag.StringVar(&param.TemplateDir, "templateDir", "templates", "template path")
	flag.StringVar(&param.OutputDir, "outputDir", "output", "output path")
	flag.Parse()
	if param.Host == "" || param.User == "" || param.Pwd == "" || param.Schema == "" || param.Table == "" {
		return errors.New("host, user, pwd, schema,table must be set")
	}
	if !contains(supportGenType, param.GenType) {
		return errors.New("genType must be java or ts")
	}
	return nil
}

func dirPath(filename string) string {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(currentDirectory, filename)
}

func toOutputFilename(modulename string, fullpath string, param Params) string {
	templateDir := param.TemplateDir
	outputDir := param.OutputDir
	newpath := strings.ReplaceAll(fullpath, templateDir, outputDir)
	filename := filepath.Base(newpath)
	var prefix = ""
	if param.GenType == "java" {
		prefix = utils.PascalCase(modulename)
	} else {
		prefix = utils.CamelCase(modulename)
	}
	newfilename := prefix + strings.TrimSuffix(filename, filepath.Ext(filename)) + filepath.Ext(filename)
	return filepath.Join(strings.ReplaceAll(newpath, filename, ""), newfilename)
}

// codegen --host=192.168.1.94 --user=root --pwd=123456 --schema=db_ems_monitor --table=df_hotel --pkg=com.huwei.hotel.ems.monitor.interfaces.output --model=output
func main() {
	err := InitParams()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	initDataBase(param)
	fmt.Println("Start code generation...", param)
	// codegen.Generate()
	templatePath := dirPath(param.TemplateDir)
	files := []string{}
	filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否为文件，如果是，则添加到files切片中
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	parameters := code.CustomParameters{
		Schema:      param.Schema,
		TableName:   param.Table,
		ModelName:   param.Model,
		PackageName: param.Pkg,
		GenType:     param.GenType,
	}
	model := code.FetchModelInfo(Db, parameters)
	fmt.Printf("ModelInfo: %+v\n", model)

	for _, v := range files {
		t := toOutputFilename(param.Model, v, param)
		err := code.Generate(model, v, t)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	defer Db.Close()
}
