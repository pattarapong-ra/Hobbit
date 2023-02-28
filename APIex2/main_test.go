package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessCalculateDisbursement(t *testing.T) {
	var mockReq requestMessage
	mockReq.ReqBody.DisbursementAmount = 35000.00
	mockReq.ReqBody.NumberOfPayment = 4
	mockReq.ReqBody.CalculateDate = "2020-01-06"
	mockReq.ReqBody.PaymentFrequency = 1
	mockReq.ReqBody.PaymentUnit = "M"
	mockReq.ReqBody.AccountNumber = 5674577629
	mockReqFormatted, err := json.Marshal(&mockReq)
	if err != nil {
		fmt.Println(err)
	}
	bodyMock := strings.NewReader(string(mockReqFormatted))
	req := httptest.NewRequest(http.MethodPost, "/dloan-payment/v1/calculate-installment-amount", bodyMock)
	/*{
		"rq_body": {
			"disbursement_amount": 35000.00,
			"number_of_payment": 4,
			"cal_date": "2020-01-05",
			"payment_frequency": 1,
			"payment_unit": "M",
			"account_number": 5074577629
		}
	}*/
	w := httptest.NewRecorder()
	calculateInstallmentAmount(w, req)
	res := w.Result()
	defer res.Body.Close()
	//fmt.Println(data)
	var mockRes respondMessage
	var actualRes respondMessage
	mockRes.ResBody.InstallmentAmount = 8795.62
	mockRes.ResBody.InterestRate = 0.025
	mockRes.ResBody.PromotionName = "Promo1"
	mockRes.ResBody.AccountNumber = 5674577629
	errorRequestDecode := json.NewDecoder(res.Body).Decode(&actualRes)
	if errorRequestDecode != nil {
		t.Errorf("expected error to be nil got %v", errorRequestDecode)
	} /*
		if string(data) != mockRespond {
			t.Errorf("expected : %s\n\t\t  got : %v", mockRespond, string(data))
		}*/
	status := assert.Equal(t, mockRes, actualRes)
	if status == true {
	}
}

func TestFailedCalculateDisbursement(t *testing.T) {
	var mockReq requestMessage
	/*mockReq.ReqBody.DisbursementAmount=35000.00
	mockReq.ReqBody.NumberOfPayment=4
	mockReq.ReqBody.CalculateDate="2020-01-06"
	mockReq.ReqBody.PaymentFrequency=1
	mockReq.ReqBody.PaymentUnit=""*/
	mockReqFormatted, err := json.Marshal(&mockReq)
	if err != nil {
		fmt.Println(err)
	}
	bodyMock := strings.NewReader(string(mockReqFormatted))
	req := httptest.NewRequest(http.MethodPost, "/dloan-payment/v1/calculate-installment-amount", bodyMock)
	/*{
		"rq_body": {
			"disbursement_amount": 35000.00,
			"number_of_payment": 4,
			"cal_date": "2020-01-05",
			"payment_frequency": 1,
			"payment_unit": "M",
			"account_number": 5074577629
		}
	}*/
	w := httptest.NewRecorder()
	calculateInstallmentAmount(w, req)
	res := w.Result()
	defer res.Body.Close()
	//fmt.Println(data)
	var mockRes respondMessage
	var actualRes respondMessage
	mockRes.ResBody.InstallmentAmount = 0
	mockRes.ResBody.InterestRate = 0
	mockRes.ResBody.PromotionName = ""
	mockRes.ResBody.AccountNumber = 0
	errorRequestDecode := json.NewDecoder(res.Body).Decode(&actualRes)
	if errorRequestDecode != nil {
		fmt.Println(errorRequestDecode)
		t.Errorf("expected error to be nil got %v", errorRequestDecode)
	} /*
		if string(data) != mockRespond {
			t.Errorf("expected : %s\n\t\t  got : %v", mockRespond, string(data))
		}*/
	status := assert.Equal(t, mockRes, actualRes)
	if status == true {
		fmt.Println("equal ja")
	} else {
		fmt.Println("not equal ja")
	}
}
