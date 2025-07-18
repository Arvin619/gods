package main

import (
	"fmt"

	"github.com/Arvin619/gods/lists/arraylist"
	"github.com/Arvin619/gods/maps/hashmap"
)

// ListSerializationExample demonstrates how to serialize and deserialize lists to and from JSON
func ListSerializationExample() {
	list := arraylist.New[string]()
	list.Add("a", "b", "c")

	// Serialization (marshalling)
	json, err := list.ToJSON()
	if err != nil {
		fmt.Println("failed to list to json:", err)
		return
	}
	fmt.Println(string(json)) // ["a","b","c"]

	// Deserialization (unmarshalling)
	json = []byte(`["a","b"]`)
	err = list.FromJSON(json)
	if err != nil {
		fmt.Println("failed to list from json:", err)
		return
	}
	fmt.Println(list) // ArrayList ["a","b"]
}

// MapSerializationExample demonstrates how to serialize and deserialize maps to and from JSON
func MapSerializationExample() {
	m := hashmap.New[string, string]()
	m.Put("a", "1")
	m.Put("b", "2")
	m.Put("c", "3")

	// Serialization (marshalling)
	json, err := m.ToJSON()
	if err != nil {
		fmt.Println("failed to m to json:", err)
		return
	}
	fmt.Println(string(json)) // {"a":"1","b":"2","c":"3"}

	// Deserialization (unmarshalling)
	json = []byte(`{"a":"1","b":"2"}`)
	err = m.FromJSON(json)
	if err != nil {
		fmt.Println("failed to m from json:", err)
		return
	}
	fmt.Println(m) // HashMap {"a":"1","b":"2"}
}

func main() {
	ListSerializationExample()
	MapSerializationExample()
}
