package main

import (
    "flag"
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "io"
    "io/ioutil"
    "math"
    "os"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"
)

var fp = fmt.Fprintln
var ff = fmt.Fprintf

var inputDir = flag.String("i", "", "Input Dir (like -i github.com/hirokidaihi/goviz)")
var outputFile = flag.String("o", "STDOUT", "Output File")
var ignoreTest = flag.Bool("ignore-test", false, "Ignore *_test.go File")
var ignoreLibs = flag.Bool("ignore-libs", false, "Ignore $GOROOT/pkg")
var level = flag.Int("level", math.MaxInt8, "Ignore $GOROOT/pkg")

func GOSRC() string {
    return filepath.Join(os.Getenv("GOPATH"), "src")
}

func main() {
    flag.Parse()

    root := NewGoImportPath(*inputDir)
    if !root.HasFiles() {
        flag.Usage()
        os.Exit(1)
    }
    if 0 > *level {
        flag.Usage()
        os.Exit(1)
    }
    output := getOutputWriter(*outputFile)
    fp(output, "digraph main{")
    fp(output, `edge[arrowhead="vee"]`)
    fp(output, "graph [rankdir=LR,compound=true,ranksep=1.0];")
    root.DumpRelation(output, *level)
    fp(output, "}")

}

func getOutputWriter(name string) *os.File {
    if name == "STDOUT" {
        return os.Stdout
    }
    if name == "STDERR" {
        return os.Stderr
    }
    f, _ := os.Create(name)
    return f
}

type GoImportPath struct {
    ImportPath string
    Files      []*GoFile
    isDumped   bool
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
    if !*ignoreTest {
        return fileNames
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
func NewGoImportPath(importPath string) *GoImportPath {
    if _, ok := InstancePool[importPath]; ok {
        return InstancePool[importPath]
    }
    dirPath := filepath.Join(GOSRC(), importPath)
    if !fileExists(dirPath) {
        InstancePool[importPath] = &GoImportPath{
            ImportPath: importPath, isDumped: false}
        return InstancePool[importPath]
    }
    fileNames := glob(dirPath)

    ret := &GoImportPath{
        ImportPath: importPath,
        isDumped:   false,
    }
    InstancePool[importPath] = ret

    goFiles := make([]*GoFile, len(fileNames))
    for idx, fileName := range fileNames {
        goFile, err := NewGoFile(fileName)
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

func (p *GoImportPath) IsSkipped() bool {
    return (*ignoreLibs && !p.HasFiles())
}
func (p *GoImportPath) DumpRelation(output io.Writer, depth int) {
    if p.isDumped {
        return
    }
    if depth <= 0 {
        return
    }
    p.isDumped = true
    p.DumpSetting(output)
    cache := make(map[string]bool)
    for _, f := range p.Files {
        for _, d := range f.Imports {
            if d.IsSkipped() {
                continue
            }
            key := strings.Join([]string{p.ImportPath, d.ImportPath}, "-")
            if !cache[key] {
                ff(output, "%s -> %s\n", escape(p.ImportPath), escape(d.ImportPath))
            }
            cache[key] = true
            d.DumpRelation(output, depth-1)
        }
    }

}

func (p *GoImportPath) DumpSetting(output io.Writer) {
    if !p.HasFiles() {
        ff(output, "%s[style=dashed]\n", escape(p.ImportPath))
        return
    }
    ff(output, `
    %s [
        shape = record
        label = "%s|%s|%s"
    ]
    `, escape(p.ImportPath),
        p.Files[0].Namespace,
        p.ImportPath,
        strings.Join(p.FileNames(), `\n`))
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

func escape(target string) string {
    return strconv.Quote(target)
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

func NewGoFile(fileName string) (*GoFile, error) {

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
        fileName, tree.Name.Name, collectPathValue(tree.Imports),
    }, nil

}

func fileExists(file string) bool {
    _, err := os.Stat(file)
    return !os.IsNotExist(err)
}

func collectPathValue(imports []*ast.ImportSpec) []*GoImportPath {
    r := make([]*GoImportPath, len(imports))
    for i, _ := range imports {
        r[i] = NewGoImportPath(unescape(imports[i].Path.Value))
    }
    return r
}
