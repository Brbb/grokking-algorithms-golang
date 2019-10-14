package main

import (
	"fmt"
	"testing"
)

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (ll *LinkedList) Append(data string) {
	newNode := Node{
		Data: data,
	}

	if ll.Length == 0 {
		ll.Head = &newNode
	} else {
		currentNode := ll.Head

		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}

		currentNode.Next = &newNode
	}

	ll.Length++
}

func (ll *LinkedList) Prepend(data string) {
	newNode := Node{
		Data: data,
	}

	if ll.Length == 0 {
		ll.Head = &newNode
	} else {
		currentNode := ll.Head
		ll.Head = &newNode
		ll.Head.Next = currentNode
	}

	ll.Length++

}

func (ll *LinkedList) Search(data string) *Node {
	if ll.IsEmpty() {
		return nil
	}

	currentNode := ll.Head

	for currentNode.Data != data {
		if currentNode.Next != nil {
			currentNode = currentNode.Next
		} else {
			return nil
		}
	}

	return currentNode
}

func (ll *LinkedList) SearchRecursively(head *Node, data string) *Node {
	if ll.IsEmpty() {
		return nil
	}

	if head.Data != data {
		if head.Next == nil {
			return nil
		}

		return ll.SearchRecursively(head.Next, data)
	}

	return head
}

// Overly verbose Delete
func (ll *LinkedList) Delete(data string) error {
	if ll.IsEmpty() {
		return fmt.Errorf("empty list")
	}

	currentNode := ll.Head
	var previousNode *Node

	// Get current and previous node
	for currentNode.Data != data {
		if currentNode.Next != nil {
			previousNode = currentNode
			currentNode = currentNode.Next
		} else {
			return fmt.Errorf("element not found")
		}
	}

	// If deleting first node
	if currentNode == ll.Head {
		// If list has more than one node
		if ll.Head.Next != nil {
			ll.Head = currentNode.Next
		} else {
			// If list has only one node
			ll.Head = nil
		}
		ll.Length--

		return nil
	}

	// If deleting last node
	if currentNode.Next == nil {
		previousNode.Next = nil
		ll.Length--
		return nil
	}

	// If deleting a middle node
	if previousNode != nil && currentNode.Next != nil {
		previousNode.Next = currentNode.Next
		ll.Length--
		return nil
	}

	return nil
}

func (ll *LinkedList) IsEmpty() bool {
	if ll.Length <= 0 {
		return true
	}

	return false
}

func (ll *LinkedList) String() string {
	if ll.IsEmpty() {
		return ""
	}

	currentNode := ll.Head
	var s string

	if currentNode.Next == nil {
		s += fmt.Sprintf("%s", currentNode.Data)
	}

	for currentNode.Next != nil {
		s += fmt.Sprintf(" | %s -> %s", currentNode.Data, currentNode.Next.Data)
		currentNode = currentNode.Next
	}

	return s
}

func Test_Append(t *testing.T) {
	ll := LinkedList{}

	// Append first node
	firstNodeData := "first node"
	ll.Append(firstNodeData)

	if ll.Length != 1 {
		t.Errorf("expected %d, got %d", 1, ll.Length)
	}

	if ll.Head.Data != firstNodeData {
		t.Errorf("expected %s, got %s", firstNodeData, ll.Head.Data)
	}

	// Append a new node
	secondNodeData := "second node"
	ll.Append(secondNodeData)

	if ll.Length != 2 {
		t.Errorf("expected %d, got %d", 2, ll.Length)
	}

	if ll.Head.Data != firstNodeData {
		t.Errorf("expected %s, got %s", firstNodeData, ll.Head.Data)
	}

	// Append a third node
	thirdNodeData := "third node"
	ll.Append(thirdNodeData)

	if ll.Length != 3 {
		t.Errorf("expected %d, got %d", 3, ll.Length)
	}

	if ll.Head.Data != firstNodeData {
		t.Errorf("expected %s, got %s", firstNodeData, ll.Head.Data)
	}
}

func Test_Prepend(t *testing.T) {
	ll := LinkedList{}

	// Preprend first node
	firstNodeData := "first node"
	ll.Prepend(firstNodeData)

	if ll.Length != 1 {
		t.Errorf("expected %d, got %d", 1, ll.Length)
	}

	if ll.Head.Data != firstNodeData {
		t.Errorf("expected %s, got %s", firstNodeData, ll.Head.Data)
	}

	// Preprend second node
	secondNodeData := "second node"
	ll.Prepend(secondNodeData)

	if ll.Length != 2 {
		t.Errorf("expected %d, got %d", 2, ll.Length)
	}

	if ll.Head.Data != secondNodeData {
		t.Errorf("expected %s, got %s", secondNodeData, ll.Head.Data)
	}
}

func Test_Search(t *testing.T) {
	ll := LinkedList{}

	// Search empty
	node := ll.Search("n1")
	if node != nil {
		t.Errorf("expected %v", nil)
	}

	// Search first item
	ll.Append("n1")
	ll.Append("n2")
	ll.Append("n3")

	node = ll.Search("n1")
	if node.Data != "n1" {
		t.Errorf("expected %s, got %s", "n1", node.Data)
	}

	// Search second item
	node = ll.Search("n2")
	if node.Data != "n2" {
		t.Errorf("expected %s, got %s", "n2", node.Data)
	}

	// Search non existing item
	node = ll.Search("n10")
	if node != nil {
		t.Errorf("expected %v, got %v", nil, node)
	}
}

func Test_SearchRecursively(t *testing.T) {
	ll := LinkedList{}

	// Search empty
	node := ll.SearchRecursively(ll.Head, "n1")
	if node != nil {
		t.Errorf("expected %v", nil)
	}

	// Search first item
	ll.Append("n1")
	ll.Append("n2")
	ll.Append("n3")

	node = ll.SearchRecursively(ll.Head, "n1")
	if node.Data != "n1" {
		t.Errorf("expected %s, got %s", "n1", node.Data)
	}

	// Search second item
	node = ll.SearchRecursively(ll.Head, "n2")
	if node.Data != "n2" {
		t.Errorf("expected %s, got %s", "n2", node.Data)
	}

	// Search non existing item
	node = ll.SearchRecursively(ll.Head, "n10")
	if node != nil {
		t.Errorf("expected %v, got %v", nil, node)
	}
}

func Test_IsEmpty(t *testing.T) {
	ll := LinkedList{}

	if ll.IsEmpty() == false {
		t.Errorf("expected %v, got %v", false, true)
	}

	// Append one item
	ll.Append("n1")

	if ll.IsEmpty() == true {
		t.Errorf("expected %v, got %v", true, false)
	}
}

func Test_DeleteNode(t *testing.T) {
	ll := LinkedList{}

	// Delete empty list
	err := ll.Delete("n1")
	if err.Error() != "empty list" {
		t.Errorf("expected %s, got %v", "empty list", nil)
	}

	// Delete non existing node
	ll = LinkedList{}
	ll.Append("n1")

	err = ll.Delete("n10")
	if err.Error() != "element not found" {
		t.Errorf("expected %s, got %v", "element not found", nil)
	}

	// Delete last element
	ll = LinkedList{}
	ll.Append("n1")
	ll.Append("n2")
	ll.Append("n3")

	err = ll.Delete("n3")
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if ll.Length != 2 {
		t.Errorf("expected %d, got %d", 2, ll.Length)
	}
	node := ll.Search("n3")
	if node != nil {
		t.Errorf("expected %v, got %v", nil, node)
	}
	node = ll.Search("n1")
	if node == nil && node.Data != "n1" {
		t.Errorf("expected %v, got %v", node, nil)
	}
	node = ll.Search("n2")
	if node == nil && node.Data != "n2" {
		t.Errorf("expected %v, got %v", node, nil)
	}

	// Delete first element
	ll = LinkedList{}
	ll.Append("n1")
	ll.Append("n2")
	ll.Append("n3")

	err = ll.Delete("n1")
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if ll.Length != 2 {
		t.Errorf("expected %d, got %d", 2, ll.Length)
	}
	node = ll.Search("n1")
	if node != nil {
		t.Errorf("expected %v, got %v", nil, node)
	}
	node = ll.Search("n2")
	if node == nil && node.Data != "n2" {
		t.Errorf("expected %v, got %v", node, nil)
	}
	node = ll.Search("n3")
	if node == nil && node.Data != "n3" {
		t.Errorf("expected %v, got %v", node, nil)
	}

	// Delete first (and only) element
	ll = LinkedList{}
	ll.Append("n1")

	err = ll.Delete("n1")
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if ll.Length != 0 {
		t.Errorf("expected %d, got %d", 2, ll.Length)
	}
	node = ll.Search("n1")
	if node != nil {
		t.Errorf("expected %v, got %v", nil, node)
	}

	// Delete middle element
	ll = LinkedList{}
	ll.Append("n1")
	ll.Append("n2")
	ll.Append("n3")

	err = ll.Delete("n2")
	if err != nil {
		t.Errorf("expected %v, got %v", nil, err)
	}
	if ll.Length != 2 {
		t.Errorf("expected %d, got %d", 2, ll.Length)
	}
	node = ll.Search("n2")
	if node != nil {
		t.Errorf("expected %v, got %v", nil, node)
	}
	node = ll.Search("n1")
	if node == nil && node.Data != "n1" {
		t.Errorf("expected %v, got %v", node, nil)
	}
	node = ll.Search("n3")
	if node == nil && node.Data != "n3" {
		t.Errorf("expected %v, got %v", node, nil)
	}
}

func Test_String(t *testing.T) {
	ll := LinkedList{}

	s := ll.String()
	expected := ""

	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}

	ll.Append("n1")
	s = ll.String()
	expected = "n1"

	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}

	ll.Append("n2")
	ll.Append("n3")

	s = ll.String()
	expected = " | n1 -> n2 | n2 -> n3"

	if s != expected {
		t.Errorf("expected %s, got %s", expected, s)
	}

}
