package leetcodeHot100

/*
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
*/

type LinkedNode struct {
	key, value int
	prev, next *LinkedNode
}
type LRUCache struct {
	capacity   int
	size       int
	m          map[int]*LinkedNode
	head, tail *LinkedNode
}

func Constructor(capacity int) LRUCache {
	lru := &LRUCache{
		size:     0,
		capacity: capacity,
		m:        make(map[int]*LinkedNode),
		head:     &LinkedNode{key: 0, value: 0, prev: nil, next: nil},
		tail:     &LinkedNode{key: 0, value: 0, prev: nil, next: nil},
	}
	//初始化双向链表
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return *lru
}

// Get ,找是否有key，如果有key的话，就把这个节点移动到双向链表的头部(所以我们要先实现把这个节点移到双向链表的头部)
func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

/*
//使用lru的代码
func main() {
	key := 0
	value := 0
	capacity := 10
	obj := Constructor(capacity)
	param_1 := obj.Get(key)
	fmt.Println(param_1)
	obj.Put(key, value)
}
*/
