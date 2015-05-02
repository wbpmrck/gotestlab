
/*
场景:

如果一个struct(A)匿名嵌套了2个struct (B,C),B和C内部又有同名的method(foo),那么我调用A.foo的时候，实际调用的是B还是C的？

结论

结果是编译不通过,go编译器不允许这种情况，报错“ambiguous selector”
*/
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
