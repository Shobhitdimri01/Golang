package main

import (
	"fmt"
	"log"
	"sort"

)
	
type mystruct struct{
	FirstName string
	LastName string
}

//calling structure through function

func (m *mystruct)printFirstname()string{
	return m.FirstName
}

type User struct{
	 Age int
	 firstName string
	 LastName string
	 PhoneNo string

}

func main() {
	fmt.Println("hello,World!!")
	var whatToSay string = "Hello Again"
	fmt.Printf("This is %s\n", whatToSay)
	Func_call,num_call := SaySomething()
	fmt.Println("The function return:", Func_call,"\nnumber called is :",num_call)
    myString:="Green"
	log.Println("The color is :",myString)
	changeUsingPointer(&myString)
	log.Println("After calling function Color is:",myString)

	user := User{
		firstName: "Shobhit",
		LastName: "Dimri",
	}
	fmt.Println(user.firstName)
	var guy mystruct
	guy.FirstName = "Nilesh"
	fmt.Println("HI ",guy.FirstName)
	fmt.Println("calling structure through function	--->",guy.printFirstname())

	//map
    myMap:=make(map[string]string)
	myMap["dog"] = "Samson"
	log.Println(myMap["dog"])

	newMap := make(map[string]int)
	newMap["score1"] = 25
	log.Println("Today's score is",newMap["score1"] )

	//using map with your own data type
	YourMap := make(map[string]mystruct)
	 me := mystruct{
		FirstName: "Ramesh",
		LastName: "Tyagi",
	 }
	 YourMap["me"] = me
     log.Println("This is after using your own data type",YourMap["me"].FirstName)

	 //Slice in Go Alternative to Array
	 var myslices []string
	 myslices = append(myslices,"John")
	 myslices = append(myslices, "fam")
	 fmt.Println(myslices)
	 var myintslice []int
	 myintslice = append(myintslice, 5)
	 myintslice = append(myintslice, 3)
	 myintslice = append(myintslice, 1)
	 myintslice = append(myintslice, 10)

	 sort.Ints(myintslice)
	 fmt.Println(myintslice)
     
	 numbers := []int{3,6,9,2,1}
	 fmt.Println(numbers[0:2])
	 //lopping through myslices
	 for _,sliced := range myslices{
		fmt.Println(sliced)
	 }

	 //lopping through map
	 animals := make (map[string]string)
	 animals["dog"] = "Roger"
	 animals["Cat"] = "Fistel"
	 animals["Bird"] = "Twinkle"
 
	 for animalType,animal := range animals{
		fmt.Println(animalType,":",animal)
	 }
	 //lopping through your own structure type
	 var Employee []User
	 Employee = append(Employee, User{21,"Smith","Rog","011-2374-55"})
	 Employee = append(Employee, User{51,"Jack","fog","011-2374-75"})
	 Employee = append(Employee, User{34,"sam","Log","011-2374-59"})
	 Employee = append(Employee, User{25,"Gorge","Mog","011-2374-50"})
	 Employee = append(Employee, User{20,"Kim","Cog","011-2374-21"})

	 for _,Info := range Employee{
		fmt.Println(Info.Age,":",Info.firstName)
	 }

}	
func SaySomething() (string,int) {
	return "Something",8
}
func changeUsingPointer(s *string){
	newValue:="red"
	*s = newValue
}
