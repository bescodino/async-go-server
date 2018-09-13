
[![Build Status](https://travis-ci.org/brun1/async-go-server.svg?branch=master)](https://travis-ci.org/brun1/async-go-server)

# async-go-server
Simple async server using golang.


This package is very small, it aims to simplify the use of an asynchronous server in go easily and not use ```http.ListenAndServe()``` because it generates blocking

**How to use**

Creation of a simple asynchronous HTTP server:

```
import "github.com/brun1/async-server-go"

server := server.Init(":0")

if err := server.Start(); err != nil {
    panic(err)
}

// ...do stuff here...

server.Stop()

```
