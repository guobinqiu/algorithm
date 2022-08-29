package algorithm

import (
	"testing"
)

func TestDag(t *testing.T) {
	dag := &DAG{}

	//定义顶点
	v1 := &Vertex{Key: "1"}
	v2 := &Vertex{Key: "2"}
	v3 := &Vertex{Key: "3"}
	v4 := &Vertex{Key: "4"}
	v5 := &Vertex{Key: "5"}

	//定义边
	dag.AddEdge(v1, v5)
	dag.AddEdge(v1, v2)
	dag.AddEdge(v1, v3)
	dag.AddEdge(v3, v4)
	dag.AddEdge(v3, v2)
	dag.AddEdge(v2, v4)

	//广度优先搜索
	dag.BFS(v1)
}
