package main

import (
  "bytes"
  "fmt"
  "io"
  "mime/multipart"
  "net/http"
  "os"
  "encoding/json"
  "flag"
)

func main() {
  var access_key string
  var filepath   string
  var build_no   string
  var app_id     string

  flag.StringVar(&access_key, "key",    "",      "API key for upload/access")
  flag.StringVar(&filepath,   "file",   "",      "Full path to file")
  flag.StringVar(&build_no,   "build",  "0.0.1", "Build no")
  flag.StringVar(&app_id,     "app_id", "0",     "App id")

  flag.Parse()

	target_url := fmt.Sprintf("http://apps.madebymany.com/admin/apps/%s/versions?auth_token=%s", app_id, access_key)
	filename := filepath
  extraParams := map[string]string{
    "build_no" : build_no,
  }

	postFile(target_url, filename, extraParams)
}

func postFile(target_url string, filename string, params map[string]string){
  // Turn our params map into JSON
  json_bytes, err := json.Marshal(params)
  if err != nil {
    panic(err.Error())
  }
  json_str := string(json_bytes)

  // Create a new buffer for the body
  body_buf := bytes.NewBufferString("")
  body_writer := multipart.NewWriter(body_buf)

  file_writer, err := body_writer.CreateFormFile("file", filename)
  if err != nil {
          panic(err.Error())
  }

  fh, err := os.Open(filename)
  if err != nil {
      panic("Cannot find file")
      panic(err.Error())
  }
  io.Copy(file_writer, fh)

  content_type := body_writer.FormDataContentType()

  body_writer.WriteField("data", json_str)

  body_writer.Close()

  fmt.Println("Uploading IPA..")
  resp, err := http.Post(target_url, content_type, body_buf)
  if err != nil {
    panic(err.Error())
  }
  defer resp.Body.Close()

  if err != nil {
    panic(err.Error())
    fmt.Println(resp.Status)
  }else{
    fmt.Println("Finished uploading..")
    fmt.Println(resp.Status)
  }

}
