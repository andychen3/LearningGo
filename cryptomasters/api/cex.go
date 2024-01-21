package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/andychen3/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required; %d received", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	response, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}

	var cryptoRate CEXresponse
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &cryptoRate)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("Status code received: %v", response.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: cryptoRate.Bid}
	return &rate, nil
}
