package funcslice

import "fmt"

// FuncSlice is a generic wrapper around a slice of type T, providing
// functional-style operations such as Map, Filter, and Reduce.
type FuncSlice[T any] struct {
	data []T
}

// NewFuncSlice creates a new FuncSlice from the given slice of type T.
func NewFuncSlice[T any](data []T) *FuncSlice[T] {
	return &FuncSlice[T]{data: data}
}

// Map applies the given function to each element of the slice and returns
// a new FuncSlice containing the results.
func (fs *FuncSlice[T]) Map(f func(T) T) *FuncSlice[T] {
	result := make([]T, len(fs.data))
	for i, v := range fs.data {
		result[i] = f(v)
	}
	return NewFuncSlice(result)
}

// Filter returns a new FuncSlice containing only the elements that satisfy
// the given predicate function.
func (fs *FuncSlice[T]) Filter(f func(T) bool) *FuncSlice[T] {
	result := make([]T, 0)
	for _, v := range fs.data {
		if f(v) {
			result = append(result, v)
		}
	}
	return NewFuncSlice(result)
}

// Reduce reduces the elements of the slice to a single value by repeatedly
// applying the given function, starting with the initial value.
func (fs *FuncSlice[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for _, v := range fs.data {
		result = f(result, v)
	}
	return result
}

// ForEach applies the given function to each element of the slice.
func (fs *FuncSlice[T]) ForEach(f func(T)) {
	for _, v := range fs.data {
		f(v)
	}
}

// Sort returns a new FuncSlice with the elements sorted according to the
// given comparison function.
func (fs *FuncSlice[T]) Sort(less func(T, T) bool) *FuncSlice[T] {
	result := make([]T, len(fs.data))
	copy(result, fs.data)
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if less(result[j], result[i]) {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return NewFuncSlice(result)
}

// Reverse returns a new FuncSlice with the elements in reverse order.
func (fs *FuncSlice[T]) Reverse() *FuncSlice[T] {
	result := make([]T, len(fs.data))
	for i, v := range fs.data {
		result[len(fs.data)-1-i] = v
	}
	return NewFuncSlice(result)
}

// ToSlice returns the underlying slice of type T.
func (fs *FuncSlice[T]) ToSlice() []T {
	return fs.data
}

// Length returns the number of elements in the slice.
func (fs *FuncSlice[T]) Length() int {
	return len(fs.data)
}

// Get returns the element at the specified index. Panics if the index is out of range.
func (fs *FuncSlice[T]) Get(index int) T {
	if index < 0 || index >= len(fs.data) {
		panic("index out of range")
	}
	return fs.data[index]
}

// Set sets the element at the specified index to the given value. Panics if the index is out of range.
func (fs *FuncSlice[T]) Set(index int, value T) {
	if index < 0 || index >= len(fs.data) {
		panic("index out of range")
	}
	fs.data[index] = value
}

// Append adds the given value to the end of the slice.
func (fs *FuncSlice[T]) Append(value T) {
	fs.data = append(fs.data, value)
}

// Prepend adds the given value to the beginning of the slice.
func (fs *FuncSlice[T]) Prepend(value T) {
	fs.data = append([]T{value}, fs.data...)
}

// Insert inserts the given value at the specified index. Panics if the index is out of range.
func (fs *FuncSlice[T]) Insert(index int, value T) {
	if index < 0 || index > len(fs.data) {
		panic("index out of range")
	}
	fs.data = append(fs.data[:index], append([]T{value}, fs.data[index:]...)...)
}

// Remove removes the element at the specified index. Panics if the index is out of range.
func (fs *FuncSlice[T]) Remove(index int) {
	if index < 0 || index >= len(fs.data) {
		panic("index out of range")
	}
	fs.data = append(fs.data[:index], fs.data[index+1:]...)
}

// Clear removes all elements from the slice.
func (fs *FuncSlice[T]) Clear() {
	fs.data = []T{}
}

// Contains checks if the slice contains the given value, using the provided equality function.
func (fs *FuncSlice[T]) Contains(value T, equals func(T, T) bool) bool {
	for _, v := range fs.data {
		if equals(v, value) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the given value in the slice,
// using the provided equality function. Returns -1 if the value is not found.
func (fs *FuncSlice[T]) IndexOf(value T, equals func(T, T) bool) int {
	for i, v := range fs.data {
		if equals(v, value) {
			return i
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the given value in the slice,
// using the provided equality function. Returns -1 if the value is not found.
func (fs *FuncSlice[T]) LastIndexOf(value T, equals func(T, T) bool) int {
	for i := len(fs.data) - 1; i >= 0; i-- {
		if equals(fs.data[i], value) {
			return i
		}
	}
	return -1
}

// Join concatenates the string representations of the elements in the slice,
// separated by the given separator.
func (fs *FuncSlice[T]) Join(separator string) string {
	result := ""
	for i, v := range fs.data {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", v)
	}
	return result
}

// String returns a string representation of the slice.
func (fs *FuncSlice[T]) String() string {
	result := "["
	for i, v := range fs.data {
		if i > 0 {
			result += ", "
		}
		result += fmt.Sprintf("%v", v)
	}
	result += "]"
	return result
}

// Clone creates a deep copy of the FuncSlice and returns it.
func (fs *FuncSlice[T]) Clone() *FuncSlice[T] {
	result := make([]T, len(fs.data))
	copy(result, fs.data)
	return NewFuncSlice(result)
}

// First returns the first element of the slice. Panics if the slice is empty.
func (fs *FuncSlice[T]) First() T {
	if len(fs.data) == 0 {
		panic("slice is empty")
	}
	return fs.data[0]
}

// Last returns the last element of the slice. Panics if the slice is empty.
func (fs *FuncSlice[T]) Last() T {
	if len(fs.data) == 0 {
		panic("slice is empty")
	}
	return fs.data[len(fs.data)-1]
}
