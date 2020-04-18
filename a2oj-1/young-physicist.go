package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleanString(stream string, seperator string) []int {
	trimmedStream := strings.TrimSpace(stream)
	splitArr := strings.Split(trimmedStream, seperator)
	// convert strings to integers and store them in a slice
	cleanArr := make([]int, len(splitArr))
	for index, val := range splitArr {
		cleanArr[index], _ = strconv.Atoi(val)
	}
	return cleanArr
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, _ := reader.ReadString('\n')
	intN, _ := strconv.Atoi(strings.TrimSpace(n))

	result := new([3]int)
	for intN > 0 {
		inputLine, _ := reader.ReadString('\n')
		point := cleanString(inputLine, " ")
		result[0] += point[0]
		result[1] += point[1]
		result[2] += point[2]
		intN--
	}
	equi := true
	for i := 0; i < 3; i++ {
		if result[i] != 0 {
			equi = false
		}
	}
	if equi {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
