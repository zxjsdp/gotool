package main

import (
	"fmt"

	"github.com/zxjsdp/gotool/gotool"
)

const (
	PackageDirName = "gotool"
	BaseInfo       = `gotool: Collection of golang utils & tools.

	GitHub: https://github.com/zxjsdp/gotool
	Installation: go get github.com/zxjsdp/gotool

All functions provided by gotool:
`
)

func main() {
	funcNameMap := gotool.GetAllFuncNames(PackageDirName)
	fmt.Println(BaseInfo)
	for packageName, funcNames := range funcNameMap {
		fmt.Printf("\tpackage: %s\n", packageName)
		for i, funcName := range funcNames {
			fmt.Printf("\t\t%d: %s\n", i+1, funcName)
		}
	}
}
