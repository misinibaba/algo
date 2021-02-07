package main

import "fmt"

type MyQueue struct {
	input []int
	output []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue)pushToPop() {
	if len(this.output) == 0 {
		for _, v := range this.input {
			this.output = append(this.output, v)
		}
		this.input = nil
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.input = append(this.input, x)
	this.pushToPop()
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.pushToPop()
	out := this.output[0]
	this.output = this.output[1:]
	return out
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.output) == 0 {
		return 0
	}
	return this.output[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.output) == 0 && len(this.input) == 0
}



func main() {
	obj := Constructor();
	obj.Push(2);
	param_2 := obj.Pop();
	fmt.Println(param_2)
	param_3 := obj.Peek();
	fmt.Println(param_3)
	param_4 := obj.Empty();
	fmt.Println(param_4)
}

