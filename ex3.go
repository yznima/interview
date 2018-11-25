// You are given a license key represented as a string S which consists 
// only alphanumeric character and dashes. The string is separated into
//  N+1 groups by N dashes.

// Given a number K, we would want to reformat the strings such that each
// group contains exactly K characters, except for the first group which
// could be shorter than K, but still must contain at least one character.
// Furthermore, there must be a dash inserted between two groups and all
// lowercase letters should be converted to uppercase.

// Given a non-empty string S and a number K, format the string according
// to the rules described above.

func licenseKeyFormatting(S string, K int) string {
    dash := 0
    for _, s := range S {
        if s == '-' {
            dash++
        }
    }
    
    oldSize := len(S) - dash
    if oldSize == 0 {
        return ""
    }
    
    var newSize int
    if oldSize % K == 0 {
        newSize = oldSize + oldSize/K - 1
    } else {
        newSize = oldSize + oldSize/K
    }
    
    bytes := make([]byte, 0, newSize)
    for _, s := range S {
        if s == '-' {
            continue
        }
        
        bytes = append(bytes, byte(unicode.ToUpper(s)))
        if (len(bytes) != cap(bytes) && (cap(bytes) - len(bytes)) % (K+1) == 0) {
            bytes = append(bytes, '-')
        }
    }
    
    return string(bytes)
}