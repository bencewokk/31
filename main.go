package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input, inputcons, math int
	var inputlist, inputlist2 []int
	fmt.Scan(&input)
	inputcons = input
	for true {

		inputlist, inputlist2 = nil, nil
		inputstr := strconv.Itoa(input)
		inputlen := len(inputstr)

		for i := 0; i != inputlen; i++ {
			char := string(inputstr[i])
			charint, _ := strconv.Atoi(char)
			switch {
			case i != inputlen-1:
				inputlist = append(inputlist, charint)
			case i == inputlen-1:
				inputlist2 = append(inputlist2, charint)
			}
		}

		szorzo := 1
		equ := 0
		for i := 0; i < len(inputlist); i++ {
			math = inputlist[inputlen-2-i] * szorzo
			szorzo = szorzo * 10
			equ += math
		}
		equ2 := (equ - inputlist2[0]*3)
		fmt.Println(input)
		input = equ2
		if equ2 == 0 {
			fmt.Println(inputcons, "is dividable by 31")
			break
		} else if equ2 < 0 {
			fmt.Println(inputcons, "is not dividableby 31")
			break
		}

	}
}
