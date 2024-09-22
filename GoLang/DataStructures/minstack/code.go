package minstack

// IMPROVED TIME AND SPACE OMPLEXITY
type MinStack struct {
	top  int
	min  []int
	data []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (stack *MinStack) Push(val int) {
	stack.top = val
	stack.data = append(stack.data, val)

	if len(stack.min) == 0 {
		stack.min = append(stack.min, val)
		return
	}

	if val <= stack.min[len(stack.min)-1] {
		stack.min = append(stack.min, val)
	}
}

func (stack *MinStack) Pop() {
	if stack.min[len(stack.min)-1] == stack.top {
		stack.min = stack.min[:len(stack.min)-1]
	}

	stack.data = stack.data[:len(stack.data)-1]
	if len(stack.data) == 0 {
		stack.top = 0
		return
	}

	stack.top = stack.data[len(stack.data)-1]
}

func (stack *MinStack) Top() int {
	return stack.top
}

func (stack *MinStack) GetMin() int {
	return stack.min[len(stack.min)-1]
}

// GOOD
// type MinStack struct {
// 	top  int
// 	data []int
// }

// func Constructor() MinStack {
// 	return MinStack{}
// }

// func (stack *MinStack) Push(val int) {
// 	stack.top = val
// 	stack.data = append(stack.data, val)
// }

// func (stack *MinStack) Pop() {
// 	stack.data = stack.data[:len(stack.data)-1]
//     if len(stack.data) == 0 {
// 		stack.top = 0
// 		return
// 	}

// 	stack.top = stack.data[len(stack.data)-1]
// }

// func (stack *MinStack) Top() int {
// 	return stack.top
// }

// func (stack *MinStack) GetMin() int {
//     min := stack.data[0]
// 	for _, v := range stack.data {
// 		if v < min {
// 			min = v
// 		}
// 	}
// 	return min
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
