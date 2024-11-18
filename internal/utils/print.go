package utils

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func PrintRed(s string) {
	println(Red + s + Reset)
}

func PrintGreen(s string) {
	println(Green + s + Reset)
}

func PrintYellow(s string) {
	println(Yellow + s + Reset)
}

func PrintBlue(s string) {
	println(Blue + s + Reset)
}

func PrintMagenta(s string) {
	println(Magenta + s + Reset)
}

func PrintCyan(s string) {
	println(Cyan + s + Reset)
}

func PrintGray(s string) {
	println(Gray + s + Reset)
}

func PrintWhite(s string) {
	println(White + s + Reset)
}
