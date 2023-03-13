package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	day_one_value := day_one()
	fmt.Println(day_one_value)

}

func day_one() int {
	highest_total := 0
	// read file
	data, err := ioutil.ReadFile("resources/day_one.txt")
	check(err)

	// Convert byte data into string
	text := string(data)

	// split the string by double new line to get a list of smaller strings
	new_text := strings.Split(text, "\n\n")
	// int_slice := make([]int, 1)

	// iterate over the list of smaller strings and convert each seperate string into a number
	// fmt.Println(len(new_text))
	for _, s := range new_text {
		// this_list_int := make([]int, 1)

		s := strings.Split(s, "\n")
		// fmt.Println(this_new_string)
		this_total := 0

		for _, t := range s {
			// convert t into int and add into the small list
			if t == "" {
				continue
			}
			t_int, err := strconv.Atoi(t)
			check(err)

			this_total += t_int

		}

		if this_total > highest_total {
			highest_total = this_total
		}

	}

	// if the sum of the list is greater than the current highest value replace the highest value with the current value
	return highest_total
}
