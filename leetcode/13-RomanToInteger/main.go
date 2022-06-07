package main

func main() {

}

func romanToIntAnswer(s string) int {
    var roman = map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000}
    if s == "" {
        return 0
    }
    num, lastint, total := 0, 0, 0
    for i
}
