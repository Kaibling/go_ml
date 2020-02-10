package main


import (
	//. "bayesian"
	. "github.com/jbrukh/bayesian"
	"fmt"
) 
	const (
    Good Class = "Good"
    Bad Class = "Bad"
)

func main() {

classifier := NewClassifier(Good, Bad)
goodStuff := []string{"tall", "rich", "handsome"}
badStuff  := []string{"poor", "smelly", "ugly"}
classifier.Learn(goodStuff, Good)
classifier.Learn(badStuff,  Bad)
//scores, likely, _ := classifier.LogScores([]string{"tall", "girl"})
scores, likely, _ := classifier.LogScores([]string{"poor", "girl"})

fmt.Println(scores)
fmt.Println(likely)
probs, likely, _ := classifier.ProbScores([]string{"green", "girl"} )
					 
fmt.Println(probs)
fmt.Println(likely)

}
