package sort

import (
  "golang.org/x/exp/constraints"
)

type Ordered interface {
  constraints.Ordered
}

func swap[T Ordered](a *T, b *T) {
  temp := *a
  *a = *b
  *b = temp
}

func Selection[T Ordered](s []T) {
  for i := range s {
    current := i
    for j := i+1; j < len(s); j++ {
      if s[j] < s[current] {
        current = j
      }
    }
    if current != i {
      swap(&s[i], &s[current])
    }
  }
}

func Insertion[T Ordered](s []T) {
  for i := 1; i < len(s); i++ {
    current := s[i]
    j := i - 1
    for j >= 0 && s[j] > current {
      s[j+1] = s[j]
      j = j - 1
    }
    s[j+1] = current
  }
}

func Bubble[T Ordered](s []T) {
  swapped := false
  out:
  for i := range s {
    for j := 0; j < len(s)-i-1; j++ {
      if s[j] > s[j+1] {
        swap(&s[j], &s[j+1])
        swapped = true
      }
    }
    if !swapped {
      break out
    }
  }
}

func mergeSorted[T Ordered](s []T, l int, m int, r int) {
  lLen := m - l + 1
  rLen := r - m
  tempL, tempR := make([]T, lLen), make([]T, rLen)
  for i := 0; i < lLen; i++ {
    tempL[i] = s[l + i]
  }
  for i := 0; i < rLen; i++ {
    tempR[i] = s[m + 1 + i]
  }
  a, b := 0, 0
  for i := l; i <= r; i++ {
    if a < lLen && (b >= rLen || tempL[a] <= tempR[b]) {
      s[i] = tempL[a]
      a++
    } else {
      s[i] = tempR[b]
      b++
    }
  }
}

func mergesort[T Ordered](s []T, l int, r int) {
  if l < r {
    m := l + (r - l) / 2
    mergesort(s, l, m)
    mergesort(s, m+1, r)
    mergeSorted(s, l, m, r)
  }
}

func Mergesort[T Ordered](s []T) {
  mergesort(s, 0, len(s)-1)
}

func partition[T Ordered](s []T, l int, r int) int {
  // choose pivot to be the last element
  pivot := s[r]
  current := l
  for i := l; i < r; i++ {
    if pivot > s[i] {
      swap(&s[current], &s[i])
      current++
    }
  }
  swap(&s[current], &s[r])
  return current
}

func quicksort[T Ordered](s []T, l int, r int) {
  if l < r {
    pivotIndex := partition(s, l, r)
    quicksort(s, l, pivotIndex-1)
    quicksort(s, pivotIndex+1, r)
  }
}

func Quicksort[T Ordered](s []T) {
  quicksort(s, 0, len(s)-1)
}
