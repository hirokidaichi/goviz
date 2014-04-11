package goimport

import (
    "fmt"
    "github.com/hirokidaichi/goviz/dotwriter"
    "go/ast"
    "go/parser"
    "go/token"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"
)

type GoImportPath struct {
    ImportPath string
    Files      []*GoFile
}

var InstancePool map[string]*GoImportPath = make(map[string]*GoImportPath)

func isMatched(pattern string, target string) bool {
    r, _ := regexp.Compile(pattern)
    return r.MatchString(target)
}

func glob(dirPath string) []string {
    fileNames, err := filepath.Glob(filepath.Join(dirPath, "/*.go"))
    if err != nil {
        panic("no gofiles")
    }

    files := make([]string, 0)

    for _, v := range fileNames {
        if isMatched("test", v) {
            continue
        }
        if isMatched("example", v) {
            continue
        }
        files = append(files, v)
    }
    return files
}

var _ dotwriter.IDotNode = &GoImportPath{}

func (self *GoImportPath) Label() string {
    if !self.HasFiles() {
        return self.ImportPath
    }
    return fmt.Sprintf("%s|%s|%s",
        self.Files[0].Namespace,
        self.ImportPath,
        strings.Join(self.FileNames(), `\n`))
}

func (self *GoImportPath) Name() string { return self.ImportPath }

func (self *GoImportPath) Shape() string {
    if !self.HasFiles() {
        return "oval"
    }
    return "record"
}

func (self *GoImportPath) Style() string {
    if !self.HasFiles() {
        return "dashed"
    }
    return "solid"

}

func (self *GoImportPath) Siblings() []dotwriter.IDotNode {
    list := make([]dotwriter.IDotNode, 0)
    for _, f := range self.Files {
        for _, d := range f.Imports {
            list = append(list, d)
        }
    }
    return list
}

func NewGoImportPath(importPath string, filter *ImportFilter) *GoImportPath {
    // aquire from InstancePool
    if _, ok := InstancePool[importPath]; ok {
        return InstancePool[importPath]
    }

    // if not applicable return nullobject
    if !filter.Applicable(importPath) {
        // if invisible return nil
        if !filter.Visible(importPath) {
            return nil
        }
        InstancePool[importPath] = &GoImportPath{
            ImportPath: importPath}
        return InstancePool[importPath]
    }

    dirPath := filepath.Join(GOSRC(), importPath)
    if !fileExists(dirPath) {
        // if invisible return nil
        if !filter.Visible(importPath) {
            return nil
        }
        InstancePool[importPath] = &GoImportPath{
            ImportPath: importPath}
        return InstancePool[importPath]
    }
    fileNames := glob(dirPath)

    ret := &GoImportPath{
        ImportPath: importPath,
    }
    InstancePool[importPath] = ret

    goFiles := make([]*GoFile, len(fileNames))
    for idx, fileName := range fileNames {
        goFile, err := NewGoFile(fileName, filter)
        if err != nil {
            panic(err)
        }
        goFiles[idx] = goFile

    }
    ret.Files = goFiles
    return ret

}

func (p *GoImportPath) HasFiles() bool {
    return (len(p.Files) != 0)
}

func (p *GoImportPath) FileNames() []string {
    fileNames := make([]string, len(p.Files))
    for idx, v := range p.Files {
        fileNames[idx] = filepath.Base(v.FileName)
    }
    return fileNames
}

func (p *GoImportPath) String() string {
    return fmt.Sprintf("%s:\n%s", p.ImportPath, p.Files)
}

func unescape(target string) string {
    str, _ := strconv.Unquote(target)
    return str
}

type GoFile struct {
    FileName  string
    Namespace string
    Imports   []*GoImportPath
}

func NewGoFile(fileName string, filter *ImportFilter) (*GoFile, error) {

    contents, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }
    fset := token.NewFileSet()
    tree, err := parser.ParseFile(fset, fileName, string(contents), 0)
    if err != nil {
        return nil, err
    }
    return &GoFile{
        fileName, tree.Name.Name, collectPathValue(tree.Imports, filter),
    }, nil

}

func fileExists(file string) bool {
    _, err := os.Stat(file)
    return !os.IsNotExist(err)
}

func collectPathValue(imports []*ast.ImportSpec, filter *ImportFilter) []*GoImportPath {
    r := make([]*GoImportPath, 0)
    for _, path := range imports {
        pathName := unescape(path.Path.Value)
        importPath := NewGoImportPath(pathName, filter)
        if importPath != nil {

            r = append(r, importPath)
        }
    }
    return r
}

func GOSRC() string {
    return filepath.Join(os.Getenv("GOPATH"), "src")
}

type ImportFilter struct {
    root     string
    seekPath string
    plotLeaf bool
}

func NewImportFilter(root string, seekPath string, plotLeaf bool) *ImportFilter {
    if seekPath == "SELF" {
        seekPath = root
    }
    impf := &ImportFilter{
        root:     root,
        seekPath: seekPath,
        plotLeaf: plotLeaf,
    }
    return impf

}

func (self *ImportFilter) Visible(path string) bool {
    return self.plotLeaf
}

func (self *ImportFilter) Applicable(path string) bool {
    if self.seekPath == "" {
        return true
    }
    if strings.Index(path, self.seekPath) == 0 {
        return true
    }
    return false
}
