package main

import (
	"github.com/guobinqiu/algorithm/queue"
	"github.com/guobinqiu/algorithm/stack"
	"github.com/guobinqiu/algorithm/tree"
	"fmt"
)

func main() {
	root := makeTree()

	//深度优先遍历
	s := stack.New()
	DFSTraversal(s, root) //expect: abdecf

	//广度优先遍历
	q := queue.New()
	BFSTraversal(q, root) //expect: abcdef
}

//构造一颗树
func makeTree() *tree.TreeNode {
	a := tree.NewNode("a")
	b := tree.NewNode("b")
	c := tree.NewNode("c")
	d := tree.NewNode("d")
	e := tree.NewNode("e")
	f := tree.NewNode("f")

	a.AddChild(b)
	a.AddChild(c)
	b.AddChild(d)
	b.AddChild(e)
	c.AddChild(f)

	return a
}

func DFSTraversal(s *stack.Stack, root *tree.TreeNode) {
	s.Push(root)
	for {
		if s.IsEmpty() {
			break
		}

		node := s.Pop().(*tree.TreeNode)

		fmt.Println(node.GetData())

		for _, child := range reverse(node.GetChildren()) {
			s.Push(child)
		}
	}
}

func BFSTraversal(q *queue.Queue, root *tree.TreeNode) {
	q.Enqueue(root)
	for {
		if q.IsEmpty() {
			break
		}

		node := q.Dequeue().(*tree.TreeNode)

		fmt.Println(node.GetData())

		for _, child := range node.GetChildren() {
			q.Enqueue(child)
		}
	}
}

func reverse(a []*tree.TreeNode) []*tree.TreeNode {
	l := len(a)
	b := make([]*tree.TreeNode, l)
	for i, v := range a {
		b[l-1-i] = v
	}
	return b
}
