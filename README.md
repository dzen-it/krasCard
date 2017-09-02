# krasCard
API for transport cards of Krasnoyarsk city.

Example. Get ballance of card:

```go
package main

import (
	"fmt"
	"krasCard"
)

func main() {
	info, err := krasCard.GetTransportCard("1122334455667788")
	if err != nil {
		panic(err)
	}
	fmt.Println(info.CurrentBalance)
}

```