package main

import "fmt"


type B struct{
}
func (this *B) foo() string{
  return "B.foo called!"
}


type C struct{
}
func (this *C) foo() string{
  return "C.foo called!"
}

type A struct{
   B
   C
}

type A2 struct{
   C
   B
}

func main() {
    a := &A{}
    a2 := &A2{}
	fmt.Println("a.foo is"+a.foo())
	fmt.Println("a2.foo is"+a2.foo())
}
