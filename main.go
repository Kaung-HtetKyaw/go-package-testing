package gopackagetesting_test

import "fmt"

func TestingPackageFunc() {
	fmt.Println(("Hello from a package func"))
}

func main() {
	fmt.Println(("Hello from a package"))
}
