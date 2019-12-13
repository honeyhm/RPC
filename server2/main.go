package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)


type BankDatabase struct {
	Ssn 			string
	AccountNumber   string
	BadCheck		int
	BadGhest		int
}



type LoanResult struct {
	Result 			string
}

var bankDatabase []BankDatabase

type API int



func (a *API) GetDB(empty string, reply *[]BankDatabase) error {
	*reply = bankDatabase
	return nil
}


func (a *API) AddItem(item BankDatabase, reply *BankDatabase) error {
	bankDatabase = append(bankDatabase, item)
	*reply = item
	return nil
}


//func (a *API) Server2Confirmation(Ssn string, ans2 *LoanResult) error {
func (a *API) Server2Confirmation(Ssn string, ans2 *string) error {
	fmt.Println("Server2Confirmation : ")

	var res LoanResult
	for _, val := range bankDatabase {
		if val.Ssn == Ssn  && val.BadCheck != 0 {
			res.Result = "No"
		}
	}
	*ans2 = res.Result

	return nil

}



func main() {

	fmt.Println("Server2 : ")

	a := BankDatabase{"10", "10",0,0}
	b := BankDatabase{"11", "11",1,0}
	c := BankDatabase{"12", "12",0,1}
	d := BankDatabase{"13", "13",1,1}

	bankDatabase = append(bankDatabase, a)
	bankDatabase = append(bankDatabase, b)
	bankDatabase = append(bankDatabase, c)
	bankDatabase = append(bankDatabase, d)

	fmt.Println("database in server2 : ", bankDatabase)

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API in server2", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4042")
	if err != nil {
		log.Fatal("Listener error in server2", err)
	}
	log.Printf("serving rpc on port %d", 4042)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving server2: ", err)
	}

}
