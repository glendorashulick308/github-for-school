package main

import "fmt"

func main() {
    // Define a function to perform some operation on two numbers
    func add(x int, y int) int {
        return x + y
    }

    // Call the function and print the result
    sum := add(5, 3)
    fmt.Println(sum)
}
