package algorithm

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	root := makeTree()

	//深度优先遍历
	s := NewStack()
	DFSTraversal(s, root) //expect: abdecf

	//广度优先遍历
	q := NewQueue()
	BFSTraversal(q, root) //expect: abcdef
}

//构造一颗树
func makeTree() *TreeNode {
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")
	e := NewNode("e")
	f := NewNode("f")

	a.AddChild(b)
	a.AddChild(c)
	b.AddChild(d)
	b.AddChild(e)
	c.AddChild(f)

	return a
}

func DFSTraversal(s *Stack, root *TreeNode) {
	s.Push(root)
	for {
		if s.IsEmpty() {
			break
		}

		node := s.Pop().(*TreeNode)

		fmt.Println(node.GetData())

		for _, child := range reverse(node.GetChildren()) {
			s.Push(child)
		}
	}
}

func BFSTraversal(q *Queue, root *TreeNode) {
	q.Enqueue(root)
	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue().(*TreeNode)

		fmt.Println(node.GetData())

		for _, child := range node.GetChildren() {
			q.Enqueue(child)
		}
	}
}

func reverse(a []*TreeNode) []*TreeNode {
	l := len(a)
	b := make([]*TreeNode, l)
	for i, v := range a {
		b[l-1-i] = v
	}
	return b
}
