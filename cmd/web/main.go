package main

import "fmt"

const version = "1.0.0"

//This is the execution point of the application
//config struct
type config struct {
	version string
	port    string
}

//application struct
type Application struct {
}

func main() {
	//cmd flags
	fmt.Println("Welcome to Larago")
}
