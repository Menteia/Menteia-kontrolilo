package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"menteia/kontrolilo"
)

type Peto struct {
	Vorto string
}

type Respondo struct {
	ĈuValida bool
	Kialo    string
}

func Funkcio(ctx context.Context, peto Peto) (Respondo, error) {
	kialo := kontrolilo.KontroliVorton(peto.Vorto)
	if kialo == nil {
		return Respondo{ĈuValida: true, Kialo: ""}
	} else {
		return Respondo{ĈuValida: false, Kialo: kialo}
	}
}
