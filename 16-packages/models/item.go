package models

import (
	"go-101/16-packages/data"
)

// AddItem Add a string to the store
func AddItem(item string) int {
	key := len(data.Store)
	data.Store[key+1] = item

	return key
}

// RemoveItem remove item from the map by key
func RemoveItem(key int) bool {
	value := data.Store[key]
	if value != "" {
		delete(data.Store, key)
		return true
	}
	return false

}

// UpdateItem Update item at passed in key
func UpdateItem(key int, item string) string {
	data.Store[key] = item

	return data.Store[key]
}

// GetAll returns the map collection
func GetAll() map[int]string {
	return data.Store
}

// Get returns a string at the key
func Get(key int) string {
	return data.Store[key]
}
