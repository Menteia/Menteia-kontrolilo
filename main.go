package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"menteia/kontrolilo"
)

type Peto struct {
	Vorto string
	IPA   string
}

type Respondo struct {
	ĈuValida bool
	Kialo    string
	IPA      []string
}

func vorto(peto Peto) (Respondo, error) {
	kialo := kontrolilo.KontroliVorton(peto.Vorto)
	if kialo == nil {
		return Respondo{ĈuValida: true, Kialo: ""}, nil
	} else {
		return Respondo{ĈuValida: false, Kialo: kialo.Error()}, nil
	}
}

func IPA(peto Peto) (Respondo, error) {
	rezulto, kialo := kontrolilo.IgiIPA(peto.IPA)
	if kialo == nil {
		return Respondo{ĈuValida: true, Kialo: "", IPA: rezulto}, nil
	} else {
		return Respondo{ĈuValida: false, Kialo: kialo.Error(), IPA: []string{}}, nil
	}
}

func Funkcio(ctx context.Context, peto Peto) (Respondo, error) {
	if len(peto.Vorto) > 0 {
		return vorto(peto)
	} else if len(peto.IPA) > 0 {
		return IPA(peto)
	}
	return Respondo{}, errors.New("Nevalida peto")
}

func main() {
	lambda.Start(Funkcio)
}
