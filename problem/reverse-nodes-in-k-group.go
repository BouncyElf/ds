package problem

/*
问题：
25.K个一组翻转链表
https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
*/

// 思路是每遇见k个就翻转一下，每次翻转后把上一次翻转后的尾节点指向当前新的头
// 当前新的尾指向正常遍历的下一个节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	newHead, last, counter := head, (*ListNode)(nil), 0
	for now := head; now != nil; now = now.Next {
		counter++
		if counter == k {
			counter = 0
			if newHead == head {
				newHead = now
			}
			start := head
			if last != nil {
				start = last.Next
				last.Next = now
			}
			term, prev := now.Next, now.Next
			for temp := start; temp != nil && temp != term; {
				next := temp.Next
				temp.Next = prev
				prev = temp
				temp = next
			}
			last = start
			now = start
		}
	}
	return newHead
}
