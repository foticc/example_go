package code

import (
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Generate() {
	tpl, err := template.ParseFiles("D:\\go_work_space\\learn\\codegen\\entity.tpl")

	check(err)

	file, err := os.Create("D:\\go_work_space\\learn\\codegen\\output.txt")

	check(err)
	defer file.Close()
	model := FetchModelInfo()
	model.Package = "com.example.hha"
	tpl.Execute(file, model)
}
