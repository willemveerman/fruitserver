package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"math/rand"
	"time"
)

var fruit = []string{"apple",
					 "pear",
					 "banana",
					 "plum",
					 "nectarine",
					 "orange",
					 "kiwi",
					 "apricot",
					 "blueberry",
					 "peach"}

var targetServer = os.Getenv("TARGET_FRUITSERVER")

var targetPort = os.Getenv("TARGET_FRUITPORT")

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	requestFruit := fruit[rand.Intn(10)]

	fmt.Println(requestFruit, " - ", r.Host, " - ", r.RemoteAddr)

	queryString := "http://"+targetServer+".default:"+targetPort+"/"+requestFruit

	fmt.Println(queryString)

	go func() {
		time.Sleep(2 * time.Second)

		resp, err := http.Get(queryString)

		if err != nil {
			fmt.Println("http.Get error = ", err)
		}

		resp.Body.Close()
	}()

	fmt.Fprintf(w, "Yo!, you requested: %s\n and your fruit is %s\n", r.URL.Path, requestFruit)
}

func main() {

	fmt.Println("Starting Server")

	sourcePort := os.Getenv("FRUITPORT")

	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":"+sourcePort, nil))
}
