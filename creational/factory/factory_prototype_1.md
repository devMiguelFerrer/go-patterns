# Factory Pattern

## Description
In progress ...

```go
package main

import "fmt"

const (
	A01 = iota
	B01
)

type Student struct {
	class, tutor, name string
}

func NewStudentFactory(class int, name string) *Student {
	switch class {
	case A01:
		return &Student{
			"A01", "Michael", name,
		}
	case A02:
		return &Student{
			"A01", "John", name,
		}
	case B01:
		return &Student{
			"A01", "James", name,
		}
	default:
		panic("Unavailable Class")
	}
}

func main() {
	s1 := NewStudentFactory(A01, "Peter")
	s2 := NewStudentFactory(B01, "Martina")
	fmt.Println(s1, s2)
}
```