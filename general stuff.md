// arrays and slices

arr := [5]int{1, 2, 3, 4, 5}  // Fixed-size array
slice := []int{1, 2, 3}       // Dynamic slice
slice = append(slice, 4, 5)  

// maps and hash tables
mp := make(map[string]int)
mp["apple"] = 3
mp["banana"] = 5
fmt.Println(mp["apple"])  // Output: 3


// sets
set := make(map[int]bool)
set[5] = true
set[10] = true
fmt.Println(set[5])   // Output: true


// stacks and queues
// Stack (LIFO)
stack := []int{}
stack = append(stack, 1)   // Push
top := stack[len(stack)-1] // Peek
stack = stack[:len(stack)-1] // Pop

// Queue (FIFO)
queue := []int{}
queue = append(queue, 1)  // Enqueue
front := queue[0]         // Peek
queue = queue[1:]         // Dequeue


// sort & search
import "sort"

arr := []int{1, 3, 5, 7, 9}
idx := sort.SearchInts(arr, 5)  // Finds index of 5
fmt.Println(idx)                // Output: 2


// Recursion and DFS/BFS

// DFS
func dfs(node int, adj map[int][]int, visited map[int]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	fmt.Println(node)

	for _, neighbor := range adj[node] {
		dfs(neighbor, adj, visited)
	}
}


// BFS
import "container/list"

func bfs(start int, adj map[int][]int) {
	queue := list.New()
	queue.PushBack(start)
	visited := make(map[int]bool)
	visited[start] = true

	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(int)
		fmt.Println(node)

		for _, neighbor := range adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
}
