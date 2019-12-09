package server_service_pack

import (
	"fmt"
	"net/http"
)

func Greet_Server_Home( w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w, "Okay now go is working")
}
