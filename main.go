package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/ozx1812/go-mapreduce/mapreduce"
	"github.com/ozx1812/go-mapreduce/mrapps"
)

// for sorting by key
type ByKey []mapreduce.KeyValue

func (a ByKey) Len() int {return len(a)}
func (a ByKey) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a ByKey) Less(i, j int) bool {return a[i].Key < a[j].Key}


func main()  {
    text := "Apple Banana Banana Cat Cat Cat "
    for i := 0; i < 20; i++ {
        fmt.Printf("\n------\nIteration i: %v\n", i)
        start := time.Now()
        text += text
        var wc mrapps.WordCount
        intermediate := []mapreduce.KeyValue{}
        kva := wc.Map(strconv.Itoa(i+1), text)
        intermediate = append(intermediate, kva...)

        // Sort by Key
        sort.Sort(ByKey(intermediate))
    
        // Reduce on distinct keys
        i := 0
        for i < len(intermediate) {
            j := i + 1
            for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
                j++
            }
            values := []string{}
            for k := i; k<j; k++{
                values = append(values, intermediate[k].Value)
            }
            output := wc.Reduce(intermediate[i].Key, values)
            fmt.Printf("%v %v\n", intermediate[i].Key, output)
            i = j
        }
        elapsed := time.Since(start)
        fmt.Printf("Time : %s", elapsed)
    }
}
