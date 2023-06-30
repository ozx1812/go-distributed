package main

import (
	"fmt"
	"strings"
	"time"
)

// Basic MapReduce Implementation
type keyvalue struct    {
    k, v string
}

type MapReduce interface  {
    mapper(key string, value string) []keyvalue
    reducer(key string, data []keyvalue) 
}

type WordCount struct {
    words map[string]int
}

func (wc WordCount) mapper(key string, value string) []keyvalue {
    var data []keyvalue
    words := strings.Split(value, " ") 
    for _, w := range words {
        data = append(data, keyvalue{w,"1"})
    }
    return data
}

func (wc WordCount) reducer(key string, data []keyvalue) {
    for _, w := range data {
        if wc.words[w.k] == 0 {
            wc.words[w.k] = 1
        }else {
            wc.words[w.k] += 1
        }
    }
}

func main()  {
    start := time.Now()
    fmt.Println("Hello main...")
    var wc WordCount
    wc.words = make(map[string]int)
    text := "hello hello how how how are are you you you"
    for i := 0; i < 20; i++ {
        text += text
    }
    data := wc.mapper("1", text)
    wc.reducer("1", data)
    fmt.Println(wc.words)
    elapsed := time.Since(start)
    fmt.Printf("time : %s", elapsed)

}
