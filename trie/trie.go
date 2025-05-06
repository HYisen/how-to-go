package trie

import (
	"iter"
)

type Node[K comparable, V any] struct {
	children map[K]*Node[K, V]
	value    *V
}

func New[K comparable, V any]() *Node[K, V] {
	return &Node[K, V]{
		children: make(map[K]*Node[K, V]),
	}
}

func (n *Node[K, V]) Set(value V, keys ...K) {
	n.set(&value, keys...)
}

func (n *Node[K, V]) set(value *V, keys ...K) {
	if len(keys) == 0 {
		n.value = value
		return
	}
	if _, ok := n.children[keys[0]]; !ok {
		n.children[keys[0]] = New[K, V]()
	}
	n.children[keys[0]].set(value, keys[1:]...)
}

func (n *Node[K, V]) Get(keys ...K) *V {
	if len(keys) == 0 {
		return n.value
	}
	if next, ok := n.children[keys[0]]; ok {
		return next.Get(keys[1:]...)
	}
	return nil
}

func (n *Node[K, V]) All() iter.Seq2[[]K, *V] {
	return func(yield func([]K, *V) bool) {
		n.push(nil, yield)
	}
}

func (n *Node[K, V]) push(path []K, yield func([]K, *V) bool) bool {
	if n.value != nil {
		yield(path, n.value)
	}
	for k, next := range n.children {
		next.push(append(path, k), yield)
	}
	return false
}

func (n *Node[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		nodes := []*Node[K, V]{n}
		for len(nodes) > 0 {
			var next []*Node[K, V]
			for _, node := range nodes {
				if node.value != nil {
					yield(*node.value)
				}
				for _, child := range node.children {
					next = append(next, child)
				}
			}
			nodes = next
		}
	}
}
