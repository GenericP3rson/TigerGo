# TigerGo

Welcome to TigerGo, a Go TigerGraph wrapper!

## Quickstart

To get started, get the package in Go.

```
go get github.com/GenericP3rson/TigerGo
```

You can then import the library in your code and create a TigerGraph connection. If you do nnot immediately have a token, that's okay! You can leave it blank, run `conn.GetToken`, then update the token!

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

## Blogs

Check out an introduction to some of the features of this library with TigerGraph's COVID-19 Starter Kit [here](https://medium.com/@shreya-chaudhary/an-introduction-to-using-tigergraph-with-go-exploring-covid-19-patient-cases-f2c0e45849e4)!