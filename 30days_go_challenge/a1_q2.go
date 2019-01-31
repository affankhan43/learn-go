// MODIFY THIS AREA
// [filename]: a1_q2.go
// [question]: 2
// [instagram username]: affan.ak43
package main

import "fmt"

func Find(s []string, x string) (int,bool) {
	for i, n := range s {
		if x == n {
			return i,true
		}
	}
	return -1,false
}

func main() {
	vowels := []string{"a", "e", "u", "i", "o"}
	scores := []int{12, 8, 1, 4, 3}
	usernames := []string{"westdabestdb", "elonmusk", "vancityreynolds", "porkbundomains", "affan.ak43"}
	score := func (username string) int{
		if len(vowels) == len(scores){
			var score int
			score = 0
			for _,ch := range username{
				index,check := Find(vowels,string(ch))
				if check {
					score+=scores[index]
				}
			}
			return score
		}else{
			return 0
		}
	}
	for _, username := range usernames {
		// write lambda function here and make it equal to score()
		fmt.Printf("%v score: %v \n", username, score(username))
	}
}
