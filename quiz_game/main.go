package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"	
)


type problem struct {
	q string 
	a string 
}


func main(){


	quiz()
}

func quiz(){
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timelimit   := flag.Int("time",30,"maximum time which can be allocated to each question before the quiz ends, default 30 seconds")
	flag.Parse()

	file, err :=  os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", *csvFilename))
	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil{
		exit("Failed to parse the CSV file")
	}

	problems := parseLines(lines)

	correct := 0
	timer := time.NewTimer(time.Duration(*timelimit)*time.Second)

	// loop:
	for i, p := range problems {
		fmt.Printf("Problem #%d:%s=",i+1,p.q)

		answerch := make(chan string)
		go func(){
			var ans string
			fmt.Scanf("%s\n",&ans)
			answerch <- ans
		}()

		select{
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d. \n",correct,len(problems))
			return
		case ans := <-answerch:
			if ans == p.a {
				correct++
			}
		}
		

	}

	fmt.Printf("You scored %d out of %d. \n",correct,len(problems))

	defer file.Close()

}


func parseLines(lines [][]string ) []problem{
	ret := make([]problem,len(lines))

	for i, line := range lines {

		ret[i] = problem{
		q : line[0],
		a : strings.TrimSpace(line[1]),
		}

	}

	return ret
}

func exit(msg string){
	fmt.Print(msg)
	os.Exit(1)
}
