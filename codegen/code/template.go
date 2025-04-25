package code

import (
	"os"
	"path/filepath"
	"text/template"
)

func Generate(model ModelInfo, tplpath string, outpath string) error {
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
	tpl.Execute(file, model)
	return nil
}
