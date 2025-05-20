package subscribe

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/srv-cashpay/merchant/dto"
)

func (r *subscribeRepository) CancelPay(request dto.GetorderID) ([]byte, int, error) {
	baseURL := dto.GetMidtransUrl()
	serverKey := dto.GetMidtransServerKey()

	url := fmt.Sprintf("%s%s/cancel", baseURL, request)
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(serverKey+":"))

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, resp.StatusCode, nil
}
