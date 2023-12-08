package main

import "fmt"

func main() {
	// s:=Square{
	// 	Length: 2,
	// }
	// c:=Circle{
	// 	Radius: 3,
	// }
	// PrintShapePerimeter(s)
	// PrintShapeSurface(s)
	// PrintShapePerimeter(c)
	// PrintShapeSurface(c)

	// d:=Dog{
	// 	Sound: "A",
	// }
	// c:=Cat{
	// 	Sound: "B",
	// }
	// h:=Hourse{
	// 	Sound: "D",
	// }
	// PrintSound(d)
	// PrintSound(c)
	// PrintSound(h)

	
}

type Dog struct {
	Sound string
}
type Cat struct {
	Sound string
}
type Hourse struct {
	Sound string
}
type Animal interface{
	PrintSound()
}

func PrintSound(a Animal){
	a.PrintSound()

}
func (d Dog)PrintSound(){
	fmt.Println("A")
}
func (c Cat)PrintSound(){
	fmt.Println("B")
}
func (h Hourse)PrintSound(){
	fmt.Println("D")
}












type Square struct {
	Length int
}
type Circle struct {
	Radius int
}
type Shape interface {
	Surface()
	Perimeter()
}

func PrintShapeSurface(s Shape) {
	s.Surface()
}
func PrintShapePerimeter(s Shape) {
	s.Perimeter()
}
func (s Square) Surface() {
	fmt.Println("surface of square ", s.Length*s.Length)
}
func (c Circle) Surface() {
	fmt.Println("surface of circle , ", c.Radius*3)
}
func (s Square) Perimeter() {
	fmt.Println("perimeter of square ", 4*s.Length)
}
func (c Circle) Perimeter() {
	fmt.Println("perimeter  of circle , ", c.Radius*3*2)
}
