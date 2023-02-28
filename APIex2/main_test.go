package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDisbursement(t *testing.T) {
	bodyMock := strings.NewReader(`{"rq_body": {"disbursement_amount": 35000.00,"number_of_payment": 4,"cal_date": "2020-01-05","payment_frequency": 1,"payment_unit": "M","account_number": 5674577629}}`)
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
	data, err := ioutil.ReadAll(res.Body)
	//fmt.Println(data)
	var mockRespond string
	mockRespond = `{"rs_body":{"installment_amount":8795.62,"promotion_name":"Promo1","interest_rate":0.025,"account_number":5674577629}}
`
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} /*
		if string(data) != mockRespond {
			t.Errorf("expected : %s\n\t\t  got : %v", mockRespond, string(data))
		}*/
	assert.Equal(t, mockRespond, string(data))
}

//{ "rs_body": { "installment_amount": 8795.62, "promotion_name": "Promo1", "interest_rate": 0.025, "account_number": 5674577629 } }
