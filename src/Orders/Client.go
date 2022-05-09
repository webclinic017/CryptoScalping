package Orders

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var Endpoint string = "https://ftx.us/api"

func New(api string, secret string) *FtxClient {

	return &FtxClient{Client: &http.Client{}, Api: api, Secret: []byte(secret)}

}

func (client *FtxClient) sign(signaturePayload string) string {

	mac := hmac.New(sha256.New, client.Secret)
	mac.Write([]byte(signaturePayload))

	return hex.EncodeToString(mac.Sum(nil))

}

func (client *FtxClient) signRequest(method string, path string, body []byte) *http.Request {

	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
	signaturePayload := ts + method + "/api/" + path + string(body)
	signature := client.sign(signaturePayload)
	req, _ := http.NewRequest(method, (Endpoint + path), bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", client.Api)
	req.Header.Set("FTX-SIGN", signature)
	req.Header.Set("FTX-TS", ts)

	return req

}

func (client *FtxClient) _get(path string, body []byte) (*http.Response, error) {

	preparedRequest := client.signRequest("GET", path, body)
	resp, err := client.Client.Do(preparedRequest)

	return resp, err

}

func (client *FtxClient) _post(path string, body []byte) (*http.Response, error) {

	preparedRequest := client.signRequest("POST", path, body)
	resp, err := client.Client.Do(preparedRequest)

	return resp, err

}

func _processResponse(resp *http.Response, result interface{}) error {

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error processing response:", err)
		return err
	}

	err = json.Unmarshal(body, result)

	if err != nil {
		log.Println("Error processing response:", err)
		return err
	}

	return nil
}
