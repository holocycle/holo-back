package main

import (
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
	"local.package/pkg/api"
)

func main() {
	converter := typescriptify.New()
	converter.CreateFromMethod = true
	converter.Indent = "  "
	converter.BackupDir = "" // no backup

	for _, model := range api.Models() {
		converter.Add(model)
	}

	err := converter.ConvertToFile("src/index.ts")
	if err != nil {
		panic(err.Error())
	}
}
