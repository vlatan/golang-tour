package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Goroutines
func Goroutine() {
	go say("world")
	say("Hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

// Channels
// ch <- v		Send v to channel ch.
// v := <-ch	Receive from ch, and assign value to v.
// The data flows in the direction of the arrow
// Like maps and slices, channels must be created before use:
// ch := make(chan int)
func Channel() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from channel c
	fmt.Println(x, y, x+y)
}

// Buffered Channels
func BufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// Range and Close
func RangeClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// Select
func Select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}

// Default Selection
func DefaultSelect() {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	// if node is nil exit
	if t == nil {
		return
	}

	// walk the left leaf
	Walk(t.Left, ch)

	// send the node (root) value to channel
	ch <- t.Value

	// walk the right leaf
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	// create channels
	ch1, ch2 := make(chan int), make(chan int)

	// Walk t1 in goroutine and close its channel when done
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	// Walk t2 in goroutine and close its channel when done
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1 // recieve value from channel ch1
		v2, ok2 := <-ch2 // recieve value from channel ch2

		// If both channels don't have the same state (opened or closed) OR
		// the values comming from both channels are not equal,
		// the trees are not equivalent.
		if ok1 != ok2 || v1 != v2 {
			return false
		}

		// when one of the channels is closed exit the loop
		if !ok1 {
			break
		}
	}

	return true
}

func testWalk(t *Tree) {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

// Exercise: Equivalent Binary Trees
func ExcerciseEquivalentBinaryTrees() {
	t1, t2 := New(1), New(2)
	testWalk(t1)
	fmt.Println(Same(t1, t2))
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()         // Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // Unlock upon this function exit
	return c.v[key]
}

// Mutex
func Mutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// Exercise: Web Crawler
func ExerciseWebCrawler() {
	Crawl("https://golang.org/", 4, fetcher)
}
