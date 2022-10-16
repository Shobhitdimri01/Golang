package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	var links = []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://flipkart.com",
	}

	c := make(chan string)
	for _, link := range links {
		go loopOverLink(link,c)
	}
	//Receiving value from channel
	//l has link that we have passed to a channel
	for l:= range c {
		go func (link string)  {
			time.Sleep(5*time.Second)
			loopOverLink(link,c)	
		}(l)
	}
}

func loopOverLink(link string, c chan string) {
	_, err := http.Get(link)
	
	if err!= nil{
		fmt.Println(link,"is down")
		//Sending data from channel
		c <- link
		return
	}

	fmt.Println(link,"is up !")
	c <- link
}