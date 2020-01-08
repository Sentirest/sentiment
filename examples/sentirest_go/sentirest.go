package sentirest

import "net/http"
import "bytes"
import "io/ioutil"
import "encoding/json"


type Request struct {
    Text string `json:"text"`
}

var Key string = ""

func GetSentiment(text string) string {

  url := "https://5svpy8pygb.execute-api.us-east-1.amazonaws.com/prod/get/sentiment"

  var request Request
  request.Text = text
  jsonStr, _ := json.Marshal(request)

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("x-api-key", Key)

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  return string(body)
}
