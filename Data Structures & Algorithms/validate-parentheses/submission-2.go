func isValid(s string) bool {
    var p []rune

    for _, ch := range s {
        if (ch == '}' || ch == ')' || ch == ']' ) && len(p) == 0 {
            return false
        } else if (ch == '}' && p[len(p)-1] == '{' ) || 
        (ch == ')' && p[len(p)-1] == '(' ) ||
         (ch == ']' && p[len(p)-1] == '[') {
            p = p[:len(p)-1]
        } else {
            p = append(p, ch)
        }
    }


    return len(p) == 0 
}
