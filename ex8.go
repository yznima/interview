// https://leetcode.com/problems/super-palindromes

// Let's say a positive integer is a superpalindrome if it is a 
// palindrome, and it is also the square of a palindrome.

// Now, given two positive integers L and R (represented as strings),
// return the number of superpalindromes in the inclusive range [L, R].

func superpalindromesInRange(L string, R string) int {
    l, err := strconv.ParseInt(L, 10, 64)
    if err != nil {
    	panic(err)
    }
    r, err := strconv.ParseInt(R, 10, 64)
    if err != nil {
    	panic(err)
    }
    
    count := 0
    for i := int64(0); i <= 1e5; i++ {
        if i > r {
            break
        }
        j, k := toPal(i)
        if isVal(l, r, j) {
        	count++
        }
        if isVal(l, r, k) {
        	count++
        }
    }
    return count
}

func isVal(l, r, i int64) bool {
    if isPal(i) {
        i2 := i * i
        return isPal(i2) && i2 >= l && i2 <= r
    }
    return false
}

func toPal(i int64) (int64, int64) {
    r, n, j, k := int64(0), i, i, i
    if n > 0 {
        r, n = n % 10, n / 10
        j = 10 * j + r
    }
    
    for n > 0 {
        r, n = n % 10, n / 10
        j, k = 10 * j + r, 10 * k + r
    }

    return j, k
}

func isPal(i int64) bool {
    r, n := int64(0), i
    for n > 0 {
        r, n = 10 * r + n % 10, n / 10
    }
    return r == i
}

/** TO SLOW
func superpalindromesInRange(L string, R string) int {
    l, err := strconv.ParseInt(L, 10, 64)
    if err != nil { panic(err) }
    r, err := strconv.ParseInt(R, 10, 64)
    if err != nil { panic(err) }

    ls := int64(math.Sqrt(float64(l)))
    rs := int64(math.Sqrt(float64(r)))
    count := 0
    for ls <= rs {
        if isPal(ls) && isPal(ls * ls) {
            fmt.Println(ls)
            count++
        }
        ls++
    }
    return count
}

func isPal(i int64) bool {
    r, n := int64(0), i
    for n > 0 {
        r, n = 10 * r + n % 10, n / 10
    }
    return r == i
}
 */
