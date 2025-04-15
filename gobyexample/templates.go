package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("value is {{.}}")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("value is {{.}}\n"))

	t1.Execute(os.Stdout, "some value")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"C", "C++", "Java"})

	Create := func(name, str string) *template.Template {
		return template.Must(template.New(name).Parse(str))
	}

	t2 := Create("t2", "Name is {{.Name}}\n")
	t2.Execute(os.Stdout, struct{ Name string }{Name: "John"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Mary"})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "something")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} \n{{end}}\n")
	t4.Execute(os.Stdout, []string{"C", "C++", "Java"})
}
