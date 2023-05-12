package string

import (
	"strconv"
	"strings"
)

func AddBinary(a string, b string) string {
	var longS, shortS string
	var longI, shortI, mv int
	if len(a) > len(b) {
		longI = len(a) - 1
		longS = a
		shortI = len(b) - 1
		shortS = b
	} else {
		longI = len(b) - 1
		longS = b
		shortI = len(a) - 1
		shortS = a
	}
	result := make([]int, len(longS))
	tmp := make([]string, len(longS))

	for longI >= 0 {
		if shortI >= 0 {
			numL, _ := strconv.Atoi(string(longS[longI]))
			numS, _ := strconv.Atoi(string(shortS[shortI]))
			sum := numL + numS + mv
			mv = caseUpdater(tmp, result, sum, mv, longI)
			shortI--
		} else if mv == 1 {
			numL, _ := strconv.Atoi(string(longS[longI]))
			sum := numL + mv
			mv = caseUpdater(tmp, result, sum, mv, longI)
		} else if longI > 0 {
			return longS[:longI+1] + strings.Join(tmp, "")
		} else {
			tmp[longI] = string(longS[longI])
		}
		longI--
	}

	if mv == 1 {
		return "1" + strings.Join(tmp, "")
	} else {
		return strings.Join(tmp, "")
	}
}

func caseUpdater(tmp []string, result []int, sum int, mv int, i int) int {
	switch sum {
	case 0:
		result[i] = 0
		tmp[i] = "0"
	case 1:
		result[i] = 1
		tmp[i] = "1"
		mv = 0
	case 2:
		result[i] = 0
		tmp[i] = "0"
		mv = 1
	case 3:
		result[i] = 1
		tmp[i] = "1"
	}

	return mv
}
