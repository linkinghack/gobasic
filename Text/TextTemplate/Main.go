package main

import (
	"fmt"
	"os"
	"text/template"
)

type Member struct {
	ID   string
	Name string
	Age  uint
}

func funcInTemplate(source string) string {
	return "<<<" + source + ">>>"
}

func main() {
	// Basic usage of text/template
	firstTemplate := `{{range . -}} ID:{{.ID}}, Name:{{.Name}}, Age:{{.Age}} {{end}}`
	tmpl := template.Must(template.New("first").Parse(firstTemplate))
	var members []Member
	members = append(members, Member{"001", "Tom", 27})
	members = append(members, Member{"002", "Jerry", 18})
	tmpl.Execute(os.Stdout, members)

	fmt.Println("")

	// Function
	secondTemplate := "U can call a user self function: {{userfunc1 .}}"
	// register a user function to the template context, before Parse()
	tmpl2 := template.Must(template.New("second").Funcs(template.FuncMap{"userfunc1": funcInTemplate}).Parse(secondTemplate))

	tmpl2.Execute(os.Stdout, "This is input of user function")

}
