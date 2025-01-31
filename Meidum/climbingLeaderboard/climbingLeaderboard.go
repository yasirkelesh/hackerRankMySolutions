package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The functin accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */
func removeDuplicates(arr []int32) []int32 {
    uniqueMap := make(map[int]bool)
    uniqueSlice := []int32{}

    for _, num := range arr {
        if !uniqueMap[int(num)] {
            uniqueMap[int(num)] = true
            uniqueSlice = append(uniqueSlice, num)
        }
    }
    return uniqueSlice
}

func binarySearch(arr []int32, target int32) int {
    left, right := 0, len(arr)-1
    for left <= right {
        mid := left + (right-left)/2
        if arr[mid] == target {
            return mid 
        } else if arr[mid] > target {
            left = mid + 1 
        } else {
            right = mid - 1 
        }
    }
    return left 
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    ranked = removeDuplicates(ranked) 
    var hig []int32

    for _, score := range player {
        position := binarySearch(ranked, score)
        hig = append(hig, int32(position+1)) 
    }

    return hig
}
/*func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    ranked = removeDuplicates(ranked)
    var hig []int32
    
    for k := 0; k <= len(player) -1;k++{
        for i := len(ranked)-1; i >= 0;i--{
            if ranked[i] == player[k]{
                hig = append(hig, int32(i+1))
                break
            }else if ranked[i] > player[k]{
                 hig = append(hig, int32(i+2))
                 break
            }else if i == 0 && player[k] >= ranked[i]{
                hig = append(hig, 1)
            }
                
        }
        
    }
    return hig
}*/

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ranked []int32

    for i := 0; i < int(rankedCount); i++ {
        rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
        checkError(err)
        rankedItem := int32(rankedItemTemp)
        ranked = append(ranked, rankedItem)
    }

    playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var player []int32

    for i := 0; i < int(playerCount); i++ {
        playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
        checkError(err)
        playerItem := int32(playerItemTemp)
        player = append(player, playerItem)
    }

    result := climbingLeaderboard(ranked, player)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
