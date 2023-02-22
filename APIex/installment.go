package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*{
	"rq_body": {
	  "disbursement_amount": 9.22,
	  "number_of_payment": 1,
	  "interest_rate": 2.50000,
	  "payment_frequency": 1,
	  "payment_unit": "M"
	}
  }*/
type rq_body struct {
	DisbursementAmount string `json:"disbursement_amount"` //decimal.Dec2
	NumberOfPayment    string `json:"number_of_payment"`
	InterestRate       string `json:"interest_rate"`
	PaymentFrequency   string `json:"payment_frequency"`
	PaymentUnit        string `json:"payment_unit"`
}

type rs_body struct {
	InstallmentAmount string `json:"installment_amount"`
}

var installments []rs_body

func calculateInstallmentAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var installmentRespond rs_body
	var installmentRequest rq_body
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Println(r)
	_ = json.NewDecoder(r.Body).Decode(&installmentRequest)
	installmentRequest.DisbursementAmount = params["disbursement-amount"]
	/*tempInterest := installmentRequest.InterestRate
	tempDisbursement := installmentRequest.DisbursementAmount
	tempFrequency := installmentRequest.PaymentFrequency*/
	interest, _ := strconv.ParseFloat(installmentRequest.InterestRate, 64)
	disbursement, _ := strconv.ParseFloat(installmentRequest.DisbursementAmount, 64)
	frequency, _ := strconv.ParseFloat(installmentRequest.PaymentFrequency, 64)
	/*interest, _ := strconv.ParseFloat(params["interest_rate"], 64)
	disbursement, _ := strconv.ParseFloat(params["disbursement_amount"], 64)
	frequency, _ := strconv.ParseFloat(params["payment_frequency"], 64)*/
	fmt.Println(params["interest_rate"])
	fmt.Println(params["disbursement_amount"])
	fmt.Println(params["payment_frequency"])
	fmt.Println(interest)
	fmt.Println(disbursement)
	fmt.Println(frequency)
	tempRespond := fmt.Sprintf("%18.2f", disbursement/((1-(1/(math.Pow(1+interest/12, frequency))))/(interest/12)))
	installmentRespond.InstallmentAmount = tempRespond
	fmt.Println(installmentRespond)
	installmentRespond.InstallmentAmount = "12345"
	installments = append(installments, installmentRespond)
	json.NewEncoder(w).Encode(&installmentRespond)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dloan-payment/v1/calculate-installment-amount", calculateInstallmentAmount).Methods("GET")
	http.ListenAndServe(":8080", router)
}

//http://localhost:8080//dloan-payment/v1/calculate-installment-amount/?disbursement_amount=35000&number_of_payment=4&interest_rate=0.09120&payment_frequency=1&payment_unit=1
