package requests

import (
	"net/http"
	"io/ioutil"
)

func readRequest(r *http.Request) ([]byte, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	return b, err 
}
  
