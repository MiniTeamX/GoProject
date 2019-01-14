package main
import "fmt"
import "time"

func main () {
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	tm := time.Unix(timestamp,0)
	fmt.Println(tm)
	fmt.Println(tm.Format("2006-01-02 15:04:05"))
    fmt.Println(tm.Format("02/01/2006 15:04:05 PM"))
	
}