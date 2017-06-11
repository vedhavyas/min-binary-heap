package heap

import "testing"

func TestHeap_Push(t *testing.T) {
	h := New()
	h.Push(5, "a")
	h.Push(2, "b")
	h.Push(3, "c")
	h.Push(1, "d")
	h.Push(0, "e")

	if l := h.Len(); l != 5 {
		t.Fatalf("Expected lenght %d but got %d", 5, l)
	}
}

func TestHeap_Pop(t *testing.T) {
	h := New()
	h.Push(5, "a")
	h.Push(2, "b")
	h.Push(3, "c")
	h.Push(1, "d")
	h.Push(0, "e")
	h.Push(7, "e")

	if l := h.Len(); l != 5 {
		t.Fatalf("Expected lenght %d but got %d", 5, l)
	}

	for _, v := range []string{"e", "d", "b", "c", "a"} {
		n := h.Pop().(string)
		if v != n {
			t.Fatalf("Expected node %s but got %s", v, n)
		}
	}
}
