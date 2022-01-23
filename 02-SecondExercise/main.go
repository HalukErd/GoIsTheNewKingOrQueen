package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type address struct {
	Name   string
	Street string
	House  string
}

type hotel struct {
	Name    string
	Address address
	City    string
	Zip     string
	Region  string
}

var region = [3]string{"Southern", "Central", "Northern"}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	grandBudapest := hotel{
		Name: "Grand Budapest Hotel",
		Address: address{
			Name:   "Kadıköy Ev",
			Street: "Yurttaş",
			House:  "25/1",
		},
		City:   "İstanbul",
		Zip:    "35476",
		Region: region[0],
	}
	hotels := []hotel{grandBudapest}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)
}
