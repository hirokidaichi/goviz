package metrics

import (
    "fmt"
    "github.com/hirokidaichi/goviz/dotwriter"
    "io"
    "sort"
)

type MetricsWriter struct {
    output io.Writer
}

type element struct {
    Name string
    Inst float32
    Ca   int
    Ce   int
}
type elementArray []*element

type elementArraySorter struct {
    Elements elementArray
    SortBy   string
}

func (self elementArray) SortBy(name string) elementArray {
    sorter := &elementArraySorter{
        Elements: self,
        SortBy:   name,
    }
    sort.Sort(sorter)
    return self
}

func (sorter *elementArraySorter) Len() int {
    return len(sorter.Elements)
}

func (sorter *elementArraySorter) Less(i, j int) bool {
    elements := sorter.Elements
    ei, ej := elements[i], elements[j]
    if sorter.SortBy == "Name" {
        return ei.Name < ej.Name
    }
    if sorter.SortBy == "Inst" {
        return ei.Inst > ej.Inst
    }
    return false
}
func (sorter *elementArraySorter) Swap(i, j int) {

    sorter.Elements[i], sorter.Elements[j] = sorter.Elements[j], sorter.Elements[i]
}

func New(output io.Writer) *MetricsWriter {
    return &MetricsWriter{output: output}
}

func (mw *MetricsWriter) Plot(nodes []dotwriter.IDotNode) {
    list := make(elementArray, len(nodes))
    for idx, v := range nodes {
        ca := afferentCoupling(v)
        ce := efferentCoupling(v)
        list[idx] = &element{Name: v.Name(), Inst: float32(ce) / (float32(ca) + float32(ce)), Ca: ca, Ce: ce}
    }
    list.SortBy("Name").SortBy("Inst")

    for _, v := range list {
        fmt.Fprintf(mw.output, "Inst:%0.3f Ca(%3d) Ce(%3d)\t%s\n", v.Inst, v.Ca, v.Ce, v.Name)
    }

}

func afferentCoupling(node dotwriter.IDotNode) int {
    num := len(node.Parents())
    return num
}

func efferentCoupling(node dotwriter.IDotNode) int {
    num := len(node.Children())
    return num
}
