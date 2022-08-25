package day1

import (
	"advent_of_code/helpers"
	"fmt"
	"strconv"
)

const path = "day1/input.txt"

func Run() {
    lines := helpers.Getlines(path)

    var linesConverted []int
    for _, line := range lines {
        value, _ := strconv.Atoi(line)
        linesConverted = append(linesConverted, value)
    }

    // Part 1
    var counter = 0
    for i := 0; i < len(linesConverted); i++ {
        if i > 0 && linesConverted[i] > linesConverted[i - 1] {
            counter++
            // fmt.Println(linesConverted[i], "(increased)")
        }
    }
    fmt.Println("There are", counter, "measurements that are larger than the previous")

    // Part 2
    counter = 0
    for i := 0; i < len(linesConverted) && i + 3 <= (len(linesConverted)); i++ {
        if i > 0 {
            var firstSlice = linesConverted[i-1 : i+2]
            var secondSlice = linesConverted[i : i+3]

            var firstSum = 0
            var secondSum = 0

            for j:= 0; j < len(firstSlice); j++ {
                firstSum += firstSlice[j]
                secondSum += secondSlice[j]
            }

            if secondSum > firstSum {
                counter++
                // fmt.Println(linesConverted[i], "(increased)")
            }
        }
    }
    fmt.Println("There are", counter, "measurements that are larger than the previous")
}
