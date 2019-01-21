package aŭtomato

import (
	"testing"
)

func TestValidajSilaboj(t *testing.T) {
	vortoj := []string{"marika", "runa", "silika", "gesmi", "druva"}
	silaboj := [][]string{
		[]string{"ma", "ri", "ka"},
		[]string{"ru", "na"},
		[]string{"si", "li", "ka"},
		[]string{"ges", "mi"},
		[]string{"dru", "va"},
	}

	aŭtomato := Krei()
	for i, vorto := range vortoj {
		rezulto, err := aŭtomato.Dividi(vorto)
		if err != nil {
			t.Errorf("Malatendita eraro de %v: %v", vorto, err.Error())
		}
		if len(rezulto) != len(silaboj[i]) {
			t.Errorf("Malsukcesis: devus esti: %v, sed %v ricevis", silaboj[i], rezulto)
		} else {
			for i2, silabo := range rezulto {
				if silabo != silaboj[i][i2] {
					t.Errorf("Malsukcesis: devus esti: %v, sed %v ricevis", silaboj[i], rezulto)
				}
			}
		}
	}
}

func TestNevalidajSilaboj(t * testing.T) {
	vortoj := []string{"marikana", "mar", "erika", "mej"}

	aŭtomato := Krei()
	for _, vorto := range vortoj {
		_, err := aŭtomato.Dividi(vorto)
		if err == nil {
			t.Errorf("Malsukcesis: %v ja estas nevalida", vorto)
		}
	}
}