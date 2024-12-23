package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

func read() (list_1 []int, list_2 []int) {
    readFile, err := os.Open("input.txt")

    list_1 = make([]int, 0)
    list_2 = make([]int, 0)

    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var val_1 int
    var val_2 int

    for fileScanner.Scan() {
        line := fileScanner.Text()
        fmt.Sscanf(line, "%d %d", &val_1, &val_2)
        list_1 = append(list_1, val_1)
        list_2 = append(list_2, val_2)
    }

    readFile.Close()
    return
}

func first(list_1 []int, list_2 []int) {

    sort.Sort(sort.IntSlice(list_1))
    sort.Sort(sort.IntSlice(list_2))

    var total_distance int
    total_distance = 0

    for i, v1 := range(list_1) {
        var v2 = list_2[i]
        var distance = v1 - v2
        var abs_distance = max(distance, -distance)
        total_distance = total_distance + abs_distance
    }

    fmt.Printf("total distance: %d\n", total_distance)
}

func second(list_1 []int, list_2 []int) {

    similarity := 0
    for _, v1 := range list_1 {
        occurrances := 0
        for _, v2 := range list_2 {
            if v1 == v2 {
                occurrances = occurrances + 1
            }
        }
        similarity = similarity + (occurrances * v1)
    }
    fmt.Printf("similarity: %d\n", similarity)
}

func main() {
    list_1, list_2 := read()
    first(list_1, list_2)

    list_1, list_2 = read()
    second(list_1, list_2)
}
