package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func main() {
	//func1()
	//func2()
	//func3()
	//func4()
	//func5()
	//func6()
	//func7()
	//func8()
	func9()
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type house struct {
	Name string
	Mark string
}
type valyrianSword struct {
	Name  string
	Owner string
}

func func9() {
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

	err := tpl.ExecuteTemplate(os.Stdout, "tplStructSlice.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func func8() {
	data := house{
		Name: "Lannister",
		Mark: "Lion",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tplStruct.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func func7() {
	data := map[string]string{
		"Lannisters": "Lion",
		"Stark":      "DireWolf",
		"Targaryen":  "Dragon",
		"Baratheon":  "Stag",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tplMap.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
func func6() {
	data := []string{"This", "is", "passing", "data", "as a slice"}
	err := tpl.ExecuteTemplate(os.Stdout, "tplSlice.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

func func5() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

func func4() {
	t, err := template.ParseFiles("templates/one.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	t, err = t.ParseFiles("templates/two.gohtml", "templates/vespa.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func func3() {
	t, err := template.ParseFiles("templates/text.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, errCrt := os.Create("index.html")
	if errCrt != nil {
		log.Fatalln(errCrt)
	}

	err = t.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func func2() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	html := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8"/>
			<title>Hello World</title>
			</head>
			<body>
			<h1>` + name + `</h1>
			</body>
			</html>`

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating file.", err)
		os.Exit(1)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(html))
}

func func1() {
	name := "Haluk Erd"

	html := `<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="UTF-8"/>
			<title>Hello World</title>
			</head>
			<body>
			<h1>` + name + `</h1>
			</body>
			</html>`
	fmt.Println(html)
	// go run main.go > index.html
}
