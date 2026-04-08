func isPalindrome(s string) bool {

    s = strings.ToLower(s)
    s = strings.ReplaceAll(s, " ", "")
    var input []rune

    for _, ch := range s {
        if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9' ){
            input = append(input, ch)
        }
    }
    for i:= 0 ; i < (len(input) / 2 ); i++ {
        if input[i] != input[len(input) - 1  - i] {
             return false
        }
    }

  
    return true 

}
