package main

import (
//	"bufio"
	"fmt"
//	"os"
)

const(
	idname = 22
)

var ids = map[int]string {
	idname : "AMF",
}
type ContactInfo struct {
	PhoneNo string
	PinCode string
}

type Person struct {
	firstName string
	LastName  string
	ContactInfo
}

func main() {
	var Alex Person
	Alex = Person{
		firstName: "Alex",
		LastName:  "George",
		ContactInfo: ContactInfo{
			PhoneNo: "4849300221",
			PinCode: "248001",
		},
	}
	Alex.UpdateName("Lokesh")
	name := "bill"
	namePointer := &name
    Alex.Print()
 fmt.Println(&namePointer)
 printPointer(namePointer)
}
 
func printPointer(namePointer *string) {
 fmt.Println(&namePointer)
}

func (p Person) Print() {
	myids := ids[idname]
	fmt.Println(myids)
	fmt.Printf("%+v",p)
}

func (a *Person) UpdateName(newName string){
	// scan := bufio.NewReader(os.Stdin)
    // name, _ := scan.ReadString('\n')
	(a).firstName = newName
	fmt.Println(&a) //<-- This will point to address
	fmt.Println(*a) //<-- This will point to actual value 

}