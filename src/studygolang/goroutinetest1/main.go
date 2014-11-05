package main 

import (     
	"fmt"
	"strconv" 
) 

func main() {     
	ch := make(chan int)     
	task("A", ch)     
	task("B", ch)     
	fmt.Printf("begin\n")     
	<-ch     
	<-ch 
	
} 

func task(name string, ch chan int) {     
	go func() {         
		i:= 1         
		for {             
			//fmt.Printf("%s %s\n", name, strconv.Itoa(i))             
			print(name + " " + strconv.Itoa(i) + "\n")
			i++         
		}         
		ch <- 1     
	}(); 
}


