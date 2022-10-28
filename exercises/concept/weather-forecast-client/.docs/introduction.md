# HTTP

HTTP or Hypertext Transfer Protocol is an application layer protocol. When you write a URL address in format of `http://website-address`, your browser is acting as a client and uses Hypertext Transfer Protocol to ask a server for resources and uses those resources to view the webpage.
In this concept we will learn how to code HTTP clients and send requests to http servers using go `net/http` package.

## Client

A client can be defined by using `http.Client` structure.
There are three main methods that you can use in http/net package to act as an HTTP client:

### Get method

`func (c *Client) Get(url string) (resp *Response, err error)`
This method lets you send GET requests to a URL and receive the response or any error that might occur. Consider the example below:

```go
package main

import (
 "fmt"
 "io"
 "log"
 "net/http"
)

func main() {
 client := &http.Client{}
 resp, err := client.Get("https://www.google.com/") // Sending a GET request to google
 if err != nil {
  log.Fatal(err)
 }
 defer resp.Body.Close() // Always close the reponse body
 body, err := io.ReadAll(resp.Body) // Reading the response HTTP body
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println(string(body)) // Printing out the response body
}
```

Here we connect to Google main webpage and display the response body.

~~~~exercism/caution
Tip: You should always close the response body. Since Go uses HTTP/1.1, it keeps a connection to the server open in case you might reuse it. Therefore using `defer resp.Body.Close()` you can make sure the response body will be closed after your function exits.  
~~~~