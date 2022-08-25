package day2

import (
	"advent_of_code/helpers"
	"fmt"
	"strconv"
	"strings"
)

const path = "day2/input.txt"

func Run() {
    var horizontalPos, depthPos= 0, 0
    lines := helpers.Getlines(path)

    // Part 1
    for i := 0; i < len(lines); i++ {
        line := strings.Split(lines[i], " ")
        direction := line[0]
        value, _ := strconv.Atoi(line[1])
        switch direction {
            case "forward":
                horizontalPos += value
            case "down":
                depthPos += value
            case "up":
                depthPos -= value
        }
    }

    fmt.Printf("(%d, %d) Value: %d\n", horizontalPos, depthPos, horizontalPos*depthPos)

    // Part 2
    var aim = 0
    horizontalPos, depthPos = 0, 0

    for i := 0; i < len(lines); i++ {
        line := strings.Split(lines[i], " ")
        direction := line[0]
        value, _ := strconv.Atoi(line[1])
        switch direction {
            case "forward":
                horizontalPos += value
                depthPos += aim * value
            case "down":
                aim += value
            case "up":
                aim -= value
        }
    }

    fmt.Printf("(%d, %d) Value: %d\n", horizontalPos, depthPos, horizontalPos*depthPos)
}
