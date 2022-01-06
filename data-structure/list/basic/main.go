package main

import "log"

type Element interface{}

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Header *Node // 头节点
}

func main() {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")

	//log.Println(list.GetLength())
	//list.Scan()

	list.Add("head node data")

	list.Insert(5, "five_insert_value")
	list.Remove("c")
	list.RemoveByPos(3)
	list.Scan()
	log.Println(list.GetLength())


}

// CreateNode
func CreateNode(v interface{}) *Node {
	return &Node{v, nil}
}

// IsEmpty
func (l *List) IsEmpty() bool {
	if l.Header == nil {
		return true
	} else {
		return false
	}
}

// GetLength
func (l *List) GetLength() int {
	cur := l.Header

	listLen := 0
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	return listLen
}

// Add 往链表表头增加一个节点
func (l *List) Add(data Element) {
	if data == nil {
		log.Println("data is nil")
		return
	}

	newNode := CreateNode(data)
	newNode.Next = l.Header
	l.Header = newNode
}

// Append 尾部增加元素
func (l *List) Append(data Element) {
	if data == nil {
		log.Println("data is nil")
		return
	}

	newNode := CreateNode(data)
	if l.IsEmpty() {
		l.Header = newNode
	} else {
		cur := l.Header
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = newNode
	}
}

// Insert 在指定位置添加元素
func (l *List) Insert(pos int, data Element) {
	if pos < 0 {
		l.Add(data)
	}else if pos > l.GetLength() {
		l.Append(data)
	} else {
		pre := l.Header
		cnt := 0
		for cnt < (pos - 1) {
			pre = pre.Next
			cnt++
		}

		newNode := CreateNode(data)
		newNode.Next = pre.Next	// 将新的链表节点指向 pre.Next所指向的节点
		pre.Next = newNode // 上一个链表所指向的next node 改为 newNode
	}
}

// Remove
func (l *List) Remove(data Element) {
	pre := l.Header
	if pre.Data == data {
		l.Header = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// RemoveByPos
func (l *List) RemoveByPos(pos int) {
	if pos < 0 {
		log.Println("pos err")
		return
	}

	if pos > l.GetLength() {
		log.Println("pos out of length")
		return
	}

	pre := l.Header
	if pos == 0 {
		l.Header = pre.Next
	} else {
		cnt := 0
		for cnt != (pos - 1) && pre.Next != nil {
			cnt++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// Scan
func (l *List) Scan() {
	if !l.IsEmpty() {
		cur := l.Header
		for cur != nil {
			log.Printf("val:%v\n", cur.Data)
			cur = cur.Next
		}
	}
}
