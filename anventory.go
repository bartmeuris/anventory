package anventory

import (
	"github.com/gorilla/mux"
	"net/http"
)


type Anventory struct {
	r *mux.Router
}

func (a *Anventory) setupAPI() {
}

func New(s Settings) (*Anventory, error) {
	ret := &Anventory{}
	ret.r = mux.NewRouter()
	
	ret.setupAPI()

	return ret, nil
}

func (a *Anventory) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	a.r.ServeHTTP(rsp, req)
}

