func moveZeroes(nums []int)  {
    n := len(nums)
    queue := Queue{elems: make([]int, 0)}
    target := 0

    // [0, 1, 0, 3, 12]
	for i := 0; i < n; i++ {
        if nums[i] == target {
            queue.Enqueue(i)
            continue  
        }           
        
        idx := queue.Dequeue()
        nums[idx], nums[i] = nums[i], nums[idx]
        
        if queue.Size() == 0 { 
            queue.Enqueue(idx)
        }
	}
}

type Queue struct {
    elems []int
}

func (q *Queue) Enqueue(elem int) {
    q.elems = append(q.elems, elem)
}

func (q *Queue) Size() int {
    return len(q.elems)
}


func (q *Queue) Dequeue() int {
    e := q.elems[0]
    q.elems = q.elems[1:]
    
    return e
}
