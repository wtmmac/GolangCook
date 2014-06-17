package SortedLinkedList 
import (
    "container/list"
)
type SortedLinkedList struct {
    *list.List
    Limit int
    compareFunc func (old, new interface{}) bool
}
func NewSortedLinkedList(limit int, compare func (old, new interface{}) bool) *SortedLinkedList {
    return &SortedLinkedList{list.New(), limit, compare}
}
func (this SortedLinkedList) findInsertPlaceElement(value interface{}) *list.Element {
    for element := this.Front(); element != nil; element = element.Next() {
        tempValue := element.Value
        if this.compareFunc(tempValue, value) {
            return element
        }
    }
    return nil
}
func (this SortedLinkedList) PutOnTop(value interface{}) {
    if this.List.Len() == 0 {
        this.PushFront(value)
        return
    }
    if this.List.Len() < this.Limit && this.compareFunc(value, this.Back().Value) {
        this.PushBack(value)
        return
    }
    if this.compareFunc(this.List.Front().Value, value) {
        this.PushFront(value)
    } else if this.compareFunc(this.List.Back().Value, value) && this.compareFunc(value, this.Front().Value) {
        element := this.findInsertPlaceElement(value)
        if element != nil {
            this.InsertBefore(value, element)
        }
    }
    if this.Len() > this.Limit {
        this.Remove(this.Back())
    }
}
