<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# linkedlist

```go
import "go-ds/list/linked_list"
```

## Index

- [type LL](<#type-ll>)
  - [func NewLinkedList[K comparable]() *LL[K]](<#func-newlinkedlist>)
  - [func (l *LL[K]) At(i int) (K, error)](<#func-llk-at>)
  - [func (l *LL[K]) Delete(v K) bool](<#func-llk-delete>)
  - [func (l *LL[K]) Each(fn func(*LLNode[K], int))](<#func-llk-each>)
  - [func (l *LL[K]) Pop() *LL[K]](<#func-llk-pop>)
  - [func (l *LL[K]) Push(v K) *LL[K]](<#func-llk-push>)
- [type LLNode](<#type-llnode>)


## type LL

```go
type LL[K comparable] struct {
    Length int
    Front  *LLNode[K]
    Back   *LLNode[K]
}
```

### func NewLinkedList

```go
func NewLinkedList[K comparable]() *LL[K]
```

### func \(\*LL\[K\]\) At

```go
func (l *LL[K]) At(i int) (K, error)
```

Returns the element at the given index

### func \(\*LL\[K\]\) Delete

```go
func (l *LL[K]) Delete(v K) bool
```

Removes the item if exists, returning a bool on whether the items have been removed or not

### func \(\*LL\[K\]\) Each

```go
func (l *LL[K]) Each(fn func(*LLNode[K], int))
```

Runs the given function to each item of the list

### func \(\*LL\[K\]\) Pop

```go
func (l *LL[K]) Pop() *LL[K]
```

Removes the last item and returns the same linked list with the last item removed

### func \(\*LL\[K\]\) Push

```go
func (l *LL[K]) Push(v K) *LL[K]
```

Adds an item to the given linked list and returns the same linked list with the added items

## type LLNode

```go
type LLNode[K comparable] struct {
    Next  *LLNode[K]
    Value K
}
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)