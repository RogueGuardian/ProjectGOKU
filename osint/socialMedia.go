package osint

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func usernameScan(stringList []string) {
	if len(stringList) == 0 {
		return
	}
	//userName := stringList[0]

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Printf("<STUB> HANDLE HTTP ERROR")
	}
	body, bErr := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if bErr != nil {
		fmt.Printf("<STUB> HANDLE SOME ERROR")
	}
	fmt.Printf(string(body))

	// Instagram User Check
	//insta := fmt.Sprintf("https://www.instagram.com/%s/", userName)
	//resp, err := http.Get(insta)
	//if err != nil {
	//	fmt.Printf("<STUB> HANDLE HTTP ERROR")
	//}
	//body, bErr := ioutil.ReadAll(resp.Body)
	//resp.Body.Close()
	//if bErr != nil {
	//	fmt.Printf("<STUB> HANDLE SOME ERROR")
	//}
	//fmt.Printf(string(body))
}
