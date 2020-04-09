package main

import (
	"fmt"
	"go-101/16-packages/models"
)

// NOTE: VISIBILITY:
// 			- Lower-case naming are not visible outside the package eg. items, addItem
//			- AddItem, Game are visible outside the package

func main() {
	var items map[int]string

	// add items
	models.AddItem("Banana")
	models.AddItem("Yam")
	models.AddItem("Cherry")
	models.AddItem("Grape")

	// remove  an item
	models.RemoveItem(2)
	items = models.GetAll()
	fmt.Println(items)

	// remove invalid item
	models.RemoveItem(5)
	items = models.GetAll()
	fmt.Println(items)

	// get single item
	item := models.Get(1)
	fmt.Println(item)

	// update item
	models.UpdateItem(1, "Mango")
	items = models.GetAll()
	fmt.Println(items)
}
