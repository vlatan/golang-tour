package generics

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// Type parameters
func TypeParameters() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (head *List[T]) Append(data T) {
	newNode := List[T]{nil, data}
	if head.next == nil {
		head.next = &newNode
	} else {
		current := head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
}

func (head *List[T]) Display() {
	if head == nil {
		fmt.Println("Linked list is empty")
		return
	}
	current := head
	fmt.Println("Linked list: ")
	for current != nil {
		if current.next == nil {
			fmt.Printf("Value: %v, Next: %v\n", current.val, nil)
		} else {
			fmt.Printf("Value: %v, Next: %v\n", current.val, current.next.val)
		}
		current = current.next
	}
}

// Generic types
func GenericTypes() {
	list := List[int]{nil, 5}
	list.Append(10)
	list.Append(15)
	list.Append(20)
	list.Append(25)
	list.Display()
}
