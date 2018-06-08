package anventory

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Anventory struct {
	r *mux.Router
}

func (a *Anventory) setupAPI() {
}

// New creates a new anventory instance using the specified settings
func New(s Settings) (*Anventory, error) {
	ret := &Anventory{}
	ret.r = mux.NewRouter()

	ret.setupAPI()

	return ret, nil
}

func (a *Anventory) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	a.r.ServeHTTP(rsp, req)
}
