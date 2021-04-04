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
	/** generate code */
	code := gocode.IncCode("RG", "", 15, 13, false)
	fmt.Println("Normal Code =>", code)
	fmt.Println("(RG, , true) =>", gocode.IncCode("RG", "", 2, 13, true))
	fmt.Println("(RG,", code, ", true) =>", gocode.IncCode("RG", code, 2, 13, true))

	fmt.Println("(RG, , false) =>", gocode.IncCode("RG", "", 3, 13, false))
	fmt.Println("(RG,", code, ", false) =>", gocode.IncCode("RG", code, 4, 13, false))

	fmt.Println("( , , false) =>", gocode.IncCode("", "", 5, 13, false))
	fmt.Println("( ,", code, ", false) =>", gocode.IncCode("", code, 6, 13, false))

	/** extract code */
	pref, seq, suff := gocode.ExtractCode(code)
	fmt.Println("Pref:", pref, "Seq:", seq, "Suff:", suff)
}
