package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc":    strings.ToUpper,
	"ft":    firstThree,
	"fhdeu": timeAndDateEuropeFormat,
	"fsqrt": sqrt,
	"fdbl":  double,
	"fsq":   square,
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
	//func1()
	func2()
	func3()

}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
	//tpl = template.Must(template.ParseFiles("templates/*")) // ERROR
	//tpl.Funcs(fm)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func timeAndDateEuropeFormat(t time.Time) string {
	return t.Format("03:04:05 02/01/2006")
}

func sqrt(x float64) float64 {
	return math.Sqrt(float64(x))
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func double(x int) int {
	return x * 2
}

func func3() {
	err := tpl.ExecuteTemplate(os.Stdout, "tplPipeline.gohtml", 6)
	if err != nil {
		log.Fatalln(err)
	}
}

func func2() {
	//t := time.Now()
	//fmt.Println(t)
	//tf := t.Format("03:04/05 02&01-2006")
	//fmt.Println(tf)
	err := tpl.ExecuteTemplate(os.Stdout, "tplDate.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
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
