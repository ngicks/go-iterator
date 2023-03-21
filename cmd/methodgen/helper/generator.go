package helper

import (
	"go/ast"
	"io"
	"sort"

	. "github.com/dave/jennifer/jen"
)

const (
	autoGenerationNotice = "// Code generated by github.com/ngicks/iterator/cmd/methodgen. DO NOT EDIT."
	defQualifier         = "github.com/ngicks/iterator/def"
)

func GetDefaultImports() map[string]string {
	return map[string]string{
		defQualifier: "def",
	}
}

type Generator struct {
	PkgName string
	Imports map[string]string

	f *File
}

func (g *Generator) GenPreamble() {
	f := NewFile(g.PkgName)
	f.PackageComment(autoGenerationNotice)
	f.PackageComment(ignoreCommentSlash + "all")
	f.ImportNames(g.Imports)

	g.f = f
}

func (g *Generator) GenSizeHint(typeInfo map[string]TypeInfo) {
	f := g.f

	f.Comment("SizeHint")
	for _, key := range keysStable(typeInfo) {
		info := typeInfo[key]
		f.
			Func(). // func
			Params(
				Id("iter").
					Id(info.TypeName).
					Types(stringToLitSlice(info.TypeParams)...),
			). // (iter <TypeName>[T, U])
			Id("SizeHint").
			Params(). // SizeHint()
			Int().    // int
			Block(
				If(
					List(Id("sizeHinter"), Id("ok")).
						Op(":=").Id("iter").Dot(info.InnerMemberName).Assert(Qual(defQualifier, "SizeHinter")),
					Id("ok"),
				).Block(
					Return(Id("sizeHinter").Dot("SizeHint").Call()),
				),
				Return(
					Lit(-1),
				),
			)
	}
}

func (g *Generator) GenReverse(typeInfo map[string]TypeInfo) {
	f := g.f

	f.Comment("Reverse")
	for _, key := range keysStable(typeInfo) {
		info := typeInfo[key]
		f.
			Func(). // func
			Params(
				Id("iter").
					Id(info.TypeName).
					Types(stringToLitSlice(info.TypeParams)...),
			). // (iter <TypeName>[T, U])
			Id("ReverseRaw").
			Params(). // ReverseRaw()
			Parens(
				List(
					Id("rev").Op("*").Id(info.TypeName).Types(stringToLitSlice(info.TypeParams)...),
					Id("ok").Bool(),
				),
			).     // (rev *<TypeName>[T,U], ok bool)
			Block( // { if rev, ok := Reverse(<inner>); ok {return &<Ty>[T,U]{<innerIter>:rev, ...reset of members unchanged ...}, true} else { return nil, false}}
				If(
					List(Id("reversedInner"), Id("ok")).
						Op(":=").Id("Reverse").Call(Id("iter").Dot(info.InnerMemberName)),
					Id("ok"),
				).Block(
					Return(
						Op("&").Id(info.TypeName).
							Types(stringToLitSlice(info.TypeParams)...).Block(
							append(
								[]Code{Id(info.InnerMemberName).Op(":").Id("reversedInner").Op(",")},
								restOfMember(GetMemberNames(info.StructType), "iter", info.InnerMemberName)...,
							)...,
						),
						True(),
					),
				),
				Return(
					Nil(), False(),
				),
			)

		f.
			Func(). // func
			Params(
				Id("iter").
					Id(info.TypeName).
					Types(stringToLitSlice(info.TypeParams)...),
			). // (iter <TypeName>[T, U])
			Id("Reverse").
			Params(). // ReverseRaw()
			Parens(
				List(
					Id("rev").Qual(defQualifier, "SeIterator").
						Types(Id(info.TypeParams[len(info.TypeParams)-1])),
					Id("ok").Bool(),
				),
			). // (rev *<TypeName>[T,U], ok bool)
			Block(
				Return(
					Id("iter").Dot("ReverseRaw").Call(),
				),
			)
	}
}

func (g *Generator) Write(w io.Writer) error {
	return g.f.Render(w)
}

func keysStable[T any](m map[string]T) []string {
	var out sort.StringSlice
	for k := range m {
		out = append(out, k)
	}
	sort.Sort(out)
	return out
}

func stringToLitSlice(input []string) []Code {
	out := make([]Code, len(input))

	for i := 0; i < len(input); i++ {
		out[i] = Id(input[i])
	}

	return out
}

func GetMemberNames(node *ast.StructType) []string {
	retSlice := []string{}

	for _, field := range node.Fields.List {
		if len(field.Names) > 0 {
			retSlice = append(retSlice, field.Names[0].Name)
		} else {
			// embedded
			idx, ok := field.Type.(*ast.IndexExpr)
			if !ok {
				continue
			}
			if ident, ok := idx.X.(*ast.Ident); ok {
				retSlice = append(retSlice, ident.Name)
			} else {
				selector, ok := idx.X.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				retSlice = append(retSlice, selector.Sel.Name)
			}
		}
	}

	return retSlice
}

func restOfMember(memberNames []string, id string, excludes string) []Code {
	out := make([]Code, 0, len(memberNames))
	for _, name := range memberNames {
		if name == excludes {
			continue
		}

		out = append(out, Id(name).Op(":").Id(id).Dot(name).Op(","))
	}
	return out
}
