package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	file ,err := os.Open("sample.txt") 
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()
	res := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		res = scanner.Text()
	}
	count := 0
	for i:= 0 ; i < len(res) ; i++ {
		if res[i] == '(' {
			count++
		} else if res[i] == ')' {
			count--
		}
	}
	fmt.Println(count)

}