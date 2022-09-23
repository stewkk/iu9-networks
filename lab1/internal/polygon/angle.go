package polygon

import (
	"github.com/stewkk/iu9-networks/lab1/internal/dynarray"
	. "github.com/stewkk/iu9-networks/lab1/internal/vertex"
)

type angleIterator struct {
	vertices dynarray.Array
	a        dynarray.Iterator
	b        dynarray.Iterator
	c        dynarray.Iterator
}

func newAngleIterator(arr dynarray.Array) angleIterator {
	return angleIterator{
		vertices: arr,
		a:        arr.Get(0),
		b:        arr.Get(2),
		c:        arr.Get(1),
	}
}

func (it *angleIterator) next() {
	nextVertex := func(v *dynarray.Iterator) {
		if !v.HasNext() {
			*v = it.vertices.First()
			return
		}
		v.Next()
	}
	nextVertex(&it.a)
	nextVertex(&it.b)
	nextVertex(&it.c)
}

func (it angleIterator) hasNext() bool {
	return it.a.HasNext()
}

func (it angleIterator) sgn() int {
	v1 := Vector{X: it.c.Vertex().X - it.a.Vertex().X, Y: it.c.Vertex().Y - it.a.Vertex().Y}
	v2 := Vector{X: it.b.Vertex().X - it.c.Vertex().X, Y: it.b.Vertex().Y - it.c.Vertex().Y}
	return AngleSign(v1, v2)
}
