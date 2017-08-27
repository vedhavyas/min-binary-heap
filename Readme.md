# heap
--
    import "github.com/vedhavyas/min-binary-heap"


## Usage

#### type Heap

```go
type Heap struct {
}
```

Heap represents Minimum Binary Heap

#### func  New

```go
func New() *Heap
```
New returns a new Minimum Binary Heap data structure

#### func (*Heap) Len

```go
func (h *Heap) Len() int
```
Len returns the total size of Heap

#### func (*Heap) Pop

```go
func (h *Heap) Pop() interface{}
```
Pop returns the value with lowest priority if multiple values with same priority
exists, will pop one of them

#### func (*Heap) Push

```go
func (h *Heap) Push(priority int, key interface{})
```
Push takes a key and its priority and adds to the heap
