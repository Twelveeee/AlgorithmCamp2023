package level3

// LRU 缓存
// https://leetcode.cn/problems/lru-cache/
// map + 双向链表实现 LRU 缓存
type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*LinkNode
	head, tail *LinkNode
}

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

func Constructor(capacity int) LRUCache {
	dummyHeadNode, dummyTailNode := &LinkNode{0, 0, nil, nil}, &LinkNode{0, 0, nil, nil}
	dummyHeadNode.next = dummyTailNode
	dummyTailNode.pre = dummyHeadNode
	return LRUCache{
		size:  0,
		cap:   capacity,
		cache: map[int]*LinkNode{},
		head:  dummyHeadNode,
		tail:  dummyTailNode,
	}
}

// 每次get都会把操作的节点移动到链表的头部
func (ca *LRUCache) Get(key int) int {
	node, ok := ca.cache[key]
	if !ok {
		return -1
	}

	out := node.val
	// 移动到head
	node.delete()
	ca.addNodeToHead(node)

	return out
}

func (ca *LRUCache) Put(key int, value int) {
	node, ok := ca.cache[key]
	if !ok {
		// insert
		if ca.size >= ca.cap {
			// 尾删
			tailNode := ca.tail.pre
			tailNode.delete()
			delete(ca.cache, tailNode.key)
			ca.size--
		}
		newNode := &LinkNode{key, value, nil, nil}
		ca.addNodeToHead(newNode)
		ca.cache[key] = newNode
		ca.size++
	} else {
		// update
		node.val = value
		node.delete()
		ca.addNodeToHead(node)
	}

}

func (ca *LRUCache) addNodeToHead(node *LinkNode) {
	node.pre = ca.head
	node.next = ca.head.next

	ca.head.next.pre = node
	ca.head.next = node
}

func (node *LinkNode) delete() {
	node.next.pre = node.pre
	node.pre.next = node.next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
