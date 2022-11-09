# Prototype Pattern

## Description
In progress ...

```go
package main

import (
	"fmt"
)

type Student struct {
	class, tutor, name string
	age                int
	subjects           []string
}

func (s *Student) Clone() *Student {
	return &Student{class: s.class, tutor: s.tutor, subjects: s.subjects}
}

var principalClass = Student{
	class:    "Principal",
	tutor:    "Michael",
	subjects: []string{"science", "math", "geography"},
}

var a001Class = Student{
	class:    "A001",
	tutor:    "Peter",
	subjects: []string{"art", "music", "history"},
}

func newStudent(prototype *Student, name string, age int) *Student {
	cp := prototype.Clone()
	cp.name = name
	cp.age = age
	return cp
}

func newPrincipalStudent(name string, age int) *Student {
	return newStudent(&principalClass, name, age)
}

func newA001Student(name string, age int) *Student {
	return newStudent(&a001Class, name, age)
}

func main() {
	s1 := newPrincipalStudent("John", 14)
	s2 := newA001Student("Martina", 15)
	fmt.Println(s1, s2)
}
```