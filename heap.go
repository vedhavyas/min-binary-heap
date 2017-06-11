package heap

// Heap represents Minimum Binary Heap
type Heap struct {
	pipe        []interface{}
	priorityMap map[interface{}]int
}

// New returns a new Minimum Binary Heap data structure
func New() *Heap {
	return &Heap{
		priorityMap: make(map[interface{}]int),
	}
}

// bubbleUp will recursively check and satisfy the heap constraints from bottom
func (h *Heap) bubbleUp(cn int) {
	if cn == 1 {
		return
	}

	cnp := h.priorityMap[h.pipe[cn-1]]
	pn := cn / 2
	pnp := h.priorityMap[h.pipe[pn-1]]
	if cnp > pnp {
		return
	}

	h.pipe = swap(h.pipe, pn-1, cn-1)
	h.bubbleUp(pn)
}

// swap will swap positions of a, b in list l
func swap(l []interface{}, a, b int) []interface{} {
	t := l[a]
	l[a] = l[b]
	l[b] = t
	return l
}

// Push takes a key and its priority and adds to the heap
func (h *Heap) Push(priority int, key interface{}) {
	_, ok := h.priorityMap[key]
	if ok {
		return
	}
	h.pipe = append(h.pipe, key)
	h.priorityMap[key] = priority
	h.bubbleUp(len(h.pipe))
}

// bubbleDown will recursively check for heap constraints from top to bottom
func (h *Heap) bubbleDown(pn int) {
	if pn == len(h.pipe) {
		return
	}

	pnp := h.priorityMap[h.pipe[pn-1]]
	cn := 2 * pn
	if cn > len(h.pipe) {
		return
	}

	cnp := h.priorityMap[h.pipe[cn-1]]
	if cn+1 <= len(h.pipe) {
		cn2p := h.priorityMap[h.pipe[cn]]
		if cn2p < cn {
			cn += 1
			cnp = cn2p
		}
	}

	if cnp >= pnp {
		return
	}

	h.pipe = swap(h.pipe, cn-1, pn-1)
	h.bubbleDown(cn)
}

// Pop returns the value with lowest priority
// if multiple values with same priority exists, will pop one of them
func (h *Heap) Pop() interface{} {
	n := h.pipe[0]
	delete(h.priorityMap, n)
	h.pipe = h.pipe[1:]
	if len(h.pipe) == 0 {
		return n
	}

	t := []interface{}{h.pipe[len(h.pipe)-1]}
	h.pipe = append(t, h.pipe[:len(h.pipe)-1]...)
	h.bubbleDown(1)
	return n
}

// Len returns the total size of Heap
func (h *Heap) Len() int {
	return len(h.pipe)
}
