package main

import(
	"fmt"
	"strings"
)

func main(){
	
	text:= "i am learning go i like learning a lot"
	words := strings.Fields(text)


	counts := map[string]int{}
	for _, word:= range words{
		counts[strings.ToLower(word)] ++
	}

	fmt.Println(counts)
}