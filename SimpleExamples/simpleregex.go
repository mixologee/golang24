package main

import (
	"fmt"
	"log"
	"regexp"
	"io/ioutil"
	"net/http"
)

func main() {
	// variable that holds what you are looking for
	// one fails, one passes since regex is case sensitive
	needle1 := "chocolate"
	needle2 := "(?i)chocolate"
	// variable that has what you are looking through
	haystack := "Chocolate is my favorite!"
	// check for a match
	match1, err := regexp.MatchString(needle1, haystack)
	// error handle
	if err != nil {
		log.Fatal(err)
	}
	// print result
	fmt.Println(match1)
	match2, err := regexp.MatchString(needle2, haystack)
	// error handle
	if err != nil {
		log.Fatal(err)
	}
	// print result
	fmt.Println(match2)
	// samples of using regex- MustCompile panics if it can't compile the regex expression you use
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdf£33£3"))
	fmt.Println(re.MatchString("roger"))
	fmt.Println(re.MatchString("iamthebestuserofthisappevaaaar"))

	// data transformation using regex
	usernames := [4]string{
		"slimshady99",
		"!asdf£33£3",
		"roger",
		"Iamthebestuserofthisappevaaaar",
	}
	// variables for regex matching 
	re1 := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
	an := regexp.MustCompile("[[:^alnum:]]")
	// loop through usernames
	for _, username := range usernames {
		if len(username) > 12 {
			username = username[:12]
			fmt.Printf("trimmed username to %v\n", username)
		}
		if !re1.MatchString(username) {
			username = an.ReplaceAllString(username, "x")
			fmt.Printf("rewrote username to %v\n", username)
		}
	}
	// parsing data with regex
	resp, err := http.Get("https://petition.parliament.uk/petitions")
	// error handle
	if err != nil {
		log.Fatal(err)
	}
	// close the GET request when we are done
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// error handle
	if err != nil {
		log.Fatal(err)
	}

	src := string(body)

	re2 := regexp.MustCompile("\\<h3\\>.*\\</h3\\>")
	titles := re2.FindAllString(src, -1)

	for _, title := range titles {
		fmt.Println(title)
	}
	
	// cleaning up HTML code from a GET response (screen scraping)
	resp1, err := http.Get("https://petition.parliament.uk/petitions")
	if err != nil {
		log.Fatal(err)
	}
	defer resp1.Body.Close()
	body1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		log.Fatal(err)
	}

	src1 := string(body1)

	re3 := regexp.MustCompile("\\<h3\\>.*\\</h3\\>")
	rHTML := regexp.MustCompile("<[^>]*>")
	titles1 := re3.FindAllString(src1, -1)

	for _, title := range titles1 {
		cleanTitle := rHTML.ReplaceAllString(title, "")
		fmt.Println(cleanTitle)

	}
}