package gocode

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2021-03-28 21:42:51
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func IncCode(pref, code string, seq, codeLength int, force bool) string {
	/** force with old code */
	var r rune
	var codes []string
	if force {
		if code != "" {
			return code
		}

		if code == "" {
			goto _new
		}
	}

	/** else, generate new code or incr.. suffix */
	if code != "" {
		codes = strings.Split(code, ".")
	}
	if pref == "" && len(codes) > 0 {
		pref = codes[0]
	}

	if len(codes) < 3 {
		goto _new
	}

	r, _ = utf8.DecodeRuneInString(codes[2:][0])
	return fmt.Sprintf("%s.%s.%s", codes[0], codes[1], string(r+1))

_new:
	now := time.Now()
	y, w := now.Local().ISOWeek()
	fmt.Println(y, w)
	code = now.Format("0601")
	suff := "A"
	seqStr := ""
	for i := 1; i <= codeLength-(len(pref+code)+4); i++ {
		seqStr += "0"
	}
	return fmt.Sprintf("%s%s.%s.%s", pref, code, seqStr+strconv.Itoa(seq), suff)
}

// ExtractCode returns prefix, seq, suffix
// Prefix is the begining of code
// Seq is sequence number of code
// Suffix is sequence number place on end of code
// suffix consists of the uppercase alphabet ex: A, B, C ...
func ExtractCode(code string) (string, int, string) {
	codes := strings.Split(code, ".")
	if len(codes) < 3 {
		return "", 1, "A"
	}
	seq, err := strconv.Atoi(codes[1])
	if err != nil {
		seq = 1
	}

	return codes[0], seq, codes[2]
}
