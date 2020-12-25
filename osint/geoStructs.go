package osint

import (
	"fmt"
	"strings"

	"../helpers"
)

type geoLocation struct {
	Address map[string]interface{} `json:address`
}

func (G* geoLocation) toString() string {
	var listing string

	// Basic Location Data
	listing = helpers.PrintColor("bold", "green", "none", "LOCATION\n========\n")
	for k,v := range G.Address {
		listing += fmt.Sprintf("[%s] %v\n", helpers.PrintColor("bold", "green", "none", strings.Title(k)), v)
	}

	return listing
}
