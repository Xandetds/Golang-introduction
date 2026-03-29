package main

import(
		"fmt"
		"time"
	)


	// uppercase like CampaignID is public. not uppercase like balance is private 
type Budget struct{
	CampaignID string
	Balance float64
	Expires time.Time
}


func main(){
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)

	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.CampaignID)


	b2 := Budget{
		Balance : 19.3,
		CampaignID: "puppies",
	}
	fmt.Printf("%#v\n", b2)


	b3 := Budget{}
	fmt.Printf("%#v\n", b3)

}