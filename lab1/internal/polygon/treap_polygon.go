package polygon

import (
	"fmt"
	"math/rand"
)

func NewTreapPolygon(vertices []Vertex) (Polygon, error) {
	res := treapPolygon{}
	if len(vertices) < 3 {
		return &res, fmt.Errorf("%w: can't construct polygon from less than 3 vertices", ErrInvalidOperation)
	}
	res.nodes = make([]node, 120000)
	res.root = res.build(vertices)
	res.angleSignSum = countPolygonAngleSignSum(vertices)
	return &res, nil
}

type treapPolygon struct {
	root         *node
	angleSignSum int
	nodes        []node
	last         int
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

	newNode := tree.newNode(v)
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
	return tree.root.subtreeSize
}

func (tree *treapPolygon) Vertex(idx int) Vertex {
	return tree.node(idx).v
}

func (tree *treapPolygon) Vertices() (vertices []Vertex) {
	vertices = make([]Vertex, 0, 5)
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

func (tree *treapPolygon) newNode(v Vertex) *node {
	node := &tree.nodes[tree.last]
	node.priority = rand.Int()
	node.subtreeSize = 1
	node.v = v
	tree.last++
	return node
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
	if from+count+1 > tree.Size() {
		vertices = tree.subset(from-1, tree.Size())
		vertices = append(vertices, tree.subset(0, count+from-tree.Size()+1)...)
		return
	}
	vertices = tree.subset(from-1, from+count+1)
	return
}

func (tree *treapPolygon) subset(start, end int) (vertices []Vertex) {
	l, m := split(tree.root, start)
	m, r := split(m, end-start)
	tmp := treapPolygon{root: m}
	vertices = tmp.Vertices()
	l = merge(l, m)
	tree.root = merge(l, r)
	return
}

func (tree *treapPolygon) build(vertices []Vertex) *node {
	if len(vertices) == 0 {
		return nil
	}
	m := len(vertices) / 2
	root := tree.newNode(vertices[m])
	root.left = tree.build(vertices[:m])
	root.right = tree.build(vertices[m+1:])
	heapify(root)
	root.recalcSize()
	return root
}

func heapify(root *node) {
	max := root
	if root.left != nil && root.left.priority > max.priority {
		max = root.left
	}
	if root.right != nil && root.right.priority > max.priority {
		max = root.right
	}
	if max != root {
		max.priority, root.priority = root.priority, max.priority
		heapify(max)
	}
}
