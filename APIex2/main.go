package main

import (
	"database/sql"
	"log"
	"net/http"

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
	psqlconn = "host=localhost port=5432 user=mooham12314 password=mooham12315 dbname=postgres sslmode=disable"
)

type requestMessage struct {
	ReqBody requestBody `json:"rq_body"`
}

type respondMessage struct {
	ResBody respondBody `json:"rs_body"`
}

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

func main() {
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	PrepareTable(db)

	r := mux.NewRouter()
	r.HandleFunc("/dloan-payment/v1/calculate-installment-amount", calculateInstallmentAmount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8089", r))
}
