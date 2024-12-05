package main

import (
	"bufio" //IO
	"fmt" //Print
	"log" //Error Handling
	"os" //Reading
	"strings" //STRINGS
	"strconv" //Conversion
	"sort" //Sorting
	"math" //Mathematics
)

func main() {
	file, err := os.Open("input.txt") //Open the file
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum := 0
	list1 := []int{} //Slices
	list2 := []int{}
	scanner := bufio.NewScanner(file) //Reading

	for scanner.Scan() {
		currLine := scanner.Text() 
		splitted := strings.SplitN(currLine, "   ",2)
		fmt.Println(splitted)
		typ, _ := strconv.Atoi(splitted[0])
		typ2, _ := strconv.Atoi(splitted[1])
		
		
		list1 = append(list1, typ)
		list2 = append(list2, typ2)
		
	}
	fmt.Println("finished reading")
//	sort.Slice(list, func(i, j int) bool) {
//		return a[i] < a[j]
//	})
	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))
	fmt.Println(list1)

	fmt.Println(list2)
	diff := []int{}
	for i := 0; i < len(list1); i++ {
		subtract := list2[i] - list1[i]
		abVal := math.Abs(float64(subtract))
		diff = append(diff, int(abVal))
	}
	
	for i := 0; i < len(diff); i++ {
		sum += diff[i]
	}
	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
