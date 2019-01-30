
package main
import (
    "log"
    "net/http"
    "os"
    "github.com/gorilla/handlers"
    "./rest-and-go/store"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    router := store.NewRouter()

    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

    log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
