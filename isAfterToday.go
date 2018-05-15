package main

import (
	"fmt"
    "time"
)


const shortForm = "2006-01-02"
const longForm = "2006-01-02T15:04:05"

func main() {
    fmt.Println(isAfterToday("2018-05-14")) // returns false
    fmt.Println(isAfterToday("2018-05-15")) // returns false
    fmt.Println(isAfterToday("2018-05-16")) // returns true

    // Extra Credit
    fmt.Println(isAfterToday("2018-05-14T11:00:33")) // returns false
    fmt.Println(isAfterToday("2018-05-15T23:59:59")) // returns false
    fmt.Println(isAfterToday("2018-05-16T00:00:01")) // returns true
}

func isAfterToday(s string) bool {
    return time.Now().Format(time.RFC3339)[:10] < s[:10]
}

// func isAfterToday(date string) bool {
//     var parsed time.Time
//     var err error
//
//     now := time.Now()
//
//     switch len(date) {
//     case len(longForm):
//         parsed, err = time.Parse(longForm, date)
//     default:
//         parsed, err = time.Parse(shortForm, date)
//     }
//
//     if err != nil {
//         return false
//     }
//
//     if now.Year() > parsed.Year() {
//         return false
//     }
//
//     if now.Month() > parsed.Month() {
//         return false
//     }
//
//     if now.Day() >= parsed.Day() {
//         return false
//     }
//
//     return true
// }
