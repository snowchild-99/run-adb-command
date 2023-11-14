package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello Sandeep")

	outputArray, err := exec.Command("pm list packages | grep tv.cloudwalker").Output()
	myString := string(outputArray)
	fmt.Println(myString)
	fmt.Println(err)

}
