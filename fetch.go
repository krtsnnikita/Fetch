// Fetch выводит ответ на запрос по заданному URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
        for _, url := range os.Args[1:] {
				prefix := strings.Contains(url, "http://")
				fmt.Println(prefix)
				if !prefix {
				url = "http://" + url
			}
                resp, err := http.Get(url)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
                        os.Exit(1)
                }

                b, err := io.Copy(os.Stdout, resp.Body)
                resp.Body.Close()
                if err != nil {
                        fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
                        os.Exit(1)
                }
                //fmt.Println(strings.Contains(url, "http://"))
                fmt.Println(b)
        }

}