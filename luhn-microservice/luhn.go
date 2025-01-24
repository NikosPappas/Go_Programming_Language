package main

import "strconv"

// LuhnCheck checks if a given number is valid according to the Luhn algorithm.
func LuhnCheck(number string) bool {
    sum := 0
    alternate := false
    for i := len(number) - 1; i >= 0; i-- {
        digit, err := strconv.Atoi(string(number[i]))
        if err != nil {
            return false // Invalid character in number
        }

        if alternate {
            digit *= 2
            if digit > 9 {
                digit -= 9
            }
        }
        sum += digit
        alternate = !alternate
    }
    return sum%10 == 0
}
