package goimport

import (
    "go/ast"
    "go/parser"
    "go/token"
    "io/ioutil"
    "strconv"
)

type Source struct {
    FileName  string
    Namespace string
    Imports   []*ImportPath
}

func NewSource(fileName string, factory *ImportPathFactory) (*Source, error) {

    contents, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }
    fset := token.NewFileSet()
    tree, err := parser.ParseFile(fset, fileName, string(contents), 0)
    if err != nil {
        return nil, err
    }
    return &Source{
        fileName, tree.Name.Name, collectPathValue(tree.Imports, factory),
    }, nil

}

func unescape(target string) string {
    str, _ := strconv.Unquote(target)
    return str
}

func collectPathValue(imports []*ast.ImportSpec, factory *ImportPathFactory) []*ImportPath {
    r := make([]*ImportPath, 0, len(imports))
    for _, path := range imports {
        pathName := unescape(path.Path.Value)
        importPath := factory.Get(pathName)
        if importPath != nil {
            r = append(r, importPath)
        }
    }
    return r
}
