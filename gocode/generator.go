package gocode

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/mmuflih/datelib"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2021-03-28 21:42:51
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

func IncCode(pref, code string, force bool) string {
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
	code = time.Now().Format(datelib.YMD_SHORT_WS)
	suff := "A"
	return fmt.Sprintf("%s.%s.%s", pref, code, suff)
}
