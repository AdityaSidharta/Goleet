package src

var romanMaps = map[string]int {
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000,
}

var subMaps = map[string][]string {
    "I": []string{"V", "X"},
    "X": []string{"L", "C"},
    "C": []string{"D", "M"},
}

func isIn(s string, values []string) bool {
    for _, k := range values {
        if s == k {
            return true
        }
    }
    return false
}

func romanToInt(s string) int {
    ns := len(s)
    value := 0
    idx := 0
    for idx < ns {
        curLetter := string(s[idx])
        if isIn(curLetter, []string {"I","X","C"}){
            if idx != ns -1 {
                nexLetter := string(s[idx + 1])
                if isIn(nexLetter, subMaps[curLetter]) {
                    value = value + romanMaps[nexLetter] - romanMaps[curLetter]
                    idx = idx + 2
                    continue
                }
            }
        }
        value = value + romanMaps[curLetter]
        idx = idx + 1
    }
    return value
}