package level1

import "fmt"

type LinkList struct {
	Val  interface{}
	Next *LinkList
}

// 打印链表
func (linkList *LinkList) PrintList() {
	if linkList == nil {
		return
	}
	fmt.Println(linkList.Val)
	linkList.Next.PrintList()
}

func (linkList *LinkList) ToList() []interface{} {
	var res []interface{}
	if linkList == nil {
		return res
	}

	res = append(res, linkList.Val)
	res = append(res, linkList.Next.ToList()...)
	return res
}

// 新建链表
func InitLinkList(val interface{}) *LinkList {
	return &LinkList{
		val, nil,
	}
}
