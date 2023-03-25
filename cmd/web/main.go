package main

import (
	"fmt"
	"net/http"

	"github.com/riteshapat/go-course/pkg/handlers"
)

const port = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("The service is running in %s", port))
	_ = http.ListenAndServe(port, nil)

	/* http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Jagannnath swami nayana patha gami bhavatume")

		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(fmt.Sprintf("number of bytes written %d", n))
	})
	*/

}
