package leetcodeHot100

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

// 定义链表节点
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 创建链表节点
func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

// 新建LRU缓存，传入的参数是容量
func NewLRUCache(capacity int) *LRUCache {
	l := &LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkedNode),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	//然后把头结点的下一个节点置为尾节点
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

// Get方法，找是否有key，如果有key的话，就把这个节点移动到双向链表的头部
func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

// Put方法，先看这个缓存里有没有key，如果有的话，就更新，如果没有的话，就新建
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok { //如果有的话，就进行更新操作，并且移动到双向链表的头部，因为这次也相当于访问
		node.value = value
		this.moveToHead(node)
	} else { //如果没有的话，就新建系欸但
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	}
}

// 把节点添加到双向链表的头部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 删除双向链表中的node
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 把节点移动到双向链表的头部
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 删除尾巴节点
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
