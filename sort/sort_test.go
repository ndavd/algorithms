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

func TestSelection(t *testing.T) {
  var test []int
  test = append(test, unsorted...)

  Selection(test)
  if !areEqual(test, sorted) {
    t.Errorf("Selection() FAILED. Expected %v, got %v\n", sorted, test)
  }
}

func TestInsertion(t *testing.T) {
  var test []int
  test = append(test, unsorted...)

  Insertion(test)
  if !areEqual(test, sorted) {
    t.Errorf("Insertion() FAILED. Expected %v, got %v\n", sorted, test)
  }
}

func TestBubble(t *testing.T) {
  var test []int
  test = append(test, unsorted...)

  Bubble(test)
  if !areEqual(test, sorted) {
    t.Errorf("Bubble() FAILED. Expected %v, got %v\n", sorted, test)
  }
}

func TestMergesort(t *testing.T) {
  var test []int
  test = append(test, unsorted...)

  Mergesort(test)
  if !areEqual(test, sorted) {
    t.Errorf("Mergesort() FAILED. Expected %v, got %v\n", sorted, test)
  }
}

func TestQuicksort(t *testing.T) {
  var test []int
  test = append(test, unsorted...)

  Quicksort(test)
  if !areEqual(test, sorted) {
    t.Errorf("Quicksort() FAILED. Expected %v, got %v\n", sorted, test)
  }
}
