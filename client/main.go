
package main

import (
	"fmt"
	"log"
	"net/rpc"
)


type BankCustomer struct {
	Ssn 			string
	AccountNumber   string
}

type LoanResult struct {
	Result 			string
}



func main() {

	fmt.Println("Client : ")

	var reply1 LoanResult
	var reply2 LoanResult
	var reply3 LoanResult
	var reply4 LoanResult

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	if err != nil {
		log.Fatal("Connection error in client: ", err)
	}

	x := BankCustomer{"10", "10"}
	y := BankCustomer{"11", "11"}
	z := BankCustomer{"12", "12"}
	w := BankCustomer{"13", "13"}

	client.Call("API.GetCustomer", x, &reply1)
	fmt.Println("Can customer x receive loan :", reply1)
	client.Call("API.GetCustomer", y, &reply2)
	fmt.Println("Can customer y receive loan :", reply2)
	client.Call("API.GetCustomer", z, &reply3)
	fmt.Println("Can customer z receive loan :", reply3)
	client.Call("API.GetCustomer", w, &reply4)
	fmt.Println("Can customer w receive loan :", reply4)

}



