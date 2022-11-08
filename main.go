package main

import (
	"runtime"
	"unsafe"
)

type animal struct {
	name string
}

//go:noinline
func newBear() *animal {
	return &animal{name: "yogi"}
}

//go:noinline
func newCat() *animal {
	return &animal{name: "garfield"}
}

//go:noinline
func (a *animal) printName() {
	println("a", uintptr(unsafe.Pointer(a)))
	println(a.name)
}

//export tinygo_getCurrentStackPointer
func getCurrentStackPointer() uintptr

//go:extern __global_base
var globalsStartSymbol [0]byte

//go:noinline
func dumpStack() {
	stackBottom := uintptr(65136) // getCurrentStackPointer == stackTop in this program
	stackTop := uintptr(unsafe.Pointer(&globalsStartSymbol))

	println("stackBottom", stackBottom)
	println("stackTop", stackTop)

	for i := stackBottom; i < stackTop; i += 4 {
		println(i, *(*uintptr)(unsafe.Pointer(i)))
	}
}

//go:noinline
func run() {
	b := newBear()
	println("b", uintptr(unsafe.Pointer(b)))

	for i := 0; i < 100000; i++ {
		runtime.GC()
		newCat()
	}

	c := newCat()
	println("c", uintptr(unsafe.Pointer(c)))

	b.printName()

	c.printName()
}

func main() {
	run()
}
