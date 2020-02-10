package main

import (
	"sync"
	"fmt"
)

type MinStack struct {
	items	[]int
	lock 	sync.RWMutex
}


/** initialize your data structure here. */
func Constructor() MinStack {
	minStack := MinStack{}
	minStack.items = []int{}
	return minStack
}


func (this *MinStack) Push(x int)  {
	this.lock.Lock()
	this.items = append(this.items, x)
	this.lock.Unlock()
}


func (this *MinStack) Pop()  {
	this.lock.Lock()
	this.items = this.items[:len(this.items)-1]
	this.lock.Unlock()
}


func (this *MinStack) Top() int {
	this.lock.Lock()
	topItem := this.items[len(this.items)-1]
	return topItem

}


func (this *MinStack) GetMin() int {
	this.lock.Lock()
	min := this.items[0]
	for i:=1;i<len(this.items);i++{
		if this.items[i]< min {
			min = this.items[i]
		}
	}
	this.lock.Unlock()
	return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(-1)
	minStack.Push(3)
	fmt.Println("min num",minStack.GetMin())
	fmt.Println(minStack.items)
	minStack.Pop()
	fmt.Println(minStack.items)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.items)

	tmp := map[int]int{}
	tmp[300] = 1
	tmp[50]=2
	for k,v := range tmp{
		fmt.Println(k,v)
	}



}