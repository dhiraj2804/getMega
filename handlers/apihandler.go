package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func ApiHandler() {
	http.HandleFunc("/createtopic", CreateTopic)
	http.HandleFunc("/deletetopic", DeleteTopic)
	http.HandleFunc("/addsubscription", AddSubscription)
	http.HandleFunc("/deletesubscription", DeleteSubscription)
	http.HandleFunc("/publish", Publish)
	http.HandleFunc("/subscribe", Subscribe)
	http.HandleFunc("/unsubscribe", UnSubscribe)
	http.HandleFunc("/ack", Ack)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func CreateTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func DeleteTopic(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func AddSubscription(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func Publish(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func UnSubscribe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func Ack(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
