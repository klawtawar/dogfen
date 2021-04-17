package main

import "fmt"

func main() {
	fmt.Println("Calculate average")

	fmt.Println("Total numbers of scores. ")
	/* Read integer */
	var TOTAL int
    fmt.Scanf("%d", &TOTAL)
    fmt.Println("Total numbers of scores are ", TOTAL)
	
	/* Create array by name score */
    scores := make([]int, TOTAL)
    var sum = 0
    
    for i :=0; i < TOTAL; i++ {
        fmt.Println("score: ", i+1)
        fmt.Scanln(&scores[i])
        
        sum = sum + scores[i]
    }
    var avg = 0
    avg = sum / TOTAL
    fmt.Println(scores)
    fmt.Println("Total sum", sum)
    fmt.Println("average", avg)   
}