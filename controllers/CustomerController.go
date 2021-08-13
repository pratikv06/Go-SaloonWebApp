package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratikv06/Go-SaloonWebApp/models"
	"github.com/pratikv06/Go-SaloonWebApp/services"
)

type CustomerController struct {
	CustSRV *services.CustomerServices
}

func NewCustomerController(custsrv *services.CustomerServices) *CustomerController {
	return &CustomerController{CustSRV: custsrv}
}

func (custController *CustomerController) CustomerRoute(route *mux.Router) {
	// fmt.Print("Hi Customer Route")
	route.HandleFunc("/customer", custController.createCustomer).Methods("POST")
	route.HandleFunc("/customer", custController.getCustomer).Methods("GET")
}

func (custController *CustomerController) getCustomer(w http.ResponseWriter, r *http.Request) {

}

func (custController *CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	var custStruct models.Customer
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Data can't be handled"))
		return
	}
	if len(body) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No data is coming"))
		return
	}
	err = json.Unmarshal(body, &custStruct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to parse data"))
		return
	}

	//using json to get input value from body Method 1
	// json.NewDecoder(r.Body).Decode(&custStruct)

	err = custController.CustSRV.AddCustomer(&custStruct)
	if err != nil {
		log.Fatal("Customer not added ", err)
	}
	json.NewEncoder(w).Encode(custStruct)
}
