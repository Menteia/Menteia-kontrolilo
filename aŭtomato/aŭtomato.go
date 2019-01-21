package aŭtomato

import (
	"errors"
	"fmt"
	"strings"
)

type Stato int

const (
	malplena Stato = iota
	k1             = iota
	k2             = iota
	k3             = iota
	k1k2           = iota
	v              = iota
	f              = iota
)

type LiteroTipo int

const (
	vokalo      LiteroTipo = iota
	konsonanto1            = iota
	konsonanto2            = iota
	konsonanto3            = iota
	fino                   = iota
	nenio                  = iota
)

var transiroj = map[Stato]map[LiteroTipo]Stato{
	malplena: map[LiteroTipo]Stato{
		konsonanto1: k1,
		konsonanto2: k2,
		konsonanto3: k3,
		fino:        k3,
	},
	k1: map[LiteroTipo]Stato{
		konsonanto2: k1k2,
		vokalo:      v,
	},
	k2: map[LiteroTipo]Stato{
		vokalo: v,
	},
	k3: map[LiteroTipo]Stato{
		vokalo: v,
	},
	k1k2: map[LiteroTipo]Stato{
		vokalo: v,
	},
	v: map[LiteroTipo]Stato{
		konsonanto1: k1,
		konsonanto2: k2,
		konsonanto3: k3,
		fino:        f,
	},
	f: map[LiteroTipo]Stato{
		konsonanto1: k1,
		konsonanto2: k2,
		konsonanto3: k3,
		vokalo:      v,
	},
}

var finajStatoj = map[Stato]struct{}{
	v: struct{}{},
	f: struct{}{},
}

type FiniaAŭtomato struct {
	stato Stato
}

func Krei() FiniaAŭtomato {
	return FiniaAŭtomato{
		stato: malplena,
	}
}

func (fa *FiniaAŭtomato) Restartigi() {
	fa.stato = malplena
}

func (fa *FiniaAŭtomato) Movi(novaLitero rune) error {
	tipo := troviTipon(novaLitero)
	sekvaj, trovita := transiroj[fa.stato]
	if trovita {
		novaStato, trovita2 := sekvaj[tipo]
		if trovita2 {
			fa.stato = novaStato
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Nevalida sekva litero: %v", novaLitero))
}

func (fa *FiniaAŭtomato) ĈuFinita() bool {
	_, trovita := finajStatoj[fa.stato]
	return trovita
}

func ĈuFinaLitero(litero rune) bool {
	return strings.ContainsRune("mns", litero)
}

func ĈuVokalaLitero(litero rune) bool {
	return strings.ContainsRune("aeiou", litero)
}

func troviTipon(litero rune) LiteroTipo {
	if strings.ContainsRune("pbtdkgf", litero) {
		return konsonanto1
	} else if strings.ContainsRune("rl", litero) {
		return konsonanto2
	} else if ĈuFinaLitero(litero) {
		return fino
	} else if strings.ContainsRune("vʃ", litero) {
		return konsonanto3
	} else if ĈuVokalaLitero(litero) {
		return vokalo
	}
	return nenio
}
