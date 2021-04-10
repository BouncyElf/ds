package problem

/*
问题：
138.复制带随机指针的链表
https://leetcode-cn.com/problems/copy-list-with-random-pointer/
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 思路1: 使用map映射老指针与新指针
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}
	res, now, m := new(Node), new(Node), make(map[*Node]*Node)
	for temp := head; temp != nil; temp = temp.Next {
		if temp == head {
			res.Val = temp.Val
			now = res
		} else {
			now.Next = &Node{
				Val: temp.Val,
			}
			now = now.Next
		}
		m[temp] = now
	}
	for origin, new := head, res; origin != nil; origin, new = origin.Next, new.Next {
		if origin.Random == nil {
			continue
		}
		new.Random = m[origin.Random]
	}
	return res
}

// 思路2: 直接将老的指针插入到新的指针之后
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	for temp := head; temp != nil; {
		next := temp.Next
		temp.Next = &Node{
			Val:  temp.Val,
			Next: next,
		}
		temp = next
	}
	for temp := head; temp != nil && temp.Next != nil; temp = temp.Next.Next {
		if temp.Random == nil {
			continue
		}
		temp.Next.Random = temp.Random.Next
	}
	res := head.Next
	for temp := head; temp.Next != nil; {
		next := temp.Next
		temp.Next = temp.Next.Next
		temp = next
	}
	return res
}
