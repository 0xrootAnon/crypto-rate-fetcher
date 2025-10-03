package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"khwahishIsCooking.com/crypto/data"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*data.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Unexpected status code: %v", res.StatusCode)
	}
	rate := data.Rate{Currency: currency, Price: float64(response.Bid)}
	return &rate, nil
}
