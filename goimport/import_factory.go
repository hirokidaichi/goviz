package goimport

import (
    "path/filepath"
    "strings"
)

func ParseRelation(
    rootPath string, seekPath string, leafVisibility bool) *ImportPathFactory {

    factory := NewImportPathFactory(
        rootPath,
        seekPath,
        leafVisibility,
    )
    factory.Root = factory.Get(rootPath)
    return factory

}

type ImportPathFactory struct {
    Root   *ImportPath
    Filter *ImportFilter
    Pool   map[string]*ImportPath
}

func NewImportPathFactory(
    rootPath string, seekPath string, leafVisibility bool) *ImportPathFactory {

    self := &ImportPathFactory{Pool: make(map[string]*ImportPath)}
    filter := NewImportFilter(
        rootPath,
        seekPath,
        leafVisibility,
    )
    self.Filter = filter
    return self
}
func (self *ImportPathFactory) GetRoot() *ImportPath {
    return self.Root
}

func (self *ImportPathFactory) GetAll() []*ImportPath {
    ret := make([]*ImportPath, 0)
    for _, value := range self.Pool {
        ret = append(ret, value)
    }
    return ret
}

func (self *ImportPathFactory) Get(importPath string) *ImportPath {
    // aquire from pool
    pool := self.Pool
    if _, ok := pool[importPath]; ok {
        return pool[importPath]
    }
    filter := self.Filter
    // if not applicable return nullobject
    if !filter.Applicable(importPath) {
        // if invisible return nil
        if !filter.Visible(importPath) {
            return nil
        }
        pool[importPath] = &ImportPath{
            ImportPath: importPath}
        return pool[importPath]
    }

    dirPath := filepath.Join(goSrc(), importPath)
    if !fileExists(dirPath) {
        // if invisible return nil
        if !filter.Visible(importPath) {
            return nil
        }
        pool[importPath] = &ImportPath{
            ImportPath: importPath}
        return pool[importPath]
    }
    ret := &ImportPath{
        ImportPath: importPath,
    }
    pool[importPath] = ret
    fileNames := glob(dirPath)
    ret.Init(self, fileNames)
    return ret
}

//ImportFilter
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
