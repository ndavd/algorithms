package sort

import (
  "testing"
)

// Check slice equality
func areEqual[T Ordered](a []T, b []T) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

var unsorted = []int{
  -11, -34, -30, 36, -23, 8, -69, 39, -27, -41,
  -24, 31, 73, -4, 80, -46, 15, -42, -100, -77,
}
var sorted = []int{
  -100, -77, -69, -46, -42, -41, -34, -30, -27, -24,
  -23, -11, -4, 8, 15, 31, 36, 39, 73, 80,
}

func testAlgorithm(t *testing.T, a func([]int), name string) {
  var test []int
  test = append(test, unsorted...)
  a(test)
  if !areEqual(test, sorted) {
    t.Errorf("%v() FAILED. Expected %v, got %v.\n", name, sorted, test)
  }
}

func TestSelection(t *testing.T) {
  testAlgorithm(t, Selection[int], "Selection")
}

func TestInsertion(t *testing.T) {
  testAlgorithm(t, Insertion[int], "Insertion")
}

func TestBubble(t *testing.T) {
  testAlgorithm(t, Bubble[int], "Bubble")
}

func TestMergesort(t *testing.T) {
  testAlgorithm(t, Mergesort[int], "Mergesort")
}

func TestQuicksort(t *testing.T) {
  testAlgorithm(t, Quicksort[int], "Quicksort")
}
