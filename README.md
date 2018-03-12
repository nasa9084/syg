# syg
[![Build Status](https://travis-ci.org/nasa9084/syg.svg?branch=master)](https://travis-ci.org/nasa9084/syg)
[![GoDoc](https://godoc.org/github.com/nasa9084/syg?status.svg)](https://godoc.org/github.com/nasa9084/syg)
-----

a very simple signal handler for golang

## SYNOPSIS

``` go
func ExampleSyg() {
    s := &http.Server{}
    waitCh := make(chan struct{})

    syg.Listen(func(os.Signal) {
        s.Shutdown(context.Background())
        close(waitCh)
    }, os.Interrupt)

    if err := s.ListenAndServe(); err != http.ErrServerClosed {
        // some error handling
    }
    <-waitCh
}
```
