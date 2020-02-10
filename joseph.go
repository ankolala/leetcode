package main

import "fmt"

// 约瑟夫问题
// 小孩
type Boy struct {
	no   int // 小孩编号
	next *Boy
}

// 编写一个函数构成单向的环形链表
// num表示约瑟夫中小孩的个数，返回环形链表的第一个小孩，也就是头结点
func AddBoy(num int) *Boy {

	head := &Boy{} // &{0, <nil>}
	curBoy := &Boy{}
	if num < 1 {
		return head
	}
	// 循环构建环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no:   i,
			next: nil,
		}

		if i == 1 {

			head = boy
			curBoy = boy
			curBoy.next = head
			fmt.Println(head, head.next,head.next.next)
		} else {

			curBoy.next = boy
			curBoy = boy
			curBoy.next = head
			fmt.Println("curBoy",curBoy)
		}
	}
	fmt.Println(head, head.next, head.next.next)
	return head
}

// 打印
func showBoy(head *Boy) {
	temp := head
	for {
		fmt.Println("boy: ", temp.no)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}


// 求出最后一个小孩，first为编号为1的小孩，startNo为游戏开始的小孩的编号，countNum为游戏规则的间隔
func playGame(first *Boy, startNo, countNum int) {
	if first.next == nil {
		fmt.Println("就一个小孩玩游戏，玩不了")
		return
	}
	// 找到first的前一个节点，由于为环形链表最后一个节点的next就是first头结点，用于判断整个环形链表最后只剩下一个小孩了
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	// 找到游戏开始的小孩
	startNode := first
	for {
		if startNode.no == startNo {
			break
		}
		startNode = startNode.next
		// tail也要跟着移动
		tail = tail.next
	}
	for {
		// 一轮游戏
		for i := 1; i <= countNum; i++ {
			startNode = startNode.next
			tail = tail.next
		}
		fmt.Println("退出的小孩：", startNode.no)
		if startNode == tail {
			fmt.Println("最后剩下的小孩编号：", startNode.no)
			break
		}
		startNode = startNode.next
		tail.next = startNode
	}
}

func main() {
	// 构造五个节点的环形链表
	head := AddBoy(5)
	// 打印
	showBoy(head)
	// 游戏
	playGame(head, 3, 2)

}