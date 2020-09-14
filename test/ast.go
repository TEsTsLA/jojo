package main

import (
	"go/ast"
	//"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	src := []byte(`package main
		import "fmt"
		func main() {
		  fmt.Println("Hello, world!")
		}
`)

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open("build.go")
	add_data := "this is add!!!"
	if err != nil {
		//return nil,err
		_file, err := os.Create("build.go")
		if err != nil {
		}
		_file.Write([]byte(add_data))
	}
	f.Write([]byte(add_data))
	//延迟关闭文件
	defer f.Close()
	ast.Print(fset, file)
}
