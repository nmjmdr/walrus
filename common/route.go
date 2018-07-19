package common

type Route struct {
  Route string
  Method string
  Handler func(w http.ResponseWriter, r *http.Request)
}
