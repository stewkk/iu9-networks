package polygon

import (
	"math/rand"
	"testing"
)

func BenchmarkInit(b *testing.B) {
	vertices := []Vertex{}
	for i := 0; i < 100000; i++ {
		vertices = append(vertices, Vertex{i, i})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewSlicePolygon(vertices)
	}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		p, _ := NewSlicePolygon([]Vertex{{1, 2}, {3, 4}, {5, 6}})
		b.StartTimer()
		for i := 0; i < 10000; i++ {
			p.Insert(rand.Intn(i+1), Vertex{i, i})
		}
	}
}

func BenchmarkSet(b *testing.B) {
	vertices := []Vertex{}
	for i := 0; i < 10000; i++ {
		vertices = append(vertices, Vertex{i, i})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		p, _ := NewSlicePolygon(vertices)
		b.StartTimer()
		for i := 0; i < len(vertices); i++ {
			p.Set(rand.Intn(len(vertices)), Vertex{i, 0})
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	vertices := []Vertex{}
	for i := 0; i < 10000; i++ {
		vertices = append(vertices, Vertex{i, i})
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		p, _ := NewSlicePolygon(vertices)
		b.StartTimer()
		for i := 0; i < len(vertices) - 5; i++ {
			p.Delete(rand.Intn(len(vertices) - i))
		}
	}
}

func BenchmarkGet(b *testing.B) {
	vertices := []Vertex{}
	for i := 0; i < 10000; i++ {
		vertices = append(vertices, Vertex{i, i})
	}
	p, _ := NewSlicePolygon(vertices)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(vertices); i++ {
			p.Vertex(rand.Intn(len(vertices)))
		}
	}

}
