# apod

a simple library for getting Astronomy Picture Of the Day information from [NASA's API](https://api.nasa.gov/api.html#apod).

## usage

```go
package main

import "github.com/stewart/apod"

var client = apod.New("")

func main() {
	// get the latest APOD
	picture, err := client.Latest()

	// get the APOD for New Years, 1999
	date, _ := apod.NewDate("1999-01-01")
	picture, err := client.ForDate(date)
}
```
