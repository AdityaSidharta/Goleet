package src

func getlength(line string) int {
    mymap := make(map[string]bool)
    length := 0
    for idx := range line {
        char := string(line[idx])
        _, ok := mymap[char]
        if ok == true {
            return length
        } else {
            mymap[char] = true
            length = length + 1
        }
    }
    return length
}

func lengthOfLongestSubstring(s string) int {
    ns := len(s)
    maxlength := 0
    for idx := 0; idx < ns; idx ++ {
        subs := string(s[idx:])
        length := getlength(subs)
        if length > maxlength{
            maxlength = length
        }
        }
    return maxlength
    }
