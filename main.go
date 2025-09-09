package main

import "fmt"

func main() {
	InitDB()
	defer DB.Close()

	fmt.Println("DB init")

}
