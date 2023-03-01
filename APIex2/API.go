package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func calculateInstallmentAmount(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var installmentRespond respondMessage
	var installmentRequest requestMessage

	err := json.NewDecoder(r.Body).Decode(&installmentRequest)
	checkError(err)

	disbursement := installmentRequest.ReqBody.DisbursementAmount
	numberOfPayment := float64(installmentRequest.ReqBody.NumberOfPayment)
	calculateDate := installmentRequest.ReqBody.CalculateDate
	installmentRespond.ResBody.AccountNumber = installmentRequest.ReqBody.AccountNumber
	currentPromo := GetPromo(calculateDate)
	interest := currentPromo.InterestRate / 100

	tempRespond := fmt.Sprintf("%.2f", disbursement/((1-(1/(math.Pow(1+interest/12, numberOfPayment))))/(interest/12)))

	installmentRespond.ResBody.InstallmentAmount, err = strconv.ParseFloat(tempRespond, 64)
	checkError(err)

	installmentRespond.ResBody.PromotionName = currentPromo.PromoName
	installmentRespond.ResBody.InterestRate = interest
	installmentRespond.ResBody.AccountNumber = installmentRequest.ReqBody.AccountNumber

	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	insertAccountDetail(db, installmentRespond)
	db.Close()

	json.NewEncoder(w).Encode(&installmentRespond)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
