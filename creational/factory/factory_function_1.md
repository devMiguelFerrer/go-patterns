# Factory Pattern

## Description
In progress ...

```go
package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Adult bool
}

func NewPerson(name string, age int) *Person {
	adult := age > 18
	return &Person{
		name, age, adult,
	}
}

func main() {
	p := NewPerson("Michael", 33)
	fmt.Println(p)
}
```