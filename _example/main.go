package main

import (
	"context"
	"fmt"

	"github.com/libdns/godaddy"
)

func main() {
	p := godaddy.Provider{
		APIToken: "YOUR_APIToken",
	}

	ret, err := p.GetRecords(context.TODO(), "your-domain")

	fmt.Println("Result:", ret)
	fmt.Println("Error:", err)
}
