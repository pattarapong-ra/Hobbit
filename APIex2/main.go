package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pattarapong-ra/Hobbit/tree/main/exercise2/sql/ex3"
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
}

type respondBody struct {
	InstallmentAmount float64 `json:"installment_amount"`
	PromotionName     string  `json:"promotion_name"`
	InterestRate      float64 `json:"interest_rate"`
	AccountNumber     int     `json:"account_number"`
}

type promotion struct {
	PromoName string
	Start_date string
	End_date string
}

type rate struct{
	PromoName string
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
	//calculate interest from promo date
	

	calculateDate :=installmentRequest.ReqBody.CalculateDate
    t, err := time.Parse(layout, calculateDate)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println(t)

	interest := installmentRequest.ReqBody.InterestRate / 100 //percentage convert
	disbursement := installmentRequest.ReqBody.DisbursementAmount
	numberOfPayment := float64(installmentRequest.ReqBody.NumberOfPayment)
	tempRespond := fmt.Sprintf("%.2f", disbursement/((1-(1/(math.Pow(1+interest/12, numberOfPayment))))/(interest/12)))
	installmentRespond.ResBody.InstallmentAmount, errorRespondFloat = strconv.ParseFloat(tempRespond, 64)
	if errorRespondFloat != nil {
		log.Fatal(errorRespondFloat)
	}
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
	r.HandleFunc("/get/swapSalary", ex3.SwapSalary).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
