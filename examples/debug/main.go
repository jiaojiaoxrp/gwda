package main

import (
	"fmt"
	"github.com/jiaojiaoxrp/gwda"
	"log"
)

func main() {
	driver, err := gwda.NewUSBDriver(nil)
	if err != nil {
		log.Fatalln(err)
	}

	source, err := driver.Source()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(source)

	fmt.Println(driver.AccessibleSource())

	gwda.SetDebug(true)
}
