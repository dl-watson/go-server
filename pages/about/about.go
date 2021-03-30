package about

import (
	"fmt"
	"net/http"
)

func Page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}