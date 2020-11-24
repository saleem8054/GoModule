package Factorial
func Factorial(n int) (int,string) {
    var f int = 1
    for i := 2; i <= n; i++ {
        f *= i
    }
    return f,"v3 executed"
}
