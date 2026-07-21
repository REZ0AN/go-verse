# Collections in GO


## Arrays

- Fixed sized
- Value Type
- Less popular
```go
var arr [5]int                    // zero-valued: [0 0 0 0 0]
arr2 := [3]int{1, 2, 3}            // literal with explicit size
arr3 := [...]int{1, 2, 3, 4}       // size inferred from literal → [4]int
```

## Slices

- Flexible Size
- Build top of Arrays
- Multiple slices can share the same underlying array (one mutate, others reflect that mutation)
- Reference Type
- Under the hood it's a small struct with *ptr  to the underlying array, len accessible value count, cap how many value we can put.
- If cap is filled then new underlying array is created with a larger capacity, copies values from the old ones, and cut off the reference to the old array, and point to the new underlying array.

```go
var s []int                        // nil slice — len=0, cap=0, ptr=nil
s1 := []int{1, 2, 3}                // slice literal
s2 := make([]int, 5)                 // len=5, cap=5, all zero-valued
s3 := make([]int, 3, 10)              // len=3, cap=10

sub := s1[0:2]                        // slicing: [low:high), high is EXCLUSIVE
sub2 := s1[:2]                         // omit low → defaults to 0
sub3 := s1[1:]                          // omit high → defaults to len(s1)

s1 = append(s1, 4)                       // grow by one
s1 = append(s1, 5, 6, 7)                  // grow by multiple
s1 = append(s1, sub...)                    // spread another slice in

len(s1)  // current element count
cap(s1)  // current underlying array capacity from this slice's start

copy(dst, src) // copies min(len(dst), len(src)) elements, returns count copied
```

## Maps

- Flexible
- Reference Type
- Underlying implementation is hashtable where current implementation holds 8 entries per bucket, if one of the bucket is filled then, it put entries to a new bucket.
- There is no guranteed order for the iteration of a Map.

```go
var m map[string]int              // nil map — reads OK, writes panic
m2 := make(map[string]int)          // usable, empty map
m3 := map[string]int{"a": 1, "b": 2} // map literal

m2["x"] = 10                          // insert/update
v := m2["x"]                            // read (returns zero value if absent)
v, ok := m2["y"]                          // "comma ok" idiom — ok=false if absent
delete(m2, "x")                            // remove a key
len(m2)                                     // number of key-value pairs

for k, v := range m2 { }                     // iterate (unordered!)
```

## Structs

- Value Type
- Way of grouping data

```go
type Point struct {
    X, Y int
    Label string
}

p1 := Point{X: 1, Y: 2, Label: "origin"} // named fields (preferred — order-independent)
p2 := Point{1, 2, "origin"}                // positional (order MUST match declaration)
p3 := Point{}                                // zero-valued struct: {0 0 ""}

fmt.Println(p1.X)   // field access via dot notation
p1.X = 99             // field mutation
```

## Pointers

- All functionality like `C pointers` except the arithmetic operation in Idiometic GO.
- This feature or design decision eliminates an enormous class of C memory-safety bugs.

```go
x := 5
p := &x            // p is *int
fmt.Println(*p)      // dereference to read: 5
*p = 10                // dereference to write
fmt.Println(x)           // 10 — x itself changed

pt := &Point{X: 1, Y: 2}   // pointer to a struct
pt.X = 99                    // Go auto-dereferences for struct field access —
                               // no need to write (*pt).X, though that also works
```


## Methods — receiver syntax

- Concept of **Pass by Value**, **Pass by Reference**

```go
func (p Point) Distance() float64 { ... }   // value receiver
func (p *Point) Move(dx, dy int) {           // pointer receiver
    p.X += dx
    p.Y += dy
}
```