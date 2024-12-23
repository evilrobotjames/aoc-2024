package main

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

func main() {

    readFile, err := os.Open("input.txt")

    list_1 := make([]int, 0)
    list_2 := make([]int, 0)

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

    fmt.Println(list_1)
    fmt.Println(list_2)

    sort.Sort(sort.IntSlice(list_1))
    sort.Sort(sort.IntSlice(list_2))
    fmt.Println(list_1)
    fmt.Println(list_2)

    var total_distance int
    total_distance = 0

    for i, v1 := range(list_1) {
        var v2 = list_2[i]
        var distance = v1 - v2
        var abs_distance = max(distance, -distance)
        fmt.Printf("%d %d = %d\n", v1, v2, abs_distance)
        total_distance = total_distance + abs_distance
    }

    fmt.Printf("total distance: %d\n", total_distance)
}
