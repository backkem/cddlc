package gogen

import (
	"fmt"
	gast "go/ast"
	"go/token"
	"strconv"
	"strings"

	"github.com/HannesKimara/cddlc/ast"
	"github.com/iancoleman/strcase"
)

type stringConverter func(string) string

var converters map[string]stringConverter

func (g *Generator) formatIdentifier(name string) *gast.Ident {
	formatted := strcase.ToCamel(name)
	formatted = g.typePrefix + formatted

	return &gast.Ident{
		Name: formatted,
	}
}

func (g *Generator) transpileIdentifier(ident *ast.Identifier, comments ...*ast.Comment) *gast.Ident {
	formatted := strcase.ToCamel(ident.Name)

	if ident.IsSocket() || ident.IsPlug() {
		formatted = strings.TrimLeft(ident.Name, "$")
	}

	if token.IsIdentifier(formatted) {
		return g.formatIdentifier(formatted)
	}
	_, err := strconv.ParseInt(formatted, 0, 64)
	if err != nil {
		panic(fmt.Sprintf("Identifier %s -> %s could not be transformed to valid go identifier", ident.Name, formatted))
	}

	name := "IntKey_" + formatted

	comment := ""
	if len(comments) > 0 && comments[0] != nil {
		comment = comments[0].Text
	}
	if len(comment) > 0 {
		name = strcase.ToCamel(comment)
	}
	return &gast.Ident{
		Name: name,
	}
}

func init() {
	converters := make(map[string]stringConverter)

	converters["camelcase"] = strcase.ToCamel
	converters["snakecase"] = strcase.ToSnake
}
