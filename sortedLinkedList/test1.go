package main
import (
    "fmt"
    "GolangTest/SortedLinkedList"
)
type WordCount struct {
    Word  string
    Count int
}
func compareValue(old, new interface {}) bool {
    if new.(WordCount).Count > old.(WordCount).Count {
        return true
    }
    return false
}
func main() {
    wordCounts := []WordCount{
        WordCount{"kate", 87},
        WordCount{"herry", 92},
        WordCount{"james", 81}}
    var aSortedLinkedList = SortedLinkedList.NewSortedLinkedList(10, compareValue)
    for _, wordCount := range wordCounts {
        aSortedLinkedList.PutOnTop(wordCount)
    }
    for element := aSortedLinkedList.List.Front(); element != nil; element = element.Next() {
        fmt.Println(element.Value.(WordCount))
    }
}
