package src

func getPalindrome(s string, istart int, iend int) string {
    curString := string(s[istart: iend + 1])
    if (istart == 0) || (iend == len(s) - 1) {
        return curString
    } else if s[istart-1] == s[iend+1]{
        return getPalindrome(s, istart-1, iend+1)
    } else {
        return curString
    }
}

func longestPalindrome(s string) string {
    var finString, propString string
    ns := len(s)
    for idx:= 0; idx < ns; idx++ {
        propString = getPalindrome(s, idx, idx)
        if len(propString) > len(finString) {
            finString = propString
        }
        if idx != ns -1{
            if s[idx] == s[idx + 1]{
                propString = getPalindrome(s, idx, idx + 1)
                if len(propString) > len(finString) {
                    finString = propString
                }
            }
        }
    }
    return finString
}