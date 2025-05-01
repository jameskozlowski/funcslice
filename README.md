# FuncSlice

`FuncSlice` is a generic wrapper around Go slices that provides functional-style operations such as `Map`, `Filter`, `Reduce`, and more. It is designed to make working with slices in Go more expressive and convenient.

## Features

- **Map**: Apply a function to each element of the slice.
- **Filter**: Filter elements based on a predicate.
- **Reduce**: Reduce the slice to a single value.
- **ForEach**: Perform an action for each element.
- **Sort**: Sort the slice using a custom comparison function.
- **Reverse**: Reverse the order of elements.
- **Get/Set**: Access or modify elements by index.
- **Append/Prepend/Insert/Remove**: Modify the slice dynamically.
- **Contains**: Check if the slice contains a specific value.
- **IndexOf/LastIndexOf**: Find the index of a value.
- **Join**: Concatenate elements into a string.
- **Clone**: Create a deep copy of the slice.
- **Length**: Get the number of elements.
- **Clear**: Remove all elements.
- **First/Last**: Access the first or last element.

## Installation

To use `FuncSlice`, you need Go 1.18 or later (for generics support). Install the package using:

```bash
go get github.com/yourusername/funcslice
```

## Useage

```go
package main

import (
    "fmt"
    "github.com/yourusername/funcslice"
)

func main() {
    // Create a new FuncSlice
    fs := funcslice.NewFuncSlice([]int{1, 2, 3, 4, 5})

    // Map: Multiply each element by 2
    mapped := fs.Map(func(x int) int { return x * 2 })
    fmt.Println("Mapped:", mapped.ToSlice())

    // Filter: Keep only even numbers
    filtered := fs.Filter(func(x int) bool { return x%2 == 0 })
    fmt.Println("Filtered:", filtered.ToSlice())

    // Reduce: Sum all elements
    sum := fs.Reduce(func(a, b int) int { return a + b }, 0)
    fmt.Println("Sum:", sum)

    // Sort: Sort in descending order
    sorted := fs.Sort(func(a, b int) bool { return a > b })
    fmt.Println("Sorted:", sorted.ToSlice())

    // Reverse: Reverse the slice
    reversed := fs.Reverse()
    fmt.Println("Reversed:", reversed.ToSlice())
}
```