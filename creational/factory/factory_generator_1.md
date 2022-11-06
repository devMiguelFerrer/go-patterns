# Factory Pattern

## Description
In progress ...

```go
package main

import "fmt"

type Student struct {
	class, tutor, name string
}
type StudentFactory struct {
	class, tutor string
}

func (st *StudentFactory) Create(name string) *Student {
	return &Student{st.class, st.tutor, name}
}

func NewStudentFactory(class, tutor string) *StudentFactory {
	return &StudentFactory{
		class, tutor,
	}
}

func main() {
	stA01 := NewStudentFactory("A01", "John")
	s1 := stA01.Create("Michael")
	s2 := stA01.Create("Peter")
	fmt.Println(s1, s2)
}
```