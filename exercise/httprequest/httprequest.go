package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"
)

type User struct {
  Name  string
  Email string
}

const (
  ApiKey = "sercet key"
)

func createData(code string) (requestParam url.Values, err error) {
  data := User{
    Name:  "liao",
    Email: "liaozhonghui@talefun.com",
  }
  datajson, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    return nil, fmt.Errorf("json parser failed. %v", err)
  }
  datastr := fmt.Sprintf("%s%s", datajson, ApiKey)

  requestParam = url.Values{}
  requestParam.Set("RequestData", string(datastr))
  return requestParam, nil
}
func Post(apiUrl string, value url.Values) (rs []byte, err error) {
  resp, err := http.PostForm(apiUrl, value)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  return ioutil.ReadAll(resp.Body)
}

func main() {
  const (
    code   = "talefun"
    apiUrl = "http://localhost:7001/ping?format=json"
  )
  requestParam, err := createData(code)
  if err != nil {
    return
  }
  data, err := Post(apiUrl, requestParam)
  if err != nil {
    return
  }

  fmt.Println(data)
  if err != nil {
    os.Exit(1)
  }

  return

}
