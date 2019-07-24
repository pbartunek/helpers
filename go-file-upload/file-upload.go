//
// based on https://tutorialedge.net/golang/go-file-upload-tutorial/
//
package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "path"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
  fmt.Println("File Upload Endpoint Hit")

  // Parse our multipart form, 10 << 20 specifies a maximum
  // upload of 10 MB files.
  r.ParseMultipartForm(10 << 20)
  // FormFile returns the first file for the given key `myFile`
  // it also returns the FileHeader so we can get the Filename,
  // the Header and the size of the file
  file, handler, err := r.FormFile("myFile")
  if err != nil {
    fmt.Println("Error Retrieving the File")
    fmt.Println(err)
    return
  }
  defer file.Close()
  fmt.Printf("Uploaded File: %+v\n", handler.Filename)
  fmt.Printf("File Size: %+v\n", handler.Size)
  fmt.Printf("MIME Header: %+v\n", handler.Header)

  // Create a temporary file within our temp-images directory that follows
  // a particular naming pattern
  tempFile, err := ioutil.TempFile(".", "*-" + path.Base(path.Clean(handler.Filename)))
  if err != nil {
    fmt.Println(err)
  }
  defer tempFile.Close()

  // read all of the contents of our uploaded file into a
  // byte array
  fileBytes, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
  }
  // write this byte array to our temporary file
  tempFile.Write(fileBytes)
  // return that we have successfully uploaded our file!
  doc := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <meta http-equiv="refresh" content="3;URL='/'" />
    <title>Document</title>
  </head>
  <body>
    <h3>Successfully uploaded file</h3>
    <p>Go to <a href="/">main page</a> to upload another one.
  </body>
</html>
 `
  fmt.Fprintf(w, doc)
}

func index(w http.ResponseWriter, r *http.Request) {
  doc := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Document</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="/upload"
      method="post"
    >
      <input type="file" name="myFile" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>`

  fmt.Fprintf(w, doc);
}

func setupRoutes() {
  http.HandleFunc("/", index)
  http.HandleFunc("/upload", uploadFile)
  http.ListenAndServe(":8080", nil)
}

func main() {
  fmt.Println("File upload server started")
  setupRoutes()
}
