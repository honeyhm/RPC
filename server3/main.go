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


//func (a *API) Server3Confirmation(Ssn string, ans3 *LoanResult) error {
func (a *API) Server3Confirmation(Ssn string, ans3 *string) error {

	var res LoanResult
	for _, val := range bankDatabase {
		if val.Ssn == Ssn  && val.BadGhest != 0 {
			res.Result = "No"
		}
	}
	*ans3 = res.Result
	return nil

}


func main() {

	fmt.Println("Server3 : ")

	a := BankDatabase{"10", "10",0,0}
	b := BankDatabase{"11", "11",1,0}
	c := BankDatabase{"12", "12",0,1}
	d := BankDatabase{"13", "13",1,1}

	bankDatabase = append(bankDatabase, a)
	bankDatabase = append(bankDatabase, b)
	bankDatabase = append(bankDatabase, c)
	bankDatabase = append(bankDatabase, d)

	fmt.Println("database in server3 : ", bankDatabase)

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API in server3", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4043")
	if err != nil {
		log.Fatal("Listener error in server3", err)
	}
	log.Printf("serving rpc on port %d", 4043)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving server3: ", err)
	}


}
