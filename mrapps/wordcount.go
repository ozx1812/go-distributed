package mrapps

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/ozx1812/go-mapreduce/mapreduce"
)

type WordCount struct{}

func (wc WordCount) Map(key, content string) []mapreduce.KeyValue {
    // function to detect the word separators
    ff := func(r rune) bool {return !unicode.IsLetter(r)}
    
    // split value into words
    words := strings.FieldsFunc(content, ff)
    keyValues := make([]mapreduce.KeyValue, len(words))
    for i, word := range words {
        keyValues[i] = mapreduce.KeyValue{Key: word, Value: "1"}
    }
    return keyValues
}

func (wc WordCount) Reduce(key string, values []string) string {
    return strconv.Itoa(len(values))
}
