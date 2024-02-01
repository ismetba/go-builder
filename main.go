package main

import (
	"fmt"

	"github.com/ismetbayandur/endPointBuilder/endpointbuilder"
)

func main() {
	epb := endpointbuilder.SetBaseUrl("https://github.com").AddPath("key1").AddPath("key2").AddParams("param1", "value1").AddParams("param2", "value2").MakeUrl()
	fmt.Println(epb)
}
