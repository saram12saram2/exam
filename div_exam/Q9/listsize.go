package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
} // - comment if no need, but likely needs to be copied from the task and be used together with the function

func ListSize(l *List) int {
	count := 0
	iterator := l.Head

	for iterator != nil {
		count++
		iterator = iterator.Next
	}
	return count
}
