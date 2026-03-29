package main

import (
	"fmt"
	"net/http")

func main(){
	ctype, err := contentType("https://www.linkedin.com/in/alexandre-tibes/")
	if err != nil{
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}


func contentType(url string) (string, error){
	resp, err := http.Get(url)
	if err != nil{
	return "", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("cant find Content-Type header")
	}

	return ctype, nil

}