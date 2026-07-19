package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var name string
    fmt.Print("Enter name: ")
    fmt.Scanln(&name) // reads a single whitespace-delimited token

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter full line: ")
    line, _ := reader.ReadString('\n') // reads until the delimiter byte
    fmt.Println("You entered:", line)
}