package multi

import "iter"

type Node[K comparable, V any] struct {
	children map[K]*Node[K, V]
	values   []V
}

func New[k comparable, V any]() *Node[k, V] {
	return &Node[k, V]{
		children: make(map[k]*Node[k, V]),
	}
}

func (n *Node[K, V]) Add(keys []K, values ...V) {
	if len(keys) == 0 {
		n.values = append(n.values, values...)
		return
	}
	if _, ok := n.children[keys[0]]; !ok {
		n.children[keys[0]] = New[K, V]()
	}
	n.children[keys[0]].Add(keys[1:], values...)
}

func (n *Node[K, V]) Get(keys []K) []V {
	if len(keys) == 0 {
		return n.values
	}
	if next, ok := n.children[keys[0]]; ok {
		return next.Get(keys[1:])
	}
	return nil
}

func (n *Node[K, V]) All() iter.Seq2[[]K, []V] {
	return func(yield func([]K, []V) bool) {
		n.push(nil, yield)
	}
}

func (n *Node[K, V]) push(path []K, yield func([]K, []V) bool) bool {
	if len(n.values) > 0 {
		yield(path, n.values)
	}
	for k, next := range n.children {
		next.push(append(path, k), yield)
	}
	return false
}
