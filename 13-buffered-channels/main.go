package main

import "fmt"

type Message struct {
	Id int
	Body []byte
}

func main() {

	ch := make(chan Message, 2)
	ch <- Message{1, []byte {99,88}}
	ch <- Message{2, []byte {11,22}}

	msg1 := <- ch
	fmt.Println(msg1)

	msg2 := <- ch
	fmt.Println(msg2)


}
