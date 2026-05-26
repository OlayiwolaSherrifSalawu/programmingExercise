package main

import (
	"embed"
	"fmt"
)

//go:embed text.txt static/*
var Embedded embed.FS

func main() {
	file, _ := Embedded.ReadFile("static/index.html")

	fmt.Println(string(file))
}
