package main

import "fmt"

//int 16 = 16 bytes
//int 32 = 32 bytes
// int 64 = 64 bytes
//floatNum float32 = 32 bytes
//floatNum float64 = 64 bytes
//string = 32 bytes
//bool = 1 byte (0 or 1)
//+ - * / % ++ -- += -= *= /= %= == != > < >= <= && || ! & | ^ << >>
//Cannot do mix type operations in Go like floatnum32 + intnum16

// String type to store text
// var stringName string = "Hello World"
//Double quotes for string
//Single quotes for characters
//Back quote
//Int division rounded down

//String type to store string
// var stringName string = "Hello World"
//Double quotes for string
// concatenate string using +
//len number of bytes -> Outside of ASCII, it will not work properly. It's based on the encoding used. UTF-8 uses one byte per character. So if you have a string with only English letters and numbers, it will work properly. If you have a string with Chinese characters, it will not work properly.
//String to byte array
//RuneCOuntintString function
//Rune is a character in GoLang
//Default value of string is empty string, integer is 0, float is 0.0, bool is false
//Cannot declare constant using := , use = instead
//Constants are declared using const keyword
// const Pi float64 = 3.14 //Constant declaration with const keyword and type after name
// const Pi2 = 3.14 //Constant declaration with const keyword and type after name


func main() {
	// fmt.Println("Hello World")
	var intNum int16 = 20
	var intNum2 int32 = 30
	var intNum3 int64 = 40


	fmt.Println(intNum)
	fmt.Println(intNum2)
	fmt.Println(intNum3)
	printMe()
}

func printMe(){
	fmt.Println("Hello World")
}

