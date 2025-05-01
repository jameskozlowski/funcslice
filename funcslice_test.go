package funcslice

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	mapped := fs.Map(func(x int) int { return x * 2 }).ToSlice()
	expectedMapped := []int{2, 4, 6, 8, 10}
	if fmt.Sprintf("%v", mapped) != fmt.Sprintf("%v", expectedMapped) {
		t.Errorf("Map failed: got %v, want %v", mapped, expectedMapped)
	}
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	filtered := fs.Filter(func(x int) bool { return x%2 == 0 }).ToSlice()
	expectedFiltered := []int{2, 4}
	if fmt.Sprintf("%v", filtered) != fmt.Sprintf("%v", expectedFiltered) {
		t.Errorf("Filter failed: got %v, want %v", filtered, expectedFiltered)
	}
}

func TestReduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	reduced := fs.Reduce(func(a, b int) int { return a + b }, 0)
	expectedReduced := 15
	if reduced != expectedReduced {
		t.Errorf("Reduce failed: got %v, want %v", reduced, expectedReduced)
	}
}

func TestForEach(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	sum := 0
	fs.ForEach(func(x int) { sum += x })
	expectedSum := 15
	if sum != expectedSum {
		t.Errorf("ForEach failed: got %v, want %v", sum, expectedSum)
	}
}

func TestSort(t *testing.T) {
	fs := NewFuncSlice([]int{3, 1, 4, 1, 5})

	sorted := fs.Sort(func(a, b int) bool { return a < b }).ToSlice()
	expectedSorted := []int{1, 1, 3, 4, 5}
	if fmt.Sprintf("%v", sorted) != fmt.Sprintf("%v", expectedSorted) {
		t.Errorf("Sort failed: got %v, want %v", sorted, expectedSorted)
	}
}

func TestReverse(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	reversed := fs.Reverse().ToSlice()
	expectedReversed := []int{5, 4, 3, 2, 1}
	if fmt.Sprintf("%v", reversed) != fmt.Sprintf("%v", expectedReversed) {
		t.Errorf("Reverse failed: got %v, want %v", reversed, expectedReversed)
	}
}

func TestLength(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	if fs.Length() != len(data) {
		t.Errorf("Length failed: got %v, want %v", fs.Length(), len(data))
	}
}

func TestGet(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	if fs.Get(0) != data[0] {
		t.Errorf("Get failed: got %v, want %v", fs.Get(0), data[0])
	}
}

func TestSet(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Set(0, 10)
	if fs.Get(0) != 10 {
		t.Errorf("Set failed: got %v, want %v", fs.Get(0), 10)
	}
}

func TestAppend(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Append(6)
	if fs.Last() != 6 {
		t.Errorf("Append failed: got %v, want %v", fs.Last(), 6)
	}
}

func TestPrepend(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Prepend(0)
	if fs.First() != 0 {
		t.Errorf("Prepend failed: got %v, want %v", fs.First(), 0)
	}
}

func TestInsert(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Insert(1, 99)
	if fs.Get(1) != 99 {
		t.Errorf("Insert failed: got %v, want %v", fs.Get(1), 99)
	}
}

func TestRemove(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Insert(1, 99)
	fs.Remove(1)
	if fs.Get(1) == 99 {
		t.Errorf("Remove failed: element not removed")
	}
}

func TestClear(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Clear()
	if fs.Length() != 0 {
		t.Errorf("Clear failed: got %v, want %v", fs.Length(), 0)
	}
}

func TestContains(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	equal := func(a, b int) bool { return a == b }
	if !fs.Contains(3, equal) {
		t.Errorf("Contains failed: expected to find 3")
	}
}

func TestIndexOf(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	equal := func(a, b int) bool { return a == b }
	if fs.IndexOf(3, equal) != 2 {
		t.Errorf("IndexOf failed: got %v, want %v", fs.IndexOf(3, equal), 2)
	}
}

func TestLastIndexOf(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	equal := func(a, b int) bool { return a == b }
	fs.Append(3)
	if fs.LastIndexOf(3, equal) != 5 {
		t.Errorf("LastIndexOf failed: got %v, want %v", fs.LastIndexOf(3, equal), 5)
	}
}

func TestJoin(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Append(3)
	joined := fs.Join(",")
	expectedJoined := "1,2,3,4,5,3"
	if joined != expectedJoined {
		t.Errorf("Join failed: got %v, want %v", joined, expectedJoined)
	}
}

func TestString(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Append(3)
	str := fs.String()
	expectedStr := "[1, 2, 3, 4, 5, 3]"
	if str != expectedStr {
		t.Errorf("String failed: got %v, want %v", str, expectedStr)
	}
}

func TestClone(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	cloned := fs.Clone()
	if fmt.Sprintf("%v", cloned.ToSlice()) != fmt.Sprintf("%v", fs.ToSlice()) {
		t.Errorf("Clone failed: got %v, want %v", cloned.ToSlice(), fs.ToSlice())
	}
}

func TestFirst(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	if fs.First() != 1 {
		t.Errorf("First failed: got %v, want %v", fs.First(), 1)
	}
}

func TestLast(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	fs := NewFuncSlice(data)

	fs.Append(3)
	if fs.Last() != 3 {
		t.Errorf("Last failed: got %v, want %v", fs.Last(), 3)
	}
}
