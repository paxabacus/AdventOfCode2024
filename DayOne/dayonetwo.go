package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	sim := 0
	list1 := []int{}
	list2 := []int{}
	simL := []int{}
	in := bufio.NewScanner(file)

	for in.Scan() {
		line := in.Text()
		split := strings.SplitN(line, "   ", 2)
		n1, _ := strconv.Atoi(split[0])
		n2, _ := strconv.Atoi(split[1])
		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	for i := 0; i < len(list1); i++ {
		curr := list1[i]
		count := 0
		for _, num := range list2{
			if num == curr{
				count++
			}
		}
		simL = append(simL, curr * count)
	}

	for i := 0; i < len(simL); i++ {
		sim += simL[i]
	}

	fmt.Println(sim)
}

