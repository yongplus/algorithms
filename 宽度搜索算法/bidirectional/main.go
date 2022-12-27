package main

/**
 * https://leetcode.cn/problems/find-if-path-exists-in-graph/
 *
有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回 true，否则返回 false 。


示例 1：

输入：n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
输出：true
解释：存在由顶点 0 到顶点 2 的路径:
- 0 → 1 → 2
- 0 → 2
示例 2：


输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
输出：false
解释：不存在由顶点 0 到顶点 5 的路径.


提示：

1 <= n <= 2 * 105
0 <= edges.length <= 2 * 105
edges[i].length == 2
0 <= ui, vi <= n - 1
ui != vi
0 <= source, destination <= n - 1
不存在重复边
不存在指向顶点自身的边

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/find-if-path-exists-in-graph
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
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
