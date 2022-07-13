package main
import (
	"fmt"
	"sync"
)

//Create a program that declares two anonymous functions.


func main(){

	var wg sync.WaitGroup
	wg.Add(2)

	//One function counts down from 100 to 0
	go func() {
		for i:=100; i >=0;i-- {
			fmt.Println("Counting Down:",i)
		}
		wg.Done()
	}()




	//One function counts up from 0 to 100.
		go func(){
			for i:=0; i <100; i++{
				fmt.Println("Counting up:",i)
			}
			wg.Done()
		}()


		fmt.Println("Starting")
		wg.Wait()
		fmt.Println("Finished")

}






// var (
// 	printNow chan bool
//     i int
// )

// func main() {
// 	printNow = make(chan bool)
   
    
// 	go printer()
// 	go sender()
//    	for {
// 	}
// }


// func printer() {
// 	for {
// 		if _, ok :=<-printNow ;ok{
// 			fmt.Println("Recieved !" , i) 
// 		}//else{
//              //return
//           //   os.Exit(0)
//        // }
        
// 	}
// }
// func sender() {
// 	for {
// 		for i = 0; i < 10; i++ {
//             fmt.Println("Call" , i)
            
// 			printNow <- true
            
// 			time.Sleep(1 * time.Millisecond)
// 		}
//        // close(printNow)
//         //return
// 	}
// }