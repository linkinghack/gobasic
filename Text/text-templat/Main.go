package main
import (
	"text/template"
	"os"
)

type Inventory struct {
	Matrial string
	Count uint
}

func main() {
	sweater := Inventory{"wool", 17}
	templ, err := template.New("test").Parse("{{.Count}} items are make of {{.Matrial}}")
	if err != nil {panic(err)}
	err = templ.Execute(os.Stdout, sweater)
	if err != nil {panic(err)}
}