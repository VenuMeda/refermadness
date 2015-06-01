package controllers

import (
  "github.com/gorilla/mux"
  "github.com/larryprice/refermadness/utils"
  "html/template"
  "net/http"
  "encoding/json"
  "io/ioutil"

  "fmt"
)

type CreateServiceControllerImpl struct {
  currentUser utils.CurrentUserAccessor
  basePage utils.BasePageCreator
}

func NewCreateServiceController(currentUser utils.CurrentUserAccessor, basePage utils.BasePageCreator) *CreateServiceControllerImpl {
  return &CreateServiceControllerImpl{
    currentUser: currentUser,
    basePage: basePage,
  }
}

func (sc *CreateServiceControllerImpl) Register(router *mux.Router) {
  router.HandleFunc("/service/create", sc.view).Methods("GET")
  router.HandleFunc("/service/create", sc.create).Methods("POST")
}

func (sc *CreateServiceControllerImpl) view(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("views/layout.html", "views/create-service.html")
  t.Execute(w, sc.basePage.Get(r))
}

func (sc *CreateServiceControllerImpl) create(w http.ResponseWriter, r *http.Request) {
  var serviceData map[string]string
  body, _ := ioutil.ReadAll(r.Body)
  if err := json.Unmarshal(body, &serviceData); err != nil {
    fmt.Println("There was an error parsing json!", err)
    return
  }

  if serviceData["name"] == "" || serviceData["description"] == "" || serviceData["url"] == "" {
    fmt.Println("bad data sent to create")
    return
  }

  fmt.Println(serviceData)
}