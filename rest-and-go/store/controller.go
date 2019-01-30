package store

import (
    "encoding/json"
    // "io"
    // "io/ioutil"
//"log"
    // "fmt"
    "net/http"
    // "strings"
    // "strconv"

    // "github.com/gorilla/mux"
    // "github.com/gorilla/context"
    // "github.com/dgrijalva/jwt-go"

)

type Controller struct {
    Repository Repository
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
    products := c.Repository.GetProducts() // list of all products
    // log.Println(products)
    data, _ := json.Marshal(products)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
    return
}