package main

import (
	"embed"
	"fmt"
)

//go:embed text.txt static/*
var embedded embed.FS

func main() {
	file, _ := embedded.ReadFile("static/index.html")

	fmt.Println(string(file))
}
