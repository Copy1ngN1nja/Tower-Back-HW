package main

import "fmt"

type Node struct {
	val int
	l   *Node
	r   *Node
}

func NewNode(x int) Node {
	return Node{
		val: x,
		l:   nil,
		r:   nil,
	}
}

func (cur Node) IsExist(x int) bool {
	switch {
	case x == cur.val:
		return true
	case x < cur.val:
		if cur.l != nil {
			return (*cur.l).IsExist(x)
		} else {
			return false
		}
	case x > cur.val:
		if cur.r != nil {
			return (*cur.r).IsExist(x)
		} else {
			return false
		}
	}
	return true
}

func (cur *Node) Add(x int) {
	switch {
	case x < cur.val:
		if cur.l != nil {
			cur.l.Add(x)
		} else {
			tmp := NewNode(x)
			cur.l = &tmp
		}
	case x > cur.val:
		if cur.r != nil {
			cur.r.Add(x)
		} else {
			tmp := NewNode(x)
			cur.r = &tmp
		}
	}
}

func (cur *Node) FindFarChild() *Node {
	cur = cur.l
	for cur.r != nil {
		cur = cur.r
	}
	return cur
}

func (cur *Node) Delete(x int) *Node {
	if cur == nil {
		return nil
	}
	switch {
	case x == cur.val:
		if cur.l == nil {
			r_child := cur.r
			cur = nil
			return r_child
		} else if cur.r == nil {
			l_child := cur.l
			cur = nil
			return l_child
		} else {
			far_child := cur.FindFarChild()
			cur.val = far_child.val
			cur.l = cur.l.Delete(far_child.val)
		}
	case x < cur.val:
		cur.l = cur.l.Delete(x)
	case x > cur.val:
		cur.r = cur.r.Delete(x)
	}
	return cur
}

func main() {
	root := NewNode(50)
	root.Add(52)
	root.Add(40)
	root.Add(51)
	fmt.Println(root.val)
	fmt.Println(root.l.val)
	fmt.Println(root.r.val)
	fmt.Println(root.r.l.val)
	fmt.Println(root.IsExist(50))
	fmt.Println(root.IsExist(60))
	root.Add(60)
	fmt.Println(root.IsExist(60))
	root.Delete(52)
	fmt.Println(root.r.r.val)
}
