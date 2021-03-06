package linkedlists

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

type linkedList struct {
	List *list.List
}

type LinkedList interface {
	// AssertEqualArray if a values and ordering of a linked list is equivalent to an array
	AssertEqualArray(t *testing.T, expectedArr []interface{})
	// AssertEqualElementValue checks if an element value is equal to what is expected
	AssertEqualElementValue(t *testing.T, element *list.Element, value int)
	// AssertNilElement checks if an element is nil
	AssertNilElement(t *testing.T, element *list.Element)
	// CreateListFromArray constructs a doubly linkedList from a given array
	CreateListFromArray(array interface{}) *linkedList
	// DeleteDuplicatesBack removes redundant nodes assuming it's a singly linked List with only head is given
	DeleteDuplicatesBack(noAdditionalDatastructures bool)
	// DeleteDuplicates removes redundant nodes assuming it's a singly linked List with only head is given
	DeleteDuplicates(noAdditionalDatastructures bool)
	// GetList gets the inner list
	GetList() *list.List
	// KthToLast returns the kth element from tail (back) assuming it's a singly linked List with given head
	KthToLast(k int) *list.Element
	// Replace replaces the linked list with a new linked list
	Replace(List *linkedList)
	// SetList replaces the List inside the linkedList with a new List
	SetList(list *list.List)
	// ToArray returns an array with values in linked List with same order and size
	ToArray() []interface{}
}

func NewLinkedList() LinkedList {
	return &linkedList{
		List: list.New(),
	}
}

// AssertEqualArray if a values and ordering of a linked list is equivalent to an array
func (l *linkedList) AssertEqualArray(t *testing.T, expectedArr []interface{}) {
	if l == nil {
		panic("AssertEqualArray: given linked list is nil")
	}
	arr := l.ToArray()
	if len(expectedArr) != len(arr) {
		t.Errorf("AssertEqualArray: linked list and expected array are not the same size")
	}
	for index, item := range arr {
		if item != expectedArr[index] {
			t.Errorf("returned linked list does not match to expected array, returned: %v, expected: %v",
				l.ToArray(), expectedArr)
			break
		}
	}
}

// AssertEqualElementValue checks if an element value is equal to what is expected
func (l *linkedList) AssertEqualElementValue(t *testing.T, element *list.Element, value int) {
	if element == nil {
		panic(fmt.Sprintf("AssertEqualElementValue: given element is nil, expected value: %v", value))
	}
	if element.Value != value {
		t.Errorf("expected element with value= %v received %v element", value, element.Value)
	}
}

// AssertNilElement checks if an element is nil
func (l *linkedList) AssertNilElement(t *testing.T, element *list.Element) {
	if element != nil {
		t.Errorf("expected nil returned %v", element.Value)
	}
}

// CreateListFromArray constructs a doubly linkedList from a given array
func (l *linkedList) CreateListFromArray(array interface{}) *linkedList {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		for index := 0; index < arr.Len(); index++ {
			l.List.PushBack(arr.Index(index).Interface())
		}
	}
	return l
}

// DeleteDuplicates removes redundant nodes assuming it's a singly linked List with only front is given
func (l *linkedList) DeleteDuplicates(noAdditionalDatastructures bool) {
	element := l.List.Front()
	if element == nil {
		return
	}
	if noAdditionalDatastructures {
		// O(n^2) time, O(1) space
		for element.Next() != nil {
			node := element
			for node.Next() != nil {
				if node.Next().Value == element.Value {
					l.List.Remove(node.Next())
				} else {
					node = node.Next()
				}
			}
			element = element.Next()
			if element == nil {
				break
			}
		}
	} else {
		// O(n) time, O(n) space
		hMap := make(map[interface{}]bool)
		hMap[element.Value] = true
		for element.Next() != nil {
			if _, ok := hMap[element.Next().Value]; ok {
				l.List.Remove(element.Next())
			} else {
				hMap[element.Next().Value] = true
				element = element.Next()
			}
		}
	}
}

// DeleteDuplicatesBack removes redundant nodes assuming it's a singly linked List with only back is given
func (l *linkedList) DeleteDuplicatesBack(noAdditionalDatastructures bool) {
	element := l.List.Back()
	if element == nil {
		return
	}
	if noAdditionalDatastructures {
		// O(n^2) time, O(1) space
		for element.Prev() != nil {
			node := element
			for node.Prev() != nil {
				if node.Prev().Value == element.Value {
					l.List.Remove(node.Prev())
				} else {
					node = node.Prev()
				}
			}
			element = element.Prev()
			if element == nil {
				break
			}
		}
	} else {
		// O(n) time, O(n) space
		hMap := make(map[interface{}]bool)
		hMap[element.Value] = true
		for element.Prev() != nil {
			if _, ok := hMap[element.Prev().Value]; ok {
				l.List.Remove(element.Prev())
			} else {
				hMap[element.Prev().Value] = true
				element = element.Prev()
			}
		}
	}
}

// GetList gets the inner list
func (l *linkedList) GetList() *list.List {
	return l.List
}

// KthToLast returns the kth element from tail (back) assuming it's a singly linked List with given head
func (l *linkedList) KthToLast(k int) *list.Element {
	elementAhead := l.List.Front()
	elementDelay := l.List.Front()
	if elementAhead == nil || k < 0 {
		return nil
	}
	if k == 0 {
		k = 1
	}
	count := 0
	for elementAhead != nil {
		count += 1
		if count > k {
			elementDelay = elementDelay.Next()
		}
		elementAhead = elementAhead.Next()
	}
	if count < k {
		return nil
	}
	return elementDelay
}

// Replace replaces the List inside the linkedList with a new List
func (l *linkedList) Replace(List *linkedList) {
	l.List = List.List
}

// SetList replaces the list with a given list
func (l *linkedList) SetList(list *list.List) {
	l.List = list
}

// ToArray returns an array with values in linked List with same order and size
func (l *linkedList) ToArray() []interface{} {
	var arr []interface{}
	if l == nil {
		return arr
	}
	element := l.List.Front()
	for element != nil {
		arr = append(arr, element.Value)
		element = element.Next()
	}
	return arr
}
