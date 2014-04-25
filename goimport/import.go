package goimport

import (
    "fmt"
    "github.com/hirokidaichi/goviz/dotwriter"
    "os"
    "path/filepath"
    "strings"
)

type ImportPath struct {
    ImportPath string
    Files      []*Source
    children   []dotwriter.IDotNode
    parents    []dotwriter.IDotNode
}

// ImportPath implements IDotNode
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

    for _, f := range self.Files {
        for _, d := range f.Imports {
            self.AddChild(d)
            d.AddParent(self)
        }
    }

}

func (self *ImportPath) AddChild(child dotwriter.IDotNode) {
    if self.children == nil {
        self.children = make([]dotwriter.IDotNode, 0)
    }
    self.children = append(self.children, child)
}

func (self *ImportPath) AddParent(parent dotwriter.IDotNode) {
    if self.parents == nil {
        self.parents = make([]dotwriter.IDotNode, 0)
    }
    self.parents = append(self.parents, parent)
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
    return self.children
}

func (self *ImportPath) Parents() []dotwriter.IDotNode {
    return self.parents
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
