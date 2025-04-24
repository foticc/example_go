package code

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func Generate(model ModelInfo, tplpath string, outpath string) error {
	fmt.Println("Vaildating template file...")
	tpl, err := template.ParseFiles(tplpath)
	if err != nil {
		return err
	}
	os.Mkdir(filepath.Dir(outpath), 0777)
	file, err := os.Create(outpath)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Println("Generating code...")
	tpl.Execute(file, model)
	return nil
}
