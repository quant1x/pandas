package tests

import "fmt"

type Person struct {
}

func (this *Person) Eat() {
	fmt.Println("Person Eat")
}

func (this *Person) Run() {
	fmt.Println("Person Run")
}

func (this *Person) Sleep() {
	fmt.Println("Person Sleep")
}

type Man struct {
	Person
}

func (this *Man) Eat() {
	// 类似于Java的 super.Eat()
	this.Person.Eat()
	fmt.Println("Man Eat")
}

func (this *Man) Run() {
	fmt.Println("Man Run")
}

// 抽象的用法: 函数指针

type AbstractDog struct {
	Sleep func()
}

func (this *AbstractDog) Eat() {
	fmt.Println("AbstractDog Eat")
	this.Sleep()
}

func (this *AbstractDog) Run() {
	fmt.Println("AbstractDog Run")
}

// Akita 秋田犬
type Akita struct {
	AbstractDog
}

func NewAkita() *Akita {
	ptr := &Akita{}
	ptr.AbstractDog.Sleep = ptr.Sleep
	return ptr
}

func (this *Akita) Sleep() {
	fmt.Println("Akita Sleep")
}

// Labrador 拉布拉多犬
type Labrador struct {
	AbstractDog
}

func NewLabrador() *Labrador {
	ptr := &Labrador{}
	ptr.AbstractDog.Sleep = ptr.Sleep
	return ptr
}

func (this *Labrador) Sleep() {
	fmt.Println("Labrador Sleep")
}
