package main

import (
	"fmt"

	"github.com/moul/anonuuid"
)

func main() {
	fmt.Println(anonuuid.New().FakeUUID("15573749-c89d-41dd-a655-16e79bed52e0"))
}
