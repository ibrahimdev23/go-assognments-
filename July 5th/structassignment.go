package main

import( 
	"fmt"
	"sort"
)



type Deposit func(int, int) int 

type Bank struct {
  name string
  balance int
  address string
}

//deposit 
func (a *Bank) Deposit(val int) int {
   a.balance = a.balance + val
  return a.balance
}

//withdrawl
func (a *Bank) Withdrawl(val int) int {
   a.balance = a.balance - val
  return a.balance
}




func main() {

	//single bank account example 
  acct_1 := Bank{
    name: "John",
    balance: 0,
    address: "New York, Ny",
  }

  acct_2 := Bank{
    name: "Roman",
    balance: 0,
    address: "Chicago, IL",
  }

//nested Struc with multiple accounts 
  accounts := []Bank{{"Cena", 0,"Minneapolis,MN"}, {"Anthony", 0, "Seattle"}, {"Venus", 0, "England"}}

  //sorting based on name 
sort.SliceStable(accounts, func(i, k int) bool {
  return accounts[i].name < accounts[k].name
})
  

fmt.Println(acct_1)
acct_1.Deposit(100)
fmt.Println(acct_1)
acct_1.Withdrawl(50)
  fmt.Println(acct_2)
  acct_2.Deposit(1000)
  fmt.Println(acct_2)

accounts[1].Deposit(1200)
  accounts[0].Deposit(300)
  accounts[2].Deposit(900)
  accounts[1].Withdrawl(200)
  fmt.Println(accounts)
}