package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    // Read the file content
    fc, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Get all the lines
    lines := strings.Split(string(fc), "\n")

    // Process each line
    validCount := 0

    for _, line := range lines {
        // Get the minimum count
        dashIdx := strings.Index(line, "-")
        min, _ := strconv.Atoi(line[:dashIdx]) // Assume input data is always valid

        // Get the maximum count
        spaceIdx := strings.Index(line, " ")
        max, _ := strconv.Atoi(line[dashIdx+1 : spaceIdx]) // Assume input data is always valid

        // Get the character
        colonIdx := strings.Index(line, ":")
        chr := line[spaceIdx+1 : colonIdx][0]

        // Get the password
        pass := line[colonIdx+2:]

        // Check the validity (part 1)
        // if isValidPart1(pass, chr, max, min) {
        //     validCount++
        // }

        // Check the validity (part 2)
        if isValidPart2(pass, chr, max, min) {
            validCount++
        }
    }

    fmt.Println(validCount)
}

func isValidPart1(pass string, chr uint8, max int, min int) bool {
    chrCount := 0
    for _, cp := range pass {
        if uint8(cp) == chr {
            chrCount++

            if chrCount > max {
                break
            }
        }
    }

    return chrCount >= min && chrCount <= max
}

func isValidPart2(pass string, chr uint8, max int, min int) bool {
    chrCount := 0

    if pass[min-1] == chr {
        chrCount++
    }

    if pass[max-1] == chr {
        chrCount++
    }

    return chrCount == 1
}
