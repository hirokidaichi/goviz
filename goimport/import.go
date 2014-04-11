package goimport

import (
    "fmt"
    "github.com/hirokidaichi/goviz/dotwriter"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

type ImportPath struct {
    ImportPath string
    Files      []*Source
}

var _ dotwriter.IDotNode = &ImportPath{}

func NewImportPath(importPath string, filter *ImportFilter) *ImportPath {
    return &ImportPath{ImportPath: importPath}
}

func (self *ImportPath) Init(factory *ImportPathFactory, fileNames []string) {

    sourceFiles := make([]*Source, len(fileNames))
    for idx, fileName := range fileNames {
        source, err := NewSource(fileName, factory)
        if err != nil {
            panic(err)
        }
        sourceFiles[idx] = source

    }
    self.Files = sourceFiles
}

func (self *ImportPath) Label() string {
    if !self.HasFiles() {
        return self.ImportPath
    }
    return fmt.Sprintf("%s|%s|%s",
        self.Files[0].Namespace,
        self.ImportPath,
        strings.Join(self.FileNames(), `\n`))
}

func (self *ImportPath) Name() string { return self.ImportPath }

func (self *ImportPath) Shape() string {
    if !self.HasFiles() {
        return "oval"
    }
    return "record"
}

func (self *ImportPath) Style() string {
    if !self.HasFiles() {
        return "dashed"
    }
    return "solid"

}

func (self *ImportPath) Children() []dotwriter.IDotNode {
    list := make([]dotwriter.IDotNode, 0)
    for _, f := range self.Files {
        for _, d := range f.Imports {
            list = append(list, d)
        }
    }
    return list
}

func (p *ImportPath) HasFiles() bool {
    return (len(p.Files) != 0)
}

func (p *ImportPath) FileNames() []string {
    fileNames := make([]string, len(p.Files))
    for idx, v := range p.Files {
        fileNames[idx] = filepath.Base(v.FileName)
    }
    return fileNames
}

func (p *ImportPath) String() string {
    return fmt.Sprintf("%s:\n%s", p.ImportPath, p.Files)
}

func fileExists(file string) bool {
    _, err := os.Stat(file)
    return !os.IsNotExist(err)
}

func goSrc() string {
    return filepath.Join(os.Getenv("GOPATH"), "src")
}

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
