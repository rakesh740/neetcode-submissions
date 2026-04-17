type LRUCache struct {
	cap int
	cache map[int]*Node
    head *Node
	tail *Node
}

type Node struct {
	left *Node
	right *Node
	key , val int
}


func (this *LRUCache) Insert(node *Node){
	// insert after head 
	head , next := this.head, this.head.right
	head.right = node
	next.left = node
	node.left = head
	node.right = next
}

func (this *LRUCache) Remove(node *Node) {
	prev, next := node.left, node.right
	prev.right = next
	next.left = prev
}

func Constructor(capacity int) LRUCache {
    c := LRUCache{
		cap: capacity, 
		cache: make(map[int]*Node, capacity),
		head: &Node{}, 
		tail: &Node{},
	}
	c.head.right = c.tail
	c.tail.left = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
    // get from hash map 
	// delete last node create new node and add to head
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.Remove(node)
	this.Insert(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
    
	// put to hash map 
	// check exceed capacity then delete from last from list 
	// add to head 

	node, ok := this.cache[key]
	if ok {
		this.Remove(node)
		delete(this.cache, key)
	}
	node = &Node{key: key, val: value}
	this.Insert(node)
	this.cache[key] = node

	if len(this.cache) > this.cap {
		last := this.tail.left
		delete(this.cache, last.key)
		this.Remove(last)
	}

}

