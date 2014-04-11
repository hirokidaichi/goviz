package main

import (
    "flag"
    "github.com/hirokidaichi/goviz/dotwriter"
    "github.com/hirokidaichi/goviz/goimport"
    "math"
    "os"
)

var inputDir = flag.String("i", "", "Input Dir (like -i github.com/hirokidaihi/goviz)")
var outputFile = flag.String("o", "STDOUT", "Output File")

var depth = flag.Int("d", math.MaxInt8, "Ignore $GOROOT/pkg")

var seekPath = flag.String("seek-in", "", "Seek Root")
var plotLeaf = flag.Bool("l", false, "")

func main() {
    flag.Parse()

    filter := goimport.NewImportFilter(
        *inputDir,
        *seekPath,
        *plotLeaf,
    )
    root := goimport.NewGoImportPath(*inputDir, filter)
    if !root.HasFiles() {
        flag.Usage()
        os.Exit(1)
    }
    if 0 > *depth {
        flag.Usage()
        os.Exit(1)
    }
    // NewDotWriter(output,root,plotPath)
    output := getOutputWriter(*outputFile)
    writer := dotwriter.New(output)
    writer.MaxDepth = *depth
    writer.PlotGraph(root)

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
