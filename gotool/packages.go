package gotool

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func GetAllFuncNames(path string) map[string][]string {
	set := token.NewFileSet()
	packs, err := parser.ParseDir(set, path, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	funcMap := make(map[string][]*ast.FuncDecl, 0)
	for _, pack := range packs {
		for _, f := range pack.Files {
			key := pack.Name

			_, keyExists := funcMap[key]
			if !keyExists {
				funcMap[key] = []*ast.FuncDecl{}
			}

			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					funcMap[key] = append(funcMap[key], fn)
				}
			}
		}
	}

	funcNameMap := make(map[string][]string)
	for packAndFileName, funcs := range funcMap {
		funcNames := make([]string, 0)
		for _, f := range funcs {
			funcNames = append(funcNames, f.Name.Name)
		}
		funcNameMap[packAndFileName] = funcNames
	}

	return funcNameMap
}
