package problem

import (
	"container/list"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for now := l; now != nil; now = now.Next {
		fmt.Print(now.Val)
		if now.Next != nil {
			fmt.Print(" -> ")
		} else {
			fmt.Println()
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) GetVal() string {
	if node == nil {
		return "<nil>"
	}
	return strconv.Itoa(node.Val)
}

func (node *TreeNode) String() string {
	return fmt.Sprintf("Val:%s, Left:%s, Right:%s",
		node.GetVal(), node.Left.GetVal(), node.Right.GetVal())
}

func buildTree(data []int, now int) *TreeNode {
	if now < 0 || now >= len(data) {
		return nil
	}
	node := &TreeNode{
		Val:   data[now],
		Left:  buildTree(data, now*2+1),
		Right: buildTree(data, now*2+2),
	}
	return node
}

func (head *TreeNode) Print() {
	if head == nil {
		return
	}
	q := list.New()
	q.PushBack(head)
	for level := 0; q.Len() != 0; level++ {
		length := q.Len()
		for range make([]struct{}, length) {
			now := q.Remove(q.Front()).(*TreeNode)
			fmt.Println("level:", level, now)
			if now.Left != nil {
				q.PushBack(now.Left)
			}
			if now.Right != nil {
				q.PushBack(now.Right)
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
