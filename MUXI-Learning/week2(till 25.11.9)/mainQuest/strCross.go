package main

import (
	"fmt"
)

type Empty interface{}

func main() {
	strLetters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	
	sigLetters := make(chan Empty)
	sigNumbers := make(chan Empty) 


	done := make(chan Empty, 2)

	var stop Empty


	go func() {
		for i := 0; i <= 24; i += 2 {
			<-sigLetters      
			fmt.Print(strLetters[i:i+2])      
			sigNumbers <- stop
		}
		done <- stop      
	}()


	go func() {
		for i := 0; i <= 24; i += 2 {
			<-sigNumbers  
			fmt.Printf("%d%d",i, i+1) 
			if i < 24 {sigLetters <- stop}
		}
		close(sigNumbers) 
		close(sigLetters) 
		done <- stop      
	}()
	
	sigLetters <- stop

	<- done
	<- done
}
		