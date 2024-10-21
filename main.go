package main

import (
	maths "github.com/ChrisShia/math-depot"
)

var stacks_ []*[]int

func main() {
	var n = 10
	stack1 := make([]int, 10)
	stack2 := make([]int, 0)
	stack3 := make([]int, 0)
	stacks_ = []*[]int{&stack1, &stack2, &stack3}
	j := 10
	for i := 0; i < n; i++ {
		stack1[i] = j
		j--
	}
	counter := maths.NewModCounter(0, 3)
	movePartition(0, 0, counter, stack1[0])
}

func movePartition2(src []int) {
	if len(src) > 1 {
		par := make([]int, 0)
		par = src[1:]
		movePartition2(par)
	} else if len(src) == 1 {
		if src[0]%2 > 0 {

		}
	}
	movePartition2(src[:1])
}

func movePartition(srcIndex, from int, counter *maths.ModuloCounter, root int) {
	var lastAppendedTo = srcIndex
	if from < len(*stacks_[srcIndex])-1 {
		movePartition(srcIndex, from+1, counter, root-1)
		lastAppendedTo = counter.ModIndex()
	}
	//} else if from > len(*stacks_[srcIndex]) {
	//	return
	//}
	var popped int
	//if len(*stacks_[srcIndex]) > 0 && (*stacks_[srcIndex])[len(*stacks_[srcIndex])-1] <= root {
	if len(*stacks_[srcIndex]) > 0 {
		popped = pop(stacks_[srcIndex])
	}
	if counter.Next() == srcIndex {
		counter.Next()
	}
	addToStack(stacks_[counter.ModIndex()], popped)
	if lastAppendedTo != srcIndex {
		//c := maths.NewModCounter(0, 3)
		counter.Offset(lastAppendedTo)
		movePartition(lastAppendedTo, order(lastAppendedTo, root), counter, root)
		//movePartition(lastAppendedTo, order(lastAppendedTo, root), c, root)
	}
}

func order(srcIndex, value int) int {
	src := *stacks_[srcIndex]
	for i := len(src) - 1; i >= 0; i-- {
		if src[i] > value {
			return i + 1
		}
	}
	return 0
}

func nextAvailableBufferStack(c *maths.ModuloCounter) int {
	c.Next()
	if c.ModIndex() == c.Start {
		return c.Next()
	}
	return c.ModIndex()
}

func popStack(stack *[]int, from, to int) []int {
	popped := (*stack)[from:to]
	return popped
}

func pop(stack *[]int) int {
	popped := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return popped
}

func addToStack(dest *[]int, value int) {
	*dest = append(*dest, value)
}
