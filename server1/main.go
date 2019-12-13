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

type BankCustomer struct {
	Ssn 			string
	AccountNumber   string
}

type LoanResult struct {
	Result 			string
}

var bankDatabase []BankDatabase

type API int
var ans1 string
var ans2 string

var s2 *rpc.Client
var s3 *rpc.Client
var err2 error
var err3 error

var t BankCustomer

//func (a *API) FinalConfirmation(customer BankCustomer, reply *LoanResult) error {
//
//	var ans LoanResult
//
//	s2.Call("API.Server2Confirmation", customer, &ans1)
//	fmt.Println("ans1: ", ans1)
//
//	s3.Call("API.Server3Confirmation", customer, &ans2)
//	fmt.Println("ans2: ", ans2)
//
//	if ans1 == "No" || ans2 == "No" {
//		ans.Result = "No"
//	}else{
//		ans.Result ="Yes"
//	}
//	*reply = ans
//	fmt.Println("ans1: ", ans1)
//	fmt.Println("ans2: ", ans2)
//
//	return nil
//}

func FinalConfirmation(customer BankCustomer) LoanResult {
	fmt.Println("FinalConfirmation: ")
	var ans LoanResult

	e1 := s2.Call("API.Server2Confirmation", customer.Ssn, &ans1)
	fmt.Println("&ans1: ", &ans1)
	fmt.Println("e1: ", e1)

	e2 :=s3.Call("API.Server3Confirmation", customer.Ssn, &ans2)
	fmt.Println("&ans2: ", &ans2)
	fmt.Println("e2: ", e2)

	if ans1 == "No" || ans2 == "No" {
		ans.Result = "No"
	}else{
		ans.Result ="Yes"
	}

	return ans
}


func (a *API) GetCustomer(customer BankCustomer , reply *LoanResult) error {
	//*reply = customer
	fmt.Println("GetCustomer: ")
	t =customer
	res := FinalConfirmation(t)
	fmt.Println("GetCustomer2: ")
	fmt.Println("res: ",res)

	*reply = res
	return nil
}


func main() {

	fmt.Println("Server1 : ")

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}
	rpc.HandleHTTP()

	s2, err2 = rpc.DialHTTP("tcp", "localhost:4042")
	if err2 != nil {
		log.Fatal("Connection error in server2 as client : ", err)
	}

	s3, err3 = rpc.DialHTTP("tcp", "localhost:4043")
	if err3 != nil {
		log.Fatal("Connection error in server3 as client : ", err)
	}


	//a := BankDatabase{"10", "10",0,0}

	//b := BankDatabase{"11", "11",1,0}
	//c := BankDatabase{"12", "12",0,1}
	//d := BankDatabase{"13", "13",1,1}
	//
	//
	//s2.Call("API.AddItem", a, &reply)
	//s2.Call("API.AddItem", b, &reply)
	//s2.Call("API.AddItem", c, &reply)
	//s2.Call("API.AddItem", d, &reply)
	//
	//s2.Call("API.GetDB", "", &db)
	//fmt.Println("Database in server2: ", db)
	//
	//
	//s3.Call("API.AddItem", a, &reply)
	//s3.Call("API.AddItem", b, &reply)
	//s3.Call("API.AddItem", c, &reply)
	//s3.Call("API.AddItem", d, &reply)
	//
	//s3.Call("API.GetDB", "", &db)
	//fmt.Println("Database in server3: ", db)
	//

	//s2.Call("API.Server2Confirmation", a, &ans1)
	//fmt.Println("ans1: ", ans1)
	//
	//s3.Call("API.Server3Confirmation", a, &ans2)
	//fmt.Println("ans2: ", ans2)


	//**
	listener0, err0 := net.Listen("tcp", ":4040")
	if err0 != nil {
		log.Fatal("Listener error", err0)
	}
	log.Printf("serving rpc on port %d", 4040)
	http.Serve(listener0, nil)
	if err0 != nil {
		log.Fatal("error serving: ", err0)
	}


	///***********************************************************************************

	//var reply LoanResult
	//var db []BankDatabase


	//s2, err2 := rpc.DialHTTP("tcp", "localhost:4042")
	//if err2 != nil {
	//	log.Fatal("Connection error in server2 as client : ", err)
	//}
	//
	//s3, err3 := rpc.DialHTTP("tcp", "localhost:4043")
	//if err3 != nil {
	//	log.Fatal("Connection error in server3 as client : ", err)
	//}
	//
	//
	//a := BankDatabase{"10", "10",0,0}
	////b := BankDatabase{"11", "11",1,0}
	////c := BankDatabase{"12", "12",0,1}
	////d := BankDatabase{"13", "13",1,1}
	////
	////
	////s2.Call("API.AddItem", a, &reply)
	////s2.Call("API.AddItem", b, &reply)
	////s2.Call("API.AddItem", c, &reply)
	////s2.Call("API.AddItem", d, &reply)
	////
	////s2.Call("API.GetDB", "", &db)
	////fmt.Println("Database in server2: ", db)
	////
	////
	////s3.Call("API.AddItem", a, &reply)
	////s3.Call("API.AddItem", b, &reply)
	////s3.Call("API.AddItem", c, &reply)
	////s3.Call("API.AddItem", d, &reply)
	////
	////s3.Call("API.GetDB", "", &db)
	////fmt.Println("Database in server3: ", db)
	////
	//
	//s2.Call("API.Server2Confirmation", a, &ans1)
	//fmt.Println("ans1: ", ans1)
	//
	//s3.Call("API.Server3Confirmation", a, &ans2)
	//fmt.Println("ans2: ", ans2)


	//s2.Call("API.Server2Confirmation", a, &reply)
	//
	//s3.Call("API.Server3Confirmation", a, &reply)

}


//package main
//
//import (
//	"log"
//	"net"
//	"net/http"
//	"net/rpc"
//)
//
//
//type Item struct {
//	Title string
//	Body  string
//}
//
//type API int
//
//var database []Item
//
//func (a *API) GetDB(empty string, reply *[]Item) error {
//	*reply = database
//	return nil
//}
//
//func (a *API) GetByName(title string, reply *Item) error {
//	var getItem Item
//
//	for _, val := range database {
//		if val.Title == title {
//			getItem = val
//		}
//	}
//
//	*reply = getItem
//
//	return nil
//}
//
//func (a *API) AddItem(item Item, reply *Item) error {
//	database = append(database, item)
//	*reply = item
//	return nil
//}
//
//func (a *API) EditItem(item Item, reply *Item) error {
//	var changed Item
//
//	for idx, val := range database {
//		if val.Title == item.Title {
//			database[idx] = Item{item.Title, item.Body}
//			changed = database[idx]
//		}
//	}
//
//	*reply = changed
//	return nil
//}
//
//func (a *API) DeleteItem(item Item, reply *Item) error {
//	var del Item
//
//	for idx, val := range database {
//		if val.Title == item.Title && val.Body == item.Body {
//			database = append(database[:idx], database[idx+1:]...)
//			del = item
//			break
//		}
//	}
//
//	*reply = del
//	return nil
//}
//
//func main() {
//	api := new(API)
//	err := rpc.Register(api)
//	if err != nil {
//		log.Fatal("error registering API", err)
//	}
//
//	rpc.HandleHTTP()
//
//	listener, err := net.Listen("tcp", ":4040")
//
//	if err != nil {
//		log.Fatal("Listener error", err)
//	}
//	log.Printf("serving rpc on port %d", 4040)
//	http.Serve(listener, nil)
//
//	if err != nil {
//		log.Fatal("error serving: ", err)
//	}
//
//	// fmt.Println("initial database: ", database)
//	// a := Item{"first", "a test item"}
//	// b := Item{"second", "a second item"}
//	// c := Item{"third", "a third item"}
//
//	// AddItem(a)
//	// AddItem(b)
//	// AddItem(c)
//	// fmt.Println("second database: ", database)
//
//	// DeleteItem(b)
//	// fmt.Println("third database: ", database)
//
//	// EditItem("third", Item{"fourth", "a new item"})
//	// fmt.Println("fourth database: ", database)
//
//	// x := GetByName("fourth")
//	// y := GetByName("first")
//	// fmt.Println(x, y)
//
//}