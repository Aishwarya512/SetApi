package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type LinkedList struct {
	item int
	next *LinkedList
}
type Set struct {
	buckets [577]*LinkedList
}

// Initializing data structure here
func Constructor() Set {
	return Set{[577]*LinkedList{}}
}

// Function AddItem adds the given item in set.
func (hash *Set) AddItem(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	vars := mux.Vars(r)
	number_str := vars["id"]
	log.Println("User asked to add - ", number_str)
	item, err := strconv.Atoi(number_str)
	if err != nil {
		log.Println("User passed a non number - ", number_str)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "Bad Request", "message" : "Item has to be a number."}`))
		return
	}
	respose := hash.CheckItem(item)
	if respose {
		fmt.Println(respose)
		log.Println("Item already exists - ", number_str)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "Bad Request", "message" : "Item already exists."}`))
		return
	}
	// if item not present, add item as head of LinkedList O(1) time complexity
	head := hash.buckets[item%577]
	if head == nil {
		hash.buckets[item%577] = &LinkedList{item, nil}
	} else {
		hash.buckets[item%577] = &LinkedList{item, head}
	}
	log.Println("Item added - ", number_str)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status" : "OK", "message" : "Item added."}`))
	t2 := time.Now()
	fmt.Printf("Time taken - %v", t2.Sub(t1))
}

// Function RemoveItem removes the given item in set.
func (hash *Set) RemoveItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number_str := vars["id"]
	log.Println("User asked to remove - ", number_str)
	item, err := strconv.Atoi(number_str)
	if err != nil {
		log.Println("User passed a non number - ", number_str)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "Bad Request", "message" : "Item has to be a number."}`))
		return
	}
	head := hash.buckets[item%577]
	// if LinkedList is empty
	if head == nil {
		log.Println("Item to be deleted is not present - ", number_str)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "Bad Request", "message" : "Item to be deleted is not present."}`))
		return
	}
	// if item is present at the beginning of the LinkedList, no need to check the rest
	if head.item == item {
		hash.buckets[item%577] = head.next
		log.Println("Item deleted - ", number_str)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "OK", "message" : "Item deleted."}`))
		return
	}
	// if not, iterate the LinkedList
	var prev *LinkedList
	for head != nil {
		if head.item == item {
			prev.next = head.next
		}
		prev = head
		head = head.next
	}
	log.Println("Item deleted - ", number_str)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status" : "OK", "message" : "Item deleted."}`))
}

// Function HasItem checks the given item in set.
func (hash *Set) HasItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number_str := vars["id"]
	log.Println("User asked to check - ", number_str)
	item, err := strconv.Atoi(number_str)
	if err != nil {
		log.Println("User passed a non number - ", number_str)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "Bad Request", "message" : "Item has to be a number."}`))
		return
	}
	respose := hash.CheckItem(item)
	if respose {
		fmt.Println(respose)
		log.Println("Item present - ", number_str)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status" : "OK", "message" : "Item is present."}`))
		return
	}
	log.Println("Item not present - ", number_str)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status" : "Bad Request", "message" : "Item is not present."}`))
}

// Function CheckItem is called in AddItem and HasItem
func (hash *Set) CheckItem(item int) bool {
	head := hash.buckets[item%577]
	for head != nil {
		//if item is present, return True
		if head.item == item {
			return true
		}
		head = head.next
	}
	return false
}

// Registering URL paths and handlers
func main() {
	MySet := Constructor()
	router := mux.NewRouter()
	router.HandleFunc("/removeItem/{id}", MySet.RemoveItem).Methods("GET")
	router.HandleFunc("/addItem/{id}", MySet.AddItem).Methods("GET")
	router.HandleFunc("/hasItem/{id}", MySet.HasItem).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Panicf("Error starting the server...%v", err)
	}
	log.Println("Server is Up and listening at port 3000...")
}
