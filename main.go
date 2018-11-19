package main

import "fmt"

import "github.com/myitcvscratch/food"

func main() {
	fmt.Println("Simple module-based main v1.0.0")
	fmt.Printf("Today we will eat %v\n", food.MainCourse)
}
