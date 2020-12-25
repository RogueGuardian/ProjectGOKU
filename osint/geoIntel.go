package osint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getGPSLocation(NS string, EW string) string {
	// Parse input coords to float and error check
	ns, nsErr := strconv.ParseFloat(NS, 64)
	ew, ewErr := strconv.ParseFloat(EW, 64)
	if nsErr != nil || ewErr != nil {
		return "Invalid Lat or Lon Coordinate"
	}
	if ns < -90 || ns > 90 {
		return "Invalid Lat Coordinate"
	}
	if ew < -180 || ew > 180 {
		return "Invalid Lon Coordinate"
	}

	// Create format string and do a Get request and recieve data in JSON
	getString := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%s&lon=%s&format=json", NS, EW)
	resp, err := http.Get(getString)
	if err != nil {
		fmt.Printf("<STB> HTTP ERROR")
		return "NA"
	}
	loc, lErr := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if lErr != nil {
		fmt.Printf("<STB> READ ERROR")
		return "NA"
	}

	// Parse JSON into relevant data
	parseLoc := geoLocation{}
	jsonErr := json.Unmarshal([]byte(loc), &parseLoc)
	if jsonErr != nil {
		fmt.Printf("<STB> JSON ERROR")
		return "XD"
	}
	return parseLoc.toString()
}

func location(stringList []string) {
	if len(stringList) == 0 {
		return
	}

	switch strings.ToLower(stringList[0]) {
	case "gps":
		if len(stringList) < 3 {
			fmt.Printf("<STB> TODO ERROR")
		}
		fmt.Printf("%s", getGPSLocation(stringList[1], stringList[2]))
		break
	}
}
