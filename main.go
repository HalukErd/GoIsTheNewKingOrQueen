package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

type house struct {
	Name string
	Mark string
}
type valyrianSword struct {
	Name  string
	Owner string
}

func main() {
	func1()
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func func1() {
	l := house{
		Name: "Lannister",
		Mark: "Lion",
	}
	s := house{
		Name: "Stark",
		Mark: "DireWolf",
	}
	t := house{
		Name: "Targaryen",
		Mark: "Dragon",
	}

	ww := valyrianSword{
		Name:  "Widow's Wail",
		Owner: "Jaime Lannister",
	}

	ok := valyrianSword{
		Name:  "Oathkeeper",
		Owner: "Brienne Tarth",
	}

	lc := valyrianSword{
		Name:  "Longclaw",
		Owner: "Jon Snow",
	}

	data := struct {
		Houses         []house
		ValyrianSwords []valyrianSword
	}{
		[]house{l, s, t},
		[]valyrianSword{ww, ok, lc},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
