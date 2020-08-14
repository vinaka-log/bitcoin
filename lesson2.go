package main

import"fmt"
import"time"
import"os/user"


func main() {
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}
