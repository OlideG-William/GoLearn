package main

// flags Args
import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLoger    *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|
		os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "Warning: ", log.LstdFlags|log.Lshortfile)
	ErrorLoger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)

}

func main() {
	start := time.Now()

	flog := log.New(os.Stdout, "my:", log.LstdFlags)
	flog.SetPrefix("MY TURN: ")
	flog.Println("from flog")
	InfoLogger.Println("ThiS is some info")
	WarningLogger.Println("This is probably important")
	ErrorLoger.Println("Something went wrong")
	fmt.Println("Compile program:", time.Since(start))
}
