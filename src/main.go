package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// day_one_value := day_one()
	day_two_answer := day_two()
	// fmt.Println(day_one_value)
	fmt.Println(day_two_answer)
}

func read_file() string {
	data, err := ioutil.ReadFile("resources/day_one.txt")
	check(err)

	return string(data)
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
func slice_str_to_int(s []string) []int {
	this_list := make([]int, 0)
	// given a string of strings return a list of ints
	for _, t := range s {
		// convert t into int and add into the small list
		if string(t) == "" {
			continue
		}

		t_int, err := strconv.Atoi(string(t))
		check(err)

		this_list = append(this_list, t_int)

	}
	return this_list
}

func sum(a []int) int {
	this_sum := 0
	for _, i := range a {
		this_sum += i
	}
	return this_sum
}

func day_two() int {
	// read file
	file_data := read_file()
	new_text := strings.Split(file_data, "\n\n")
	// double_split_data := strings.Split(file_data, "\n")
	all_sums := make([]int, 0)

	for _, i := range new_text {

		s := strings.Split(i, "\n")

		a_list := slice_str_to_int(s)
		this_list_sum := sum(a_list)

		all_sums = append(all_sums, this_list_sum)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(all_sums)))

	// get top three and sum
	sum_top_three := all_sums[0] + all_sums[1] + all_sums[2]
	// fmt.Println(all_sums)
	// sort the all sums slice to get the highest first
	return sum_top_three
}
