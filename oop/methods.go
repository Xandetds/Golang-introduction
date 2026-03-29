package main

import(
		"fmt"
		"time"
	)


type Budget struct{
	CampaignID string
	Balance float64
	Expires time.Time
}

func (b Budget) TimeLeft() time.Duration{
	return b.Expires.Sub(time.Now().UTC())
}


func (b*Budget) Update(sum float64){
	b.Balance += sum
}




func main(){
	b1 := Budget{"Kittens", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1.TimeLeft())


	b1.Update(10.5)
	fmt.Println(b1.Balance)


}