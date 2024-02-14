package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// CLI example
	level := flag.String("level", "CRITICAL", "log Level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
			// log.Println(line)
		}
	}
}

// func main() {
// 	// 	fmt.Println("Hello, World!")
// 	// Web Server example
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, "Hello, Web Services!")
// 	})
// 	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
// 		http.ServeFile(w, r, "./home.html")
// 	})
// 	http.ListenAndServe("localhost:3000", nil)
// 	// http.ListenAndServe("127.0.0.1:3000", nil)
// }
