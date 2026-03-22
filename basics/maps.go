package main

import("fmt")

func main(){
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2330.21,
		"MSFT": 287.98,
	}

	fmt.Println(len(stocks))


	fmt.Println(stocks["MSFT"])

	// zero if not found
	fmt.Println(stocks["TSLA"])



	value, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}


	// set
	stocks["TSLA"] = 832.44
	fmt.Println(stocks)


	delete(stocks, "AMZN")
	fmt.Println(stocks)


	for key:= range stocks{
		fmt.Println(key)
	}


	for key, value := range stocks{
	fmt.Printf("%s -> %.2f \n", key, value)
	}

}