package main

import "github.com/testsla/jojo/server/Admin/core"

func main() {
	core.New().Router().DBConnect().Middleware().HttpListen()
}
