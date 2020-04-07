package main

import "fmt"

func main() {
	contacts := make(map[string]string)
	contacts["monday"] = "2nd day Street"
	contacts["tuesday"] = "3rd day Street"
	contacts["wednesday"] = "4th day Street"
	contacts["thursday"] = "5th day Street"
	contacts["friday"] = "6th day Street"
	contacts["saturday"] = "7th day Street"

	fmt.Println(len(contacts))
	fmt.Println(contacts)

	// deleting a property in a map
	contacts["whatday"] = "?th day Street"
	fmt.Println(contacts)
	delete(contacts, "whatday")
	fmt.Println(contacts)

}
