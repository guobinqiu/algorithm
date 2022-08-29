package algorithm

import "fmt"

type Vertex struct {
	Key      string
	Parents  []*Vertex
	Children []*Vertex
	Value    interface{}
}

type DAG struct{}

func (dag *DAG) AddEdge(from, to *Vertex) {
	from.Children = append(from.Children, to)
	to.Parents = append(to.Parents, from)
}

func (dag *DAG) BFS(root *Vertex) {
	q := NewQueue()

	visitMap := make(map[string]bool)
	visitMap[root.Key] = true

	q.Enqueue(root)

	for {
		if q.Size() == 0 {
			fmt.Println("done")
			break
		}

		current := q.Dequeue().(*Vertex)

		//fmt.Println("bfs key", current.Key)

		for _, v := range current.Children {
			fmt.Printf("from:%v to:%s\n", current.Key, v.Key)

			if v.Key == root.Key {
				panic("back root")
			}

			if _, ok := visitMap[v.Key]; !ok {
				visitMap[v.Key] = true
				q.Enqueue(v)
			}
		}
	}
}
