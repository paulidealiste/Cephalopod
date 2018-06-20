// Package cephalostructures implements composite data structures such as trees and graphs
package cephalostructures

import (
	"github.com/paulidealiste/Cephalopod/cephaloutils"
)

// STACK STRUCTURE //

// StackMethods describes common stack methods
type StackMethods interface {
	Empty() bool
	Push(item interface{})
	Pop() interface{}
	Peek() interface{}
	Size() bool
}

// Stack implementation of a FILO principle
type Stack struct {
	Items []interface{}
}

// Empty returns whether current stack is empty
func (st *Stack) Empty() bool {
	return len(st.Items) == 0
}

// Push adds a new item onto/ontop a stack
func (st *Stack) Push(item interface{}) {
	st.Items = append(st.Items, item)
}

// Pop removes the top element from the stack and returns it
func (st *Stack) Pop() interface{} {
	popped := st.Items[len(st.Items)-1]
	st.Items = st.Items[:len(st.Items)-1]
	return popped
}

// Peek return the top element without deleting it from the stack
func (st *Stack) Peek() interface{} {
	peeked := st.Items[len(st.Items)-1]
	return peeked
}

// Size return the current stack size i.e. the number of elements
func (st *Stack) Size() int {
	return len(st.Items)
}

// QUEUE STRUCTURE //

// QueueMethods describes common queue methods
type QueueMethods interface {
	Empty() bool
	Enqueue(item interface{})
	Dequeue() interface{}
	Size() int
}

// Queue implementation of a FIFO principle
type Queue struct {
	Items []interface{}
}

// Empty returns whether current queue is empty
func (qu *Queue) Empty() bool {
	return len(qu.Items) == 0
}

// Enqueue adds the item at the rear/start of the Queue
func (qu *Queue) Enqueue(item interface{}) {
	qu.Items = append([]interface{}{item}, qu.Items...)
}

// Dequeue removes first added element - removes from face/end
func (qu *Queue) Dequeue() interface{} {
	var dequeued interface{}
	dequeued, qu.Items = qu.Items[len(qu.Items)-1], qu.Items[:len(qu.Items)-1]
	return dequeued
}

// Size returns the current queue size
func (qu *Queue) Size() int {
	return len(qu.Items)
}

// NETWORK/TREE COMMON //

// SepiaNode defines one node(vertex) of the network as a common data type. It contains
// data with any type for storage, an internal id, and a connections map, using the id
// as the key.
type SepiaNode struct {
	id    int
	key   string
	Data  interface{}
	Title string
}

// SepiaNodeMethods interface defines common methods for any network-like structure.
// InsertAdjacent should add new node as a connection of the node in question.
type SepiaNodeMethods interface {
	InsertAdjacent(node *SepiaNode)
}

// GRAPH STRUCTURE //

// GraphNode defines one vertex of the graph. It embeds SepiaNode common while adding
// Connections map wich defines all of the immediate neighbours of the node in question.
type GraphNode struct {
	SepiaNode
	Connections map[string]*GraphNode
}

// InsertAdjacent adds a new node(vertex) as a direct connection to the node in question.
func (gn *GraphNode) InsertAdjacent(n *GraphNode) {
	gn.Connections[n.key] = n
}

// Graph defines a graph representation of a network structure. Graph is implemented
// by an adjacency list.
type Graph struct {
	NodeList  []*GraphNode
	NodeCount int
}

// InsertNode performs an inserting operation of the new node to the overall graph
func (gr *Graph) InsertNode(key string, title string, data interface{}) {
	gr.NodeCount++
	nno := &GraphNode{
		SepiaNode: SepiaNode{
			id:    cephaloutils.RandomID(),
			key:   key,
			Data:  data,
			Title: title,
		},
		Connections: make(map[string]*GraphNode),
	}
	gr.NodeList = append(gr.NodeList, nno)
}

// RemoveNode removes the node, as well as all its related edges.
func (gr *Graph) RemoveNode(key string) {
	var survivors int
	for _, gn := range gr.NodeList {
		if gn.key != key {
			gr.NodeList[survivors] = gn
			survivors++
		}
		delete(gn.Connections, key)
	}
	gr.NodeList = gr.NodeList[:survivors]
}

// GraphCatalog returns keyed map/list of all the nodes
func (gr *Graph) GraphCatalog() map[string]*GraphNode {
	nodemap := make(map[string]*GraphNode)
	for _, gn := range gr.NodeList {
		nodemap[gn.key] = gn
	}
	return nodemap
}

// DirectedEdge adds edge between the two nodes of choice, form source to target.
func (gr *Graph) DirectedEdge(source *GraphNode, target *GraphNode) {
	source.InsertAdjacent(target)
}

// UndirectedEdge adds both edges between the nodes of choice.
func (gr *Graph) UndirectedEdge(n1 *GraphNode, n2 *GraphNode) {
	n1.InsertAdjacent(n2)
	n2.InsertAdjacent(n1)
}

// TREE STRUCTURE //

// Tree is a basic binary tree-like structure
type Tree struct {
	root  *SepiaNode
	left  *Tree
	right *Tree
}

// SetRootNode sets the new data/key generic node as a tree root
func (tr *Tree) SetRootNode(key string, title string, data interface{}) {
	tr.root = createSepiaNode(key, title, data)
}

// InsertLeft is a generic insert left node/tree of the tree
func (tr *Tree) InsertLeft(key string, title string, data interface{}) {
	tr.left = &Tree{
		root: createSepiaNode(key, title, data),
	}
}

// InsertRight is a generic insert left node/tree of the tree
func (tr *Tree) InsertRight(key string, title string, data interface{}) {
	tr.right = &Tree{
		root: createSepiaNode(key, title, data),
	}
}

func createSepiaNode(key string, title string, data interface{}) *SepiaNode {
	return &SepiaNode{
		id:    cephaloutils.RandomID(),
		key:   key,
		Data:  data,
		Title: title,
	}
}
