package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mooham12314"
	password = "mooham12315"
	dbname   = "postgres"
	layout   = "2006-01-02"
)

type requestMessage struct {
	ReqBody requestBody `json:"rq_body"`
}

type respondMessage struct {
	ResBody respondBody `json:"rs_body"`
}

//decimal.Dec2
type requestBody struct {
	DisbursementAmount float64 `json:"disbursement_amount"`
	NumberOfPayment    int     `json:"number_of_payment"`
	CalculateDate      string  `json:"cal_date"`
	PaymentFrequency   int     `json:"payment_frequency"`
	PaymentUnit        string  `json:"payment_unit"`
	AccountNumber      int     `json:"account_number"`
}

type respondBody struct {
	InstallmentAmount float64 `json:"installment_amount"`
	PromotionName     string  `json:"promotion_name"`
	InterestRate      float64 `json:"interest_rate"`
	AccountNumber     int     `json:"account_number"`
}

type promotion struct {
	PromoName    string
	InterestRate float64
}

func calculateInstallmentAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var installmentRespond respondMessage
	var installmentRequest requestMessage
	var errorRespondFloat error
	errorRequestDecode := json.NewDecoder(r.Body).Decode(&installmentRequest)
	if errorRequestDecode != nil {
		log.Fatal(errorRequestDecode)
	}
	disbursement := installmentRequest.ReqBody.DisbursementAmount
	numberOfPayment := float64(installmentRequest.ReqBody.NumberOfPayment)
	calculateDate := installmentRequest.ReqBody.CalculateDate
	installmentRespond.ResBody.AccountNumber = installmentRequest.ReqBody.AccountNumber
	currentPromo := GetPromo(calculateDate)
	interest := currentPromo.InterestRate / 100
	tempRespond := fmt.Sprintf("%.2f", disbursement/((1-(1/(math.Pow(1+interest/12, numberOfPayment))))/(interest/12)))
	installmentRespond.ResBody.InstallmentAmount, errorRespondFloat = strconv.ParseFloat(tempRespond, 64)
	if errorRespondFloat != nil {
		log.Fatal(errorRespondFloat)
	}

	installmentRespond.ResBody.PromotionName = currentPromo.PromoName
	installmentRespond.ResBody.InterestRate = interest
	installmentRespond.ResBody.AccountNumber = installmentRequest.ReqBody.AccountNumber

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	insertAccountDetail(db, installmentRespond)

	json.NewEncoder(w).Encode(&installmentRespond)
}

func InitializeDB() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()
	defer fmt.Println("Closed!")
	// check db
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected!")
	PrepareTable(db)
}

func main() {
	InitializeDB()
	r := mux.NewRouter()
	r.HandleFunc("/dloan-payment/v1/calculate-installment-amount", calculateInstallmentAmount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8089", r))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
