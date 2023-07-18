package main

import (
	"github.com/LorezV/gophkeeper/internal/client/app"
	"github.com/LorezV/gophkeeper/internal/client/app/build"
)

func main() {
	build.CheckBuild()
	app.Execute()
}
