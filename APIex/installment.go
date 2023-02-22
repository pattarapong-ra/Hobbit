package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type requestMessage struct {
	ReqBody requestBody `json:"rq_body"`
}

type respondMessage struct {
	ResBody respondBody `json:"rs_body"`
	Err     string      `json:"error_body"`
}

//decimal.Dec2
type requestBody struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	InterestRate       float64 `json:"interest_rate"`
	PaymentFrequency   int     `json:"payment_frequency"`
	PaymentUnit        string  `json:"payment_unit"`
}

type respondBody struct {
	InstallmentAmount float64 `json:"installment_amount"`
}

func calculateInstallmentAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var installmentRespond respondMessage
	var installmentRequest requestMessage
	errorDecode := json.NewDecoder(r.Body).Decode(&installmentRequest)
	if errorDecode != nil {
		installmentRespond.Err = string(errorDecode.Error())
	}
	interest := installmentRequest.ReqBody.InterestRate / 100 //percentage convert
	disbursement := installmentRequest.ReqBody.DisbursementAmount
	numberOfPayment := float64(installmentRequest.ReqBody.NumberOfPayment)
	tempRespond := fmt.Sprintf("%.2f", disbursement/((1-(1/(math.Pow(1+interest/12, numberOfPayment))))/(interest/12)))
	installmentRespond.ResBody.InstallmentAmount, _ = strconv.ParseFloat(tempRespond, 64)
	json.NewEncoder(w).Encode(&installmentRespond)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dloan-payment/v1/calculate-installment-amount", calculateInstallmentAmount).Methods("GET")
	http.ListenAndServe(":8080", router)
}
