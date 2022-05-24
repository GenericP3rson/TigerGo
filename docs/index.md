# Welcome to TigerGo!

A TigerGraph Wrapper in Golang

## Quickstart

First, import the TigerGo library.

```
go get github.com/GenericP3rson/TigerGo
```

Next, create a solution on TigerGraph cloud, then connect to the solution in your Go file.

```
package main

import "github.com/GenericP3rson/TigerGo"

func main() {
   conn := TigerGo.TigerGraphConnection(
      Token: "TOKEN",
      Host: "https://SUBDOMAIN.i.tgcloud.io",
      GraphName: "GRAPHNAME",
      Username: "tigergraph",
      Password: "PASSWORD"
   )
}
```

If you do not already have a token, you can leave the field blank, generate a token using `conn.GetToken()`, then replace the `Token` field.

```
package main

import(
   "fmt"
   "github.com/GenericP3rson/TigerGo"
)

func main() {
   conn := TigerGo.TigerGraphConnection(
      Token: "", // Leaving it empty for now
      Host: "https://SUBDOMAIN.i.tgcloud.io",
      GraphName: "GRAPHNAME",
      Username: "tigergraph",
      Password: "PASSWORD"
   )
   fmt.Println(conn.GetToken())
}
```

## Blogs and Tutorials

Check out a few blogs and tutorials using TigerGo!

- [An Introduction to Using TigerGraph with Go: Exploring COVID-19 Patient Cases](https://towardsdatascience.com/an-introduction-to-using-tigergraph-with-go-exploring-covid-19-patient-cases-f2c0e45849e4)
