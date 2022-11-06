# Factory Pattern

## Description
In progress ...

```go
package main

import "fmt"

type Person interface {
	SayHello() string
}
type person struct {
	name  string
	age   int
	adult bool
}

func (p *person) SayHello() string {
	return fmt.Sprintf("Hi, my name is %s", p.name)
}

func NewPerson(name string, age int) Person {
	adult := age > 18
	return &person{
		name, age, adult,
	}
}

func main() {
	p := NewPerson("Michael", 33)
	fmt.Println(p)
	fmt.Println(p.SayHello())
}

```