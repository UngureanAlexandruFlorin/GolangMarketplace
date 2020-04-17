package main

import (
    "fmt"
    "bytes"
    "os"
    "net/http"
     "golang.org/x/oauth2/jwt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func userHandler (requestWriter http.ResponseWriter, req *http.Request) {
	var bufferSize int = 10;
	var result bytes.Buffer;

	file, error := os.Open("./static/index.html");
    check(error);

    var buffer []byte = make([]byte, bufferSize);
    readBytes, error := file.Read(buffer);
    check(error);

    for readBytes > 0 {
    	readBytes, error = file.Read(buffer);

    	if (readBytes == 0) {
    		continue;
    	}

    	check(error);

    	result.WriteString(string(buffer[:readBytes]));
    }

    fmt.Fprintf(requestWriter, result.String());
    
}

func main() {

    http.HandleFunc("/", userHandler);
    http.ListenAndServe(":8080", nil);

    fmt.Printf("Server started on port 8080!");

}