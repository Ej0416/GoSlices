package main

import (
	"fmt"
)

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	message := [3]string{primary,secondary,tertiary}
	costPerMessage := [3]int{len(primary), len(secondary) + len(primary), len(tertiary) + len(secondary) + len(primary)}

	return  message, costPerMessage
}


func main(){
	fmt.Println("app start")
}