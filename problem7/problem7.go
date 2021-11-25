package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type singleList struct {
	len  int
	head *Node
}

func initList() *singleList {
	return &singleList{}
}



func (s *singleList) AddBack(val int) {
	Node := &Node{
		val: val,
	}
	if s.head == nil {
		s.head = Node
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = Node
	}
	s.len++
	return
}


func (s *singleList) Front() (int, error) {
	if s.head == nil {
		return -1, fmt.Errorf("Single List is empty")
	}
	return s.head.val, nil
}

func (s *singleList) Size() int {
	return s.len
}

func (s *singleList) Traverse() error {
	if s.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
	return nil
}

func main() {
	singleList := initList()
	fmt.Printf("AddFront: A\n")
	singleList.AddFront(2)
	fmt.Printf("AddFront: B\n")
	singleList.AddFront(1)
	fmt.Printf("AddBack: C\n")
	singleList.AddBack(4)

	fmt.Printf("Size: %d\n", singleList.Size())

	err := singleList.Traverse()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("RemoveFront\n")
	err = singleList.RemoveFront()
	if err != nil {
		fmt.Printf("RemoveFront Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	fmt.Printf("RemoveBack\n")
	err = singleList.RemoveBack()
	if err != nil {
		fmt.Printf("RemoveBack Error: %s\n", err.Error())
	}

	err = singleList.Traverse()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Size: %d\n", singleList.Size())
}
