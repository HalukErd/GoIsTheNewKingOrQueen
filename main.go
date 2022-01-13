package main

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type turkishBot struct{}

func main() {
	var eb englishBot
	var sb spanishBot
	var tb turkishBot
	printGreeting(eb)
	printGreeting(sb)
	printGreeting(tb)
}

func printGreeting(b bot) {
	println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//	eb.getGreeting()
//}
//func printGreeting(sb spanishBot) {
//	sb.getGreeting()
//}
//func printGreeting(tb turkishBot) {
//	tb.getGreeting()
//}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (turkishBot) getGreeting() string {
	return "Selam!"
}
