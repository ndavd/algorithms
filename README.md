## Implementation of common algorithms and data structures in Go :blue_heart:

### Algorithms
Supported algorithms:
- sort
  - Selection
  - Insertion
  - Bubble
  - Mergesort
  - Quicksort

### Usage
```go
// main.go
package main

import (
  "github.com/ndavd/algorithms/sort"
  "fmt"
)

func main() {
  a := []byte("Algorithms in Go!") // [A l g o r i t h m s   i n   G o !]
  sort.Quicksort(a)
  fmt.Printf("%c", a)              // [    ! A G g h i i l m n o o r s t]
}
```
