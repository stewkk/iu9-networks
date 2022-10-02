package polygon

import (
	"fmt"
	"math/rand"
)

func NewTreapPolygon(vertices []Vertex) (Polygon, error) {
	// TODO: построение за O(n)
	res := treapPolygon{}
	if len(vertices) < 3 {
		return &res, fmt.Errorf("%w: can't construct polygon from less than 3 vertices", ErrInvalidOperation)
	}
	for i, v := range vertices {
		res.Insert(i, v)
	}
	return &res, nil
}

type treapPolygon struct {
	root *node
	angleSignSum int
}

func (tree *treapPolygon) Delete(idx int) error {
	if tree.Size() == 3 {
		return fmt.Errorf("%w: can't delete vertex from 3-vertex polygon", ErrInvalidOperation)
	}
	if idx < 0 || idx >= tree.Size() {
		return ErrOutOfBounds
	}
	tree.angleSignSum -= polylineAngleSignSum(tree.verticesOfAngles(idx-1, 3))
	root := &tree.root
	for {
		leftSize := size((*root).left)
		if idx == leftSize {
			*root = merge((*root).left, (*root).right)
			tree.angleSignSum += polylineAngleSignSum(tree.verticesOfAngles(idx-1, 2))
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
	if idx < 0 || idx > tree.Size() {
		return ErrOutOfBounds
	}

	tree.angleSignSum -= polylineAngleSignSum(tree.verticesOfAngles(idx-1, 2))

	newNode := newNode(v)
	if tree.root == nil {
		tree.root = newNode
		return nil
	}
	l, r := split(tree.root, idx)
	tree.root = merge(l, newNode)
	tree.root = merge(tree.root, r)

	tree.angleSignSum += polylineAngleSignSum(tree.verticesOfAngles(idx-1, 3))

	return nil
}

func (tree *treapPolygon) Set(idx int, v Vertex) error {
	if idx < 0 || idx >= tree.Size() {
		return ErrOutOfBounds
	}
	tree.angleSignSum -= polylineAngleSignSum(tree.verticesOfAngles(idx-1, 3))
	tree.node(idx).v = v
	tree.angleSignSum += polylineAngleSignSum(tree.verticesOfAngles(idx-1, 3))
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

func (p *treapPolygon) IsConvex() bool {
	return p.Size() == abs(p.angleSignSum)
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

func (tree *treapPolygon) verticesOfAngles(from, count int) (vertices []Vertex) {
	if from <= 0 {
		vertices = tree.subset(tree.Size()+from-1, tree.Size())
		vertices = append(vertices, tree.subset(0, count+from+1)...)
		return
	}
	if from + count + 1 > tree.Size() {
		vertices = tree.subset(from-1, tree.Size())
		vertices = append(vertices, tree.subset(0, count+from-tree.Size()+1)...)
		return
	}
	vertices = tree.subset(from-1, from+count+1)
	return
}

func (tree *treapPolygon) subset(start, end int) (vertices []Vertex) {
	l, m := split(tree.root, start)
	m, r := split(m, end - start)
	tmp := treapPolygon{root: m}
	vertices = tmp.Vertices()
	l = merge(l, m)
	tree.root = merge(l, r)
	return
}
