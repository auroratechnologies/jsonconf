JsonConf [![GoDoc](https://godoc.org/github.com/auroratechnologies/jsonconf?status.svg)](https://godoc.org/github.com/auroratechnologies/jsonconf)
======

######JsonConf

JsonConf implements a method of reading environment variables and defaulting to a configuration json file if the environment variable is not found.


This helps setup a single configuration file that can be used in production, development and local environments, while all having different configurations if the appropriate environment variables are setup.

###Full example
```go
package main

import (
	"fmt"
	"github.com/auroratechnologies/jsonconf"
)


func main() {
	jsonconf.LoadConfig("conf.json")
	redisaddr, err := jsonconf.GetVar("REDIS_ADDR"))
	if (err != nil){
		fmt.Println(err) // No key found
	}
	fmt.Println(redisaddr)
}
```

Full package and usage documentation can be located at [the godocs site](https://godoc.org/github.com/auroratechnologies/jsonconf).
