package extractor

import "go/ast"

func ExtractFieldsForStruct(f *ast.File) map[string][]*ast.Field {
	structName2Fields := make(map[string][]*ast.Field, 0)

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			structName := typeSpec.Name.Name
			structName2Fields[structName] = structType.Fields.List
		}
	}

	return structName2Fields
}
