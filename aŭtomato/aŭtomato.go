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
		fino:        k3,
		vokalo:      v,
	},
}

var finajStatoj = map[Stato]struct{}{
	v: struct{}{},
	f: struct{}{},
}

type FiniaAŭtomato struct {
	stato        Stato
	antaŭeFinita bool
}

func Krei() FiniaAŭtomato {
	return FiniaAŭtomato{
		stato:        malplena,
		antaŭeFinita: false,
	}
}

func (fa *FiniaAŭtomato) Restartigi() {
	fa.stato = malplena
	fa.antaŭeFinita = false
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
	return errors.New(fmt.Sprintf("Nevalida sekva litero: %v", string(novaLitero)))
}

func (fa *FiniaAŭtomato) Dividi(vorto string) ([]string, error) {
	defer fa.Restartigi()

	silaboj := make([]string, 0, 3)
	literoj := []rune(vorto)
	var aktualaSilabo strings.Builder
	for i, litero := range literoj {
		err := fa.Movi(litero)
		if err != nil {
			return []string{}, err
		}
		if fa.ĈuFinita() {
			if fa.antaŭeFinita {
				if i < len(literoj)-1 && ĈuVokalaLitero(literoj[i+1]) {
					silaboj = append(silaboj, aktualaSilabo.String())
					aktualaSilabo.Reset()
					aktualaSilabo.WriteRune(litero)
				} else {
					aktualaSilabo.WriteRune(litero)
					silaboj = append(silaboj, aktualaSilabo.String())
					aktualaSilabo.Reset()
				}
			} else {
				aktualaSilabo.WriteRune(litero)
			}
			fa.antaŭeFinita = true
		} else {
			if fa.antaŭeFinita && aktualaSilabo.Len() > 0 {
				silaboj = append(silaboj, aktualaSilabo.String())
				aktualaSilabo.Reset()
			}
			aktualaSilabo.WriteRune(litero)
			fa.antaŭeFinita = false
		}
	}

	if fa.ĈuFinita() {
		if aktualaSilabo.Len() > 0 {
			silaboj = append(silaboj, aktualaSilabo.String())
		}
		if len(silaboj) > 3 {
			return []string{}, errors.New(fmt.Sprintf("Tro da silaboj en %v", vorto))
		}
		return silaboj, nil
	} else {
		return []string{}, errors.New("Ne finita")
	}
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
