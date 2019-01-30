package main

// // Prize : ...
// type Prize struct {
// 	Category string `json:"category"`
// }

// // Winner : ...
// type Winner struct {
// 	ID            string  `json:"id"`
// 	Firstname     string  `json:"firstname"`
// 	Lastname      string  `json:"surname"`
// 	PrizeCategory []Prize `json:"prizes"`
// }

// // Laureates : ...
// type Laureates struct {
// 	Winners []Winner `json:"Laureates"`
// }

// func main() {
// 	resp, _ := http.Get("http://api.nobelprize.org/v1/laureate.json?bornCountry=poland&gender=female")
// 	respData, _ := ioutil.ReadAll(resp.Body)
// 	var mx Laureates
// 	json.Unmarshal(respData, &mx)
// 	fmt.Println(mx)
// }
