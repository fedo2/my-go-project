package main

import (
	"fmt"
	"github.com/fedo2/my-go-project/pkg/utils"
)

func main() {
	fmt.Println("Vítejte v aplikaci!")
	result := utils.Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", result)
}