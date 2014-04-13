package main

import (
    "fmt"
    "github.com/hirokidaichi/goviz/dotwriter"
    "github.com/hirokidaichi/goviz/goimport"
    "github.com/hirokidaichi/goviz/metrics"
    "github.com/jessevdk/go-flags"
    "os"
)

type options struct {
    InputDir   string `short:"i" long:"input" required:"true" description:"intput ploject name"`
    OutputFile string `short:"o" long:"output" default:"STDOUT" description:"output file"`
    Depth      int    `short:"d" long:"depth" default:"128" description:"max plot depth of the dependency tree"`
    Reversed   string `short:"f" long:"focus" description:"focus on the specific module"`
    SeekPath   string `short:"s" long:"search" default:"" description:"top directory of searching"`
    PlotLeaf   bool   `short:"l" long:"leaf" default:"false" description:"whether leaf nodes are plotted"`
    UseMetrics bool   `short:"m" long:"metrics" default:"false" description:"display module metrics"`
}

func getOptions() *options {
    options := new(options)
    _, err := flags.Parse(options)
    if err != nil {
        os.Exit(1)
    }
    return options

}

func main() {
    options := getOptions()
    factory := goimport.ParseRelation(
        options.InputDir,
        options.SeekPath,
        options.PlotLeaf,
    )
    root := factory.GetRoot()
    if !root.HasFiles() {
        fmt.Fprintf(os.Stderr, "%s has no .go files\n", root.ImportPath)
        os.Exit(1)
    }
    if 0 > options.Depth {
        fmt.Fprintf(os.Stderr, "-d or --depth should have positive int\n")
        os.Exit(1)
    }
    output := getOutputWriter(options.OutputFile)
    if options.UseMetrics {
        metrics_writer := metrics.New(output)

        metrics_writer.Plot(pathToNode(factory.GetAll()))
        return
    }

    writer := dotwriter.New(output)
    writer.MaxDepth = options.Depth
    if options.Reversed == "" {
        writer.PlotGraph(root)
        return
    }
    writer.Reversed = true

    rroot := factory.Get(options.Reversed)
    if !rroot.HasFiles() {

        os.Exit(1)
    }

    writer.PlotGraph(rroot)

}

func pathToNode(pathes []*goimport.ImportPath) []dotwriter.IDotNode {
    r := make([]dotwriter.IDotNode, len(pathes))

    for i, _ := range pathes {
        r[i] = pathes[i]
    }
    return r
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
