package osint

import (
)

type geoLocation struct {
	// TODO
	// Add all the addressparts shown in the API documentation
	// https://nominatim.org/release-docs/develop/api/Output/#xml

	Emergency string `xml:"addressparts>emergency"`
	Number string `xml:"addressparts>house_number"`
	Name string `xml:addressparts>house_name"`
	Road string `xml:"addressparts>road"`
	Village string `xml:"addressparts>village"`
	Town string `xml:"addressparts>town"`
	City string `xml:"addressparts>city"`
	County string `xml:"addressparts>county"`
	Postal string `xml:"addressparts>postcode"`
	Country string `xml:"addressparts>country"`
}

func (G* geoLocation) toString() string {
	// TODO
	// Add checks to see if items are not empty

	var listing string
	listing = "LOCATION\n========\n"
	listing += "[House Number] " + G.Number + "\n"
	listing += "[Road] " + G.Road + "\n"
	listing += "[Village] " + G.Village + "\n"
	listing += "[Town] " + G.Town + "\n"
	listing += "[City] " + G.City + "\n"
	listing += "[County] " + G.County + "\n"
	listing += "[Postal] " + G.Postal + "\n"
	listing += "[Country] " + G.Country + "\n"
	return listing
}
