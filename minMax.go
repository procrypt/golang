package main

type MinMaxStack struct {
	// Write your code here.
	stack       []int
	minMaxStack []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	// Write your code here.
	return &MinMaxStack{}
}

func (stack *MinMaxStack) Peek() int {
	// Write your code here.
	return stack.stack[len(stack.stack)-1]
}

func (stack *MinMaxStack) Pop() int {
	// Write your code here.
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	pop := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return pop
}

func (stack *MinMaxStack) Push(number int) int {
	// Write your code here.
	newMinMax := entry{min:number, max:number}
	if len(stack.minMaxStack) > 0 {
		lastMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(lastMinMax.min, number)
		newMinMax.max = max(lastMinMax.max, number)
	}
	stack.stack = append(stack.stack, number)
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
}

func (stack *MinMaxStack) GetMin() int {
	// Write your code here.
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	// Write your code here.
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}


func min(a,b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a,b int) int {
	if a < b {
		return b
	}
	return a
}