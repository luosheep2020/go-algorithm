package main

import "fmt"

type Study interface {
	learn()
}

type Student struct {
	name string
	book string
}

func (s Student) learn() {
	fmt.Printf("%s 在读 %s", s.name, s.book)
}
func main() {
	student := Student{
		name: "张三",
		book: "go语言入门",
	}

	student.learn()
}
