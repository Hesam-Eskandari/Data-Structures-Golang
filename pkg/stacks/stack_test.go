package Stacks

import (
	"fmt"
	"reflect"
	"testing"
)

func (s *Stack) assertEqualArray(t *testing.T, array interface{}) {
	if s == nil {
		panic("assertEqualArray: expected a stack, received nil")
	}
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		top := s.top
		if s.length != arr.Len() {
			t.Errorf("assertEqualArray error: stack and array don't have the same length")
		} else {
			for index := arr.Len() - 1; index >= 0; index-- {

				if top.value != arr.Index(index).Interface() {
					t.Errorf("assertEqualArray error: the element %v of stack with value %v is not "+
						"equal to item %v from array with value %v", arr.Len()-index, top.value, index, arr.Index(index))
					break
				}
				top = top.prev
			}
		}
	default:
		panic(fmt.Sprintf("expected an array, received %v", reflect.TypeOf(array).Kind()))
	}

}

func (s *Stack) assertEqual(t *testing.T, stack *Stack) {
	if s == nil || stack == nil {
		panic("assertEqual: expected a stack, received nil")
	}
	if s.length != stack.length {
		t.Errorf("assertEqual: two stacks don't have the same length. \n"+
			"stack 1: %v \n stack 2: %v", s.ToArray(), stack.ToArray())
	}
}

func (s *Stack) assertEqualTop(t *testing.T, value interface{}) {
	if s == nil {
		panic("assertEqualTop: expected a stack, received nil")
	}
	if s.top != nil && value != s.top.value {
		t.Errorf("assertEqualTop: expected the top value to be %v but"+
			" it's equal to %v", value, s.top.value)
	}
	if s.top == nil && value != nil {
		t.Errorf("assertEqualTop: expected the top value to be %v but"+
			" it's equal to %v", value, s.top.value)
	}
}

func (s *Stack) assertLength(t *testing.T, length int) {
	if s == nil {
		panic("assertEqualArray: expected a stack, received nil")
	}
	if s.length != length {
		t.Errorf("expected the stack to be of length %v, but its length is %v", length, s.length)
	}
}

func setUpStack() (stack *Stack, expectedArray []int) {
	stack = newStack()
	expectedArray = []int{1}
	stack.AppendArray(expectedArray)
	return
}

func TestStack_Array(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.assertEqualArray(t, expectedArray)

}

func TestStack_Len(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.assertLength(t, len(expectedArray))
}

func TestStack_Top(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.assertEqualTop(t, expectedArray[len(expectedArray)-1])
}

func TestStack_Push(t *testing.T) {
	stack, expectedArray := setUpStack()
	add := 999
	stack.Push(add)
	stack.assertEqualTop(t, add)
	stack.assertLength(t, len(expectedArray)+1)
	stack.assertEqualArray(t, append(expectedArray, add))
}

func TestStack_Pop(t *testing.T) {
	stack, expectedArray := setUpStack()
	stack.Pop()
	var newTopValue interface{}
	if len(expectedArray) > 1 {
		newTopValue = expectedArray[len(expectedArray)-2]
	} else {
		newTopValue = nil
	}
	stack.assertEqualTop(t, newTopValue)
}
