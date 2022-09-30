package polygon

import "math/rand"

func NewTreapPolygon(vertices []Vertex) Polygon {
	// TODO: построение за O(n)
	res := treapPolygon{}
	for i, v := range vertices {
		res.Insert(i, v)
	}
	return &res
}

type treapPolygon struct {
	root *node
}

func (tree *treapPolygon) Delete(idx int) error {
	root := &tree.root
	for {
		leftSize := size((*root).left)
		if idx == leftSize {
			*root = merge((*root).left, (*root).right)
			return nil
		}
		(*root).subtreeSize--
		if idx < leftSize {
			root = &(*root).left
		} else {
			root = &(*root).right
			idx -= leftSize + 1
		}
	}
}

func (tree *treapPolygon) Insert(idx int, v Vertex) error {
	newNode := newNode(v)
	if tree.root == nil {
		tree.root = newNode
		return nil
	}
	l, r := split(tree.root, idx)
	tree.root = merge(l, newNode)
	tree.root = merge(tree.root, r)
	return nil
}

func (tree *treapPolygon) Set(idx int, v Vertex) error {
	tree.node(idx).v = v
	return nil
}

func (tree *treapPolygon) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.subtreeSize
}

func (tree *treapPolygon) Vertex(idx int) Vertex {
	return tree.node(idx).v
}

func (tree *treapPolygon) Vertices() (vertices []Vertex) {
	vertices = []Vertex{}
	var recVertices func(cur *node)
	recVertices = func(cur *node) {
		if cur == nil {
			return
		}
		recVertices(cur.left)
		vertices = append(vertices, cur.v)
		recVertices(cur.right)
	}
	recVertices(tree.root)
	return
}

func (*treapPolygon) Polyline(from int, to int) []Vertex {
	return []Vertex{}
}

func (*treapPolygon) IsConvex() bool {
	return false
}

type node struct {
	left        *node
	right       *node
	v           Vertex
	subtreeSize int
	priority    int
}

func newNode(v Vertex) *node {
	return &node{
		left:        nil,
		right:       nil,
		v:           v,
		subtreeSize: 1,
		priority:    rand.Int(),
	}
}

func (tree *treapPolygon) node(idx int) *node {
	root := tree.root
	for {
		leftSize := size(root.left)
		if idx == leftSize {
			return root
		}
		if idx < leftSize {
			root = root.left
		} else {
			root = root.right
			idx -= leftSize + 1
		}
	}
}

func split(root *node, key int) (l, r *node) {
	if root == nil {
		return nil, nil
	}
	if key <= size(root.left) {
		l, root.left = split(root.left, key)
		root.recalcSize()
		return l, root
	} else {
		root.right, r = split(root.right, key-size(root.left)-1)
		root.recalcSize()
		return root, r
	}
}

func size(root *node) int {
	if root == nil {
		return 0
	}
	return root.subtreeSize
}

func (root *node) recalcSize() {
	root.subtreeSize = size(root.left) + 1 + size(root.right)
}

func merge(l, r *node) *node {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.priority > r.priority {
		l.right = merge(l.right, r)
		l.recalcSize()
		return l
	} else {
		r.left = merge(l, r.left)
		r.recalcSize()
		return r
	}
}

func (p *treapPolygon) countPolygonAngleSignSum() int {
	return 0
}

func (p *treapPolygon) polylineAngleSignSum(from, count int) int {
	return 0
}
