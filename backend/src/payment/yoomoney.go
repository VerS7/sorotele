package payment

import (
	"bytes"
	"net/http"
	"net/url"
	"strconv"
)

type YooMoney struct {
	config YooMoneyConfig
}

type YooMoneyConfig struct {
	Token      string
	Reciever   string
	SuccessUrl string
}

type QuickpayForm struct {
	Form        string
	Target      string
	PaymentType string
	Label       string
	Sum         int
}

func YooMoneyInit(c YooMoneyConfig) (*YooMoney, error) {
	return &YooMoney{
		config: c,
	}, nil
}

func (y *YooMoney) Quickpay(form QuickpayForm) (*url.URL, error) {
	var baseUrl string = "https://yoomoney.ru/quickpay/confirm.xml"
	data := url.Values{
		"receiver":      []string{y.config.Reciever},
		"successUrl":    []string{y.config.SuccessUrl},
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
