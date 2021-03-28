package main

import (
	"fmt"

	"github.com/mmuflih/go-trs-code/gocode"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2021-03-28 21:52:43
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func main() {
	code := gocode.IncCode("RG", "", false)
	fmt.Println("Normal Code =>", code)
	fmt.Println("(RG, , true) =>", gocode.IncCode("RG", "", true))
	fmt.Println("(RG,", code, ", true) =>", gocode.IncCode("RG", code, true))

	fmt.Println("(RG, , false) =>", gocode.IncCode("RG", "", false))
	fmt.Println("(RG,", code, ", false) =>", gocode.IncCode("RG", code, false))

	fmt.Println("( , , false) =>", gocode.IncCode("", "", false))
	fmt.Println("( ,", code, ", false) =>", gocode.IncCode("", code, false))
}
