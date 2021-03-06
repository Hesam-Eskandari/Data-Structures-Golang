package Stacks

import (
	"github.com/Data-Structures-Golang/pkg/utils"
	"reflect"
	"sort"
)

type (
	node struct {
		value interface{}
		prev  *node
	}
	stack struct {
		top    *node
		length int
		min    *node
	}
)

type Stack interface {
	// Append appends a second stack on top of the original stack
	Append(stack *stack) error
	// AppendArray pushes the item values of the array starting from index zero to the stack
	AppendArray(array interface{})
	// AppendReverse appends a second stack from top to end to the top of the primary stack
	AppendReverse(stack *stack) error
	// Len returns the number of nodes in the stack
	Len() int
	// Pop removes the top node and returns its value. It also assigns the previous node as the new top
	Pop() (detachedHead interface{})
	// Push adds a node with given value to the stack. Also moves the top to the new added node
	Push(value interface{}) interface{}
	// Reverse returns a new stack that has values of the original stack reversed
	Reverse() *stack
	// Sort returns a non-decreasing sorted copy of the primary stack
	// runs in O(n log(n)) in time and O(n) of space complexity where n = s.Len()
	Sort() *stack
	// SortN replace a stack with a non-decreasing sorted version of it
	// only one additional stack is allowed to be used
	// only push, pop, top and len methods are allowed to be used
	// runs in O(n) of space and O(n^2) of time complexity where n = stack.Len()
	// assume that the values are of type int
	SortN()
	// Top returns the value of the top of the stack
	Top() interface{}
	// ToArray returns an array of values in the stack. The last index of array corresponds to the top of the stack
	ToArray() (arr []interface{})
}

// NewStack constructs and returns a new empty stack
func NewStack() *stack {
	return &stack{
		top:    nil,
		length: 0,
		min:    nil,
	}
}

// Top returns the value of the top of the stack
func (s *stack) Top() interface{} {
	if s.top == nil {
		return nil
	}
	return s.top.value
}

// Len returns the number of nodes in the stack
func (s *stack) Len() int {
	return s.length
}

// Push adds a node with given value to the stack. Also moves the top to the new added node
func (s *stack) Push(value interface{}) interface{} {
	top := &node{
		value: value,
		prev:  s.top,
	}
	s.top = top
	s.length += 1
	s.calcMin(value, false)
	return s.top.value
}

// Pop removes the top node and returns its value. It also assign the previous node as the new top
func (s *stack) Pop() (detachedHead interface{}) {
	if s.top == nil {

		panic(StackTopIsNilException{"Pop"})
	}
	detachedHead = s.top.value
	if s.top.prev == nil {
		s.top = nil
	} else {
		s.top = s.top.prev
	}
	s.length -= 1
	s.calcMin(detachedHead, true)
	return
}

// ToArray returns an array of values in the stack. The last index of array corresponds to the top of the stack
func (s *stack) ToArray() (arr []interface{}) {
	arr = make([]interface{}, s.length)
	node := s.top
	for index := s.length - 1; index >= 0; index-- {
		arr[index] = node.value
		node = node.prev
	}
	return
}

// AppendArray pushes the item values of the array starting from index zero to the stack
func (s *stack) AppendArray(array interface{}) {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arr := reflect.ValueOf(array)
		for index := 0; index < arr.Len(); index++ {
			s.Push(arr.Index(index).Interface())
		}
	}
}

// Reverse returns a new stack that has values of the original stack reversed
func (s *stack) Reverse() *stack {
	stack := NewStack()
	for node := s.top; node != nil; {
		stack.Push(node.value)
		node = node.prev
	}
	return stack
}

// AppendReverse appends a second stack from top to end to the top of the primary stack
func (s *stack) AppendReverse(stack *stack) error {
	if s == nil {
		return &utils.StructIsNilException{DataStructure: "AppendReverse", FuncName: "Stack"}
	}
	if stack == nil {
		return &utils.StructIsNilException{DataStructure: "AppendReverse", FuncName: "Stack"}
	}
	for node := stack.top; node != nil; {
		s.Push(node.value)
		node = node.prev
	}
	return nil
}

// Append appends a second stack on top of the original stack
func (s *stack) Append(stack *stack) error {
	return s.AppendReverse(stack.Reverse())
}

// calcMin updates the min linked list if necessary
func (s *stack) calcMin(value interface{}, isPop bool) {
	var valueType interface{}
	if s.min != nil {
		valueType = s.min.value
	} else if value != nil {
		s.min = &node{
			value: value,
		}
		return
	} else {
		return
	}
	if isPop && value == s.min.value {
		s.min = s.min.prev
		return
	}
	switch valueType.(type) {
	case int:
		if value.(int) <= s.min.value.(int) {
			s.pushMin(value)
		}
	case int8:
		if value.(int8) <= s.min.value.(int8) {
			s.pushMin(value)
		}
	case int16:
		if value.(int16) <= s.min.value.(int16) {
			s.pushMin(value)
		}
	case int32:
		if value.(int32) <= s.min.value.(int32) {
			s.pushMin(value)
		}
	case int64:
		if value.(int64) <= s.min.value.(int64) {
			s.pushMin(value)
		}
	case float32:
		if value.(float32) <= s.min.value.(float32) {
			s.pushMin(value)
		}
	case float64:
		if value.(float64) <= s.min.value.(float64) {
			s.pushMin(value)
		}
	default:
		return
	}

}

// pushMin pushes a new minimum to the min singly linked list head (top)
func (s *stack) pushMin(value interface{}) {
	if s.min == nil {
		s.min = &node{value, nil}
	} else {
		node := &node{
			value: value,
			prev:  s.min,
		}
		s.min = node
	}
}

// SortN replace a stack with a non-decreasing sorted version of it
// only one additional stack is allowed to be used
// only push, pop, top and len methods are allowed to be used
// runs in O(n) of space and O(n^2) of time complexity where n = stack.Len()
// assume that the values are of type int
func (s *stack) SortN() {
	if s.Top() == nil {
		return
	}
	switch s.Top().(type) {
	case int:
	default:
		return
	}
	stack := NewStack()
	for s.Top() != nil {
		if stack.Top() == nil {
			stack.Push(s.Pop())
			continue
		}
		value := s.Top().(int)
		count := 0
		for stack.Top() != nil && stack.Top().(int) < value {
			s.Push(stack.Pop())
			count += 1
		}
		stack.Push(value)
		for ; count > 0; count-- {
			stack.Push(s.Pop())
		}
		s.Pop()
	}
	for stack.Top() != nil {
		s.Push(stack.Pop())
	}
}

// Sort returns a non-decreasing sorted copy of the primary stack
// runs in O(n log(n)) in time and O(n) of space complexity where n = s.Len()
func (s *stack) Sort() *stack {
	if s.Top() == nil {
		return nil
	}
	stack := s
	switch s.Top().(type) {
	case int:
		arr := make([]int, 0, s.Len())
		for stack.Top() != nil {
			arr = append(arr, stack.Pop().(int))
		}
		sort.Ints(arr)
		stack.AppendArray(arr)
	case float64:
		arr := make([]float64, 0, s.Len())
		for stack.Top() != nil {
			arr = append(arr, stack.Pop().(float64))
		}
		sort.Float64s(arr)
		stack.AppendArray(arr)
	default:
		return nil
	}
	return stack
}
