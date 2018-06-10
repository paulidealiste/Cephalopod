package cephalostructures

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	testqueue := Queue{}
	fmt.Println(testqueue.Empty())
	testqueue.Enqueue(5.8)
	testqueue.Enqueue(8.9)
	testqueue.Enqueue(7.5)
	fmt.Println(testqueue.Items)
	testqueue.Dequeue()
	fmt.Println(testqueue.Items)
	fmt.Println(testqueue.Size())
}

func TestGraph(t *testing.T) {
	testgraph := Graph{}
	testgraph.InsertNode("kxp", "lush", [5]int{1, 2, 3, 4, 5})
	testgraph.InsertNode("exp", "soap", [5]int{5, 6, 7, 9, 10})
	testgraph.InsertNode("pxp", "ginger", [5]int{10, 11, 12, 13, 14})
	testgraph.InsertNode("rxp", "cinnamon", [5]int{14, 15, 16, 17, 18})
	testcatalog := testgraph.GraphCatalog()
	testgraph.DirectedEdge(testcatalog["kxp"], testcatalog["exp"])
	testgraph.UndirectedEdge(testcatalog["pxp"], testcatalog["rxp"])
	testgraph.DirectedEdge(testcatalog["kxp"], testcatalog["exp"])
	testgraph.RemoveNode("pxp")
	testcatalog = testgraph.GraphCatalog()
	fmt.Println(testcatalog)
	fmt.Println(testcatalog["rxp"])
	testgraph.InsertNode("mxp", "cinnamon", [5]int{20, 25, 30, 31, 32})
}
