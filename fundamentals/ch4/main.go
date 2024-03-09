package main

import (
	"io"
	"net/http"
	"os"
)

// // CLI APP that takes a string and prints it in all caps with an exclamation point
// func main() {
// 	fmt.Println("What would you like me to scream?")
// 	in := bufio.NewReader(os.Stdin)
// 	s, _ := in.ReadString('\n')
// 	s = strings.TrimSpace(s)
// 	s = strings.ToUpper(s)
// 	fmt.Println(s + "!")
// }

// Web Service
func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
