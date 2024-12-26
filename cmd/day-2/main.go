package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func read() ([][]int) {
    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }

    var reports = make([][]int, 0)

	fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
        line := fileScanner.Text()
		values := strings.Split(line, " ")

		report := make([]int, 0)
		for _, value := range values {
			i, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			report = append(report, i)
		}
		reports = append(reports, report)
    }

    readFile.Close()
    return reports
}

func abs(x, y int) int {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func is_safe(report []int) bool {
	var previous int
    var has_increased = false
	var has_decreased = false

    for i, current := range report {
		if i == 0 {
			previous = current
			continue
		}

		if current > previous {
			has_increased = true
		}

		if current < previous {
			has_decreased = true
		}

		distance := abs(current, previous)
		if distance < 1 || distance > 3 {
			return false
		}

		if has_decreased && has_increased {
			return false
		}
		previous = current
	}

	return true
}


func is_safe_removed(report []int) bool {

	if is_safe(report) {
		return true
	}

	// really need to learn how to manupulate slices properly
	for i_to_remove, _ := range report {
		report_removed := make([]int, 0)
		for i, val := range report {
			if i != i_to_remove {
				report_removed = append(report_removed, val)
			}
		}

		if is_safe(report_removed) {
			return true
		}
	}
	return false
}

func analyse(removed bool) {
	
	reports := read()

	safe_reports := 0
    for _, report := range reports {

		if !removed {
			if is_safe(report) {
				safe_reports += 1
			}
		} else {
			if is_safe_removed(report) {
				safe_reports += 1
			}
		}
	}
	fmt.Printf("Safe reports: %d\n", safe_reports)
}


func main() {
	analyse(false)
	analyse(true)
}
