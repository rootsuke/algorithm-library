package basic_data_structure

type Node struct {
	key  int
	next *Node
	prev *Node
}

func NewNode() *Node {
	node := &Node{
		next: nil,
		prev: nil,
	}

	node.next = node
	node.prev = node

	return node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	node := NewNode()
	return &LinkedList{head: node}
}

func (dll *LinkedList) Insert(key int) {
	next := dll.head.next
	node := &Node{key: key, next: next, prev: dll.head}
	dll.head.next = node
	next.prev = node
}

func (dll *LinkedList) Delete(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dll *LinkedList) DeleteKey(key int) {
	for n := dll.head.next; n != dll.head; n = n.next {
		if n.key == key {
			dll.Delete(n)
			return
		}
	}
}

func (dll *LinkedList) DeleteFirst() {
	dll.Delete(dll.head.next)
}

func (dll *LinkedList) DeleteLast() {
	dll.Delete(dll.head.prev)
}
