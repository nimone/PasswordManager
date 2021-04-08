package helper

import (
	"fmt"
)

func HandleError(err error, customMsg ...string) bool {
	if err != nil {
		fmt.Println("Error:", err)

		for _, msg := range customMsg {
			fmt.Println(" ", msg)
		}
		return true
	}
	return false
}