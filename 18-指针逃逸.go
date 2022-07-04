package main

type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) // 局部变量s逃逸到堆
	s.Name = name
	s.Age = age
	return s
}

func main() {
	s := StudentRegister("Kobe", 18)
	println(s.Name, ":", s.Age)
}
