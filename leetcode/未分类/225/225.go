package main

import "fmt"

type MyStack struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor() (s MyStack) {
	return
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	x := this.queue[len(this.queue) - 1]
	this.queue = this.queue[:len(this.queue) - 1]
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.queue[len(this.queue) - 1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func main() {
	 obj := Constructor();
	 obj.Push(1);
	 obj.Push(2);
	 param_2 := obj.Top();
	 fmt.Println(param_2)
	 param_3 := obj.Pop();
	 fmt.Println(param_3)
	 param_4 := obj.Empty();
	 fmt.Println(param_4)
}

