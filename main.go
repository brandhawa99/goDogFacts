package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type DogFact struct {
	Fact []string `json:"facts"`
}

func main() {
	getFactOrExit()
}

func getFactOrExit() {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("type 'exit' to exit program")
		fmt.Println("Else press enter to continue...")

		text, _ := reader.ReadString('\n')

		//when reading string from STDIN it contains a \n at the end so need to account for that
		if text == "exit\n" {
			fmt.Println("Exiting...")
			fmt.Println("Thank you for trying this cmd app")
			break
		} else {
			fmt.Println("Fact: " + getDogFact() + "\n")
			//give user time to read fact before asking for new fact or exit
			time.Sleep(2 * time.Second)
		}

	}

}

func getDogFact() string {

	res, err := http.Get("http://dog-api.kinduff.com/api/facts")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("api not avail")
	}

	var facts DogFact

	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &facts)

	return facts.Fact[0]
}
