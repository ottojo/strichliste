package main

import (
	"github.com/ottojo/strichlisteUtil"
	"fmt"
	"strconv"
)

func main() {
	api := strichlisteUtil.StrichlisteApi{Url: "https://see-bör.se"}
	u, err := api.GetUsers()
	m := make(map[*strichlisteUtil.User]float64)
	for i := 0; i < len(u); i++ {
		u[i].Transactions, _ = api.GetTransactions(&u[i])
		for _, t := range u[i].Transactions {
			if t.Value < 0 {
				m[&u[i]] += t.Value
			}
		}
		fmt.Println(u[i].Name + ": " + strconv.FormatFloat(m[&u[i]], 'f', 4, 64))
	}

	fmt.Println(err)
	fmt.Println(u)

}
