package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.SaveFile("my_card")
	fmt.Println("File contents are : ")
	SavedContent := cards.ReadFile("my_card")
	deal(cards,10)
	SavedContent.print()
	fmt.Println("*************+++++++++++**************")
	cards.Shuffle()
	fmt.Println("Cards are shuffled now!!")
	cards.print()
	/* odd /even
	var even []int 
	var odd []int 
	myslice := []int{0,1,2,3,4,5,6,7,8,9,10}
	for _,num := range myslice{
		if num%2 == 0{
			even = append(even,num)
		}else{
				odd = append(odd,num)
		}
	}
	fmt.Println("Even num are : ",even)
	fmt.Println("odd num are : ",odd)
*/
}
