package main

import (
	"fmt"
	"testing"
)

func TestGetMessage(t *testing.T) {
	testCases := []string{"🗺️","👍","👏","😋👌","🐔"}

	for i, testCase := range testCases {

		t.Run(fmt.Sprintf("%v", testCase), func(t *testing.T) {
			fmt.Println(i,": ",GetMessage())
		})

	}

}
