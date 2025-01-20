package payment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type YooMoney struct {
	Config YooMoneyConfig
}

type YooMoneyConfig struct {
	Token             string
	Secure            string
	Reciever          string
	SuccessUrl        string
	RecievePaymentUrl string
}

type QuickpayForm struct {
	Form        string
	Target      string
	PaymentType string
	Label       string
	Sum         int
}

type Operation struct {
	OperationID string    `json:"operation_id"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Label       string    `json:"label"`
	Amount      float64   `json:"amount"`
	Datetime    time.Time `json:"datetime"`
}

func YooMoneyInit(c YooMoneyConfig) (*YooMoney, error) {
	return &YooMoney{
		Config: c,
	}, nil
}

func (y *YooMoney) Quickpay(form QuickpayForm) (*url.URL, error) {
	var baseUrl string = "https://yoomoney.ru/quickpay/confirm.xml"
	data := url.Values{
		"receiver":      []string{y.Config.Reciever},
		"successUrl":    []string{y.Config.SuccessUrl},
		"quickpay-form": []string{form.Form},
		"paymentType":   []string{form.PaymentType},
		"targets":       []string{form.Target},
		"label":         []string{form.Label},
		"sum":           []string{strconv.Itoa(form.Sum)},
	}

	resp, err := http.Post(baseUrl, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	return resp.Request.URL, nil
}

func (y *YooMoney) GetHistory(account string) ([]Operation, error) {
	var baseUrl string = "https://yoomoney.ru/api/operation-history"

	client := &http.Client{}

	data := url.Values{
		"label": []string{account},
	}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+y.Config.Token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var ops []Operation
	if err := json.NewDecoder(resp.Body).Decode(&ops); err != nil {
		return nil, err
	}

	return ops, nil
}
