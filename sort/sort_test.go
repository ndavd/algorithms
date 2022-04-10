package sort

import (
  "math/rand"
  "testing"
  "time"
)

/* UNIT TESTING */

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

/* BENCHMARK TESTING */

// Check if slice is sorted
func isSorted[T Ordered](a []T) bool {
  for i := 0; i < len(a)-1; i++ {
    if a[i + 1] < a[i] {
      return false
    }
  }
  return true
}

func benchmarkAlgorithm(b *testing.B, a func([]int), name string) {
  for i := 0; i < b.N; i++ {
    var test []int
    rand.Seed(time.Now().Unix())
    for i := 0; i < 10000; i++ {
      test = append(test, rand.Intn(200) - 100)
    }
    a(test)
    if !isSorted(test) {
      b.Errorf("%v() FAILED. Not properly sorted!\n", name)
    }
  }
}

func BenchmarkSelection(b *testing.B) {
  benchmarkAlgorithm(b, Selection[int], "Selection")
}

func BenchmarkInsertion(b *testing.B) {
  benchmarkAlgorithm(b, Insertion[int], "Insertion")
}

func BenchmarkBubble(b *testing.B) {
  benchmarkAlgorithm(b, Bubble[int], "Bubble")
}

func BenchmarkMergesort(b *testing.B) {
  benchmarkAlgorithm(b, Mergesort[int], "Mergesort")
}

func BenchmarkQuicksort(b *testing.B) {
  benchmarkAlgorithm(b, Quicksort[int], "Quicksort")
}
