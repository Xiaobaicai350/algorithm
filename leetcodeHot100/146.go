package leetcodeHot100

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

// 把节点移动到双向链表的头部
func (this *LRUCache) moveToHead(node *LinkedNode) {
	//删除当前节点
	this.removeNode(node)
	this.addToHead(node)
}

// Get ,找是否有key，如果有key的话，就把这个节点移动到双向链表的头部(所以我们要先实现把这个节点移到双向链表的头部)
func (this *LRUCache) Get(key int) int {

	if node, ok := this.m[key]; ok {
		this.moveToHead(node)
		return node.value
	} else { //如果没有直接返回-1
		return -1
	}
}

// Put 往lru中添加/更新数据
func (this *LRUCache) Put(key int, value int) {
	//先判断是新增还是更新
	if node, ok := this.m[key]; ok { //如果存在，说明是更新
		//更新的话，也是一次访问
		node.value = value
		//移动到头部
		this.moveToHead(node)
	} else { //如果不存在，说明是添加
		node := &LinkedNode{key: key, value: value, prev: nil, next: nil}
		//添加到链表和map
		this.m[key] = node
		//添加到双向链表头部
		this.addToHead(node)
		this.size++
		if this.size > this.capacity { //如果超过容量了，需要进行淘汰策略了
			//删除双向链表中的尾节点
			tail := this.deleteTail()
			//删除map的对应节点
			delete(this.m, tail.key)
			this.size--
		}
	}
}

// 删除当前节点
func (this *LRUCache) removeNode(node *LinkedNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

// 把当前节点加到链表头部
func (this *LRUCache) addToHead(node *LinkedNode) {
	//一共要改四个指针，一个原来的head的next的prev,一个是head的next，两个node的指针
	next := this.head.next
	next.prev = node
	this.head.next = node
	node.prev = this.head
	node.next = next
}

func (this *LRUCache) deleteTail() *LinkedNode {
	node := this.tail.prev
	this.tail.prev = node.prev
	node.prev.next = this.tail
	return node
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
