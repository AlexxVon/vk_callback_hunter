package main

import "fmt"
func getEnteringCodeStatus (peerId float64) bool {

	var b bool

	fmt.Println("userEnterCode", userEnterCode)

	for j, _ := range (userEnterCode) {

		fmt.Println(j, "   ", peerId)
		if j == peerId {
			 b = true
			 
		} else {
			b = false
		}
	}
	return b

}