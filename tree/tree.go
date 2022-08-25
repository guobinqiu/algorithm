package tree

type TreeNode struct {
	data     interface{}
	parent   *TreeNode
	children []*TreeNode
}

func NewNode(data interface{}) *TreeNode {
	return &TreeNode{
		data: data,
	}
}

func (n *TreeNode) AddChild(child *TreeNode) {
	n.children = append(n.children, child)
}

func (n *TreeNode) GetChildren() []*TreeNode {
	return n.children
}

func (n *TreeNode) SetParent(parent *TreeNode) {
	n.parent = parent
}

func (n *TreeNode) GetParent() *TreeNode {
	return n.parent
}

func (n *TreeNode) IsRoot() bool {
	return n.parent == nil
}

func (n *TreeNode) IsLeaf() bool {
	return len(n.children) == 0
}

func (n *TreeNode) GetData() interface{} {
	return n.data
}
