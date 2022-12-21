package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type Queue[T any] struct {
	Elements []T
	Size     int
}

func (q *Queue[T]) Enqueue(elem T) {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return *new(T)
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element // Slice off the element once it is dequeued.
}

func (q *Queue[T]) GetLength() int {
	return len(q.Elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue[T]) Peek() (any, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	//adj是一个二维数组 adj[i][j] 描述的index => values相邻关系
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, edge := range edges {
		x := edge[0]
		y := edge[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}

	//visited 表示访问过得
	visited := make([]bool, n)
	queue := Queue[int]{Size: 1000}
	queue.Enqueue(source)
	visited[source] = true

	for !queue.IsEmpty() {
		vertex := queue.Dequeue()
		if vertex == destination {
			break
		}

		for _, next := range adj[vertex] {
			if !visited[next] {
				queue.Enqueue(next)
				visited[next] = true
			}
		}
	}
	return visited[destination]
}

func main() {
	t1 := time.Now()
	/*
		n := 3
		edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
		source := 0
		destination := 2
	*/
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source := 0
	destination := 5
	result := validPath(n, edges, source, destination)
	log.Println(result)
	log.Println(time.Now().Sub(t1))
}
