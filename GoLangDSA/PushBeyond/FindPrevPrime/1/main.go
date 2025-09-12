package piscine


func FindPrevPrime(nb int) int {
    for i := nb; i >= 2; i-- {
        if IsPrime(i) {
            return i
        }
    }
    return 0
}


func IsPrime(num int) bool {
    i := 2
    for i*i <= num {
        if num % i == 0 {
            return false
        }
        i++
    }
    return true
}