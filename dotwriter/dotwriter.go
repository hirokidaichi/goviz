package dotwriter

import (
    "fmt"
    "io"
    "strconv"
)

var FilesShown int

type IDotNode interface {
    Name() string
    Label(int) string
    Shape() string
    Style() string
    Children() []IDotNode
    Parents() []IDotNode
}

type DotWriter struct {
    output   io.Writer
    MaxDepth int
    Reversed bool
}

type plotCtx struct {
    nodeFlags map[string]bool
    edgeFlags map[string]bool
    level     int
}

func (ctx *plotCtx) isPlottedNode(node IDotNode) bool {
    _, ok := ctx.nodeFlags[node.Name()]
    return ok
}
func (ctx *plotCtx) setPlotted(node IDotNode) {
    _, ok := ctx.nodeFlags[node.Name()]
    if !ok {
        ctx.nodeFlags[node.Name()] = true
    }

}

func (ctx *plotCtx) isDepthOver() bool {
    return (ctx.level <= 0)
}
func (ctx *plotCtx) Deeper() *plotCtx {
    return &plotCtx{
        nodeFlags: ctx.nodeFlags,
        edgeFlags: ctx.edgeFlags,
        level:     ctx.level - 1,
    }
}
func newPlotContext(level int) *plotCtx {
    return &plotCtx{
        level:     level,
        nodeFlags: make(map[string]bool),
        edgeFlags: make(map[string]bool),
    }
}
func (ctx *plotCtx) isPlottedEdge(nodeA, nodeB IDotNode) bool {
    edgeName := fmt.Sprintf("%s->%s", nodeA.Name(), nodeB.Name())
    _, ok := ctx.edgeFlags[edgeName]
    if !ok {
        ctx.edgeFlags[edgeName] = true
    }
    return ok
}

func New(output io.Writer) *DotWriter {
    return &DotWriter{output: output}
}

func (dw *DotWriter) PlotGraph(root IDotNode) {
    dw.printLine("digraph main{")
    dw.printLine("\tedge[arrowhead=vee]")
    dw.printLine("\tgraph [rankdir=LR,compound=true,ranksep=1.0];")
    dw.plotNode(newPlotContext(dw.MaxDepth), root)
    dw.printLine("}")
}

func (dw *DotWriter) plotNode(ctx *plotCtx, node IDotNode) {
    if ctx.isPlottedNode(node) {
        return
    }
    if ctx.isDepthOver() {
        return
    }
    ctx.setPlotted(node)
    dw.plotNodeStyle(node)
    for _, s := range dw.getDependency(node) {
        dw.plotEdge(ctx, node, s)
        dw.plotNode(ctx.Deeper(), s)
    }
}

func (dw *DotWriter) getDependency(node IDotNode) []IDotNode {
    if dw.Reversed {
        return node.Parents()
    }
    return node.Children()
}
func (dw *DotWriter) plotNodeStyle(node IDotNode) {
    dw.printFormat("\t/* plot %s */\n", node.Name())
    dw.printFormat("\t%s[shape=%s,label=\"%s\",style=%s]\n",
        escape(node.Name()),
        escape(node.Shape()),
        node.Label(FilesShown),
        escape(node.Style()),
    )
}

func (dw *DotWriter) plotEdge(ctx *plotCtx, nodeA, nodeB IDotNode) {
    if ctx.isPlottedEdge(nodeA, nodeB) {
        return
    }
    dir := "forward"
    if dw.Reversed {
        dir = "back"
    }
    dw.printFormat("\t%s -> %s[dir=%s]\n", escape(nodeA.Name()), escape(nodeB.Name()), dir)
}

func (dw *DotWriter) printLine(str string) {
    fmt.Fprintln(dw.output, str)
}

func (dw *DotWriter) printFormat(pattern string, args ...interface{}) {
    fmt.Fprintf(dw.output, pattern, args...)
}

func escape(target string) string {
    return strconv.Quote(target)
}
