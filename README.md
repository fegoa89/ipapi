[![Build Status](https://semaphoreci.com/api/v1/fegoa89/ipapi/branches/master/badge.svg)](https://semaphoreci.com/fegoa89/ipapi)[![GoDoc](https://godoc.org/github.com/fegoa89/ipapi?status.svg)](https://godoc.org/github.com/fegoa89/ipapi)

# IPAPI
Golang API wrapper that finds the location of an IP address using ipapi.co.

### FindLocation

Returns the complete location information for an IP address specified in the function parameter. 

```golang
ipapi.FindLocation("178.13.214.11")
```

### ClientLocation

Returns the complete location of the client (device) that’s making the request. You do not need to specify the IP address, it is inferred from the request.

```golang
ipapi.ClientLocation("178.13.214.11")
```

### Usage

```golang
package main

import (
	"fmt"
	"github.com/fegoa89/ipapi"
)

func main() {
	// Location of a specific IP
	ipapi.FindLocation("178.13.214.11")

	// Location of client's IP
	ipapi.ClientLocation()
}
```
