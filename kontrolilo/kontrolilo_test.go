package kontrolilo

import (
	"testing"
)

func TestValidajVortoj(t *testing.T) {
	vortoj := []string{"marika", "runa", "silika", "gesmi", "druva"}
	for _, vorto := range vortoj {
		err := KontroliVorton(vorto)
		if err != nil {
			t.Errorf("Eraron ricevis: %v, sed %v estas valida", err, vorto)
		}
	}
}

func TestNevalidajVortoj(t *testing.T) {
	vortoj := []string{"manteia", "ana", "sra", "arra"}
	for _, vorto := range vortoj {
		err := KontroliVorton(vorto)
		if err == nil {
			t.Errorf("%v estas nevalida, sed neniun eraron", vorto)
		}
	}
}

func TestIPA(t *testing.T) {
	testo1 := "marika"
	prava := "mə'rikə"
	rezulto, err := IgiIPA(testo1)
	if err != nil {
		t.Errorf("Malatendita eraro de marika: %v", err.Error())
	}
	if len(rezulto) != 1 {
		t.Errorf("Malprava kvanto da vortoj: atendita 1, ricevis %v", len(rezulto))
	}
	if rezulto[0] != prava {
		t.Errorf("Malprava IPA: atendita %v, ricevis %v", prava, rezulto[0])
	}

	testo2 := "sagi to darena ʃona siri fora"
	prava2 := []string{
		"sagi",
		"to",
		"də'reɪnə",
		"ʃonə",
		"siri",
		"forə",
	}
	rezulto2, err2 := IgiIPA(testo2)
	if err2 != nil {
		t.Errorf("Malatendita eraro de %v: %v", testo2, err.Error())
	}
	if len(rezulto2) != len(prava2) {
		t.Errorf("Malprava kvanto da vortoj: atendita %v, ricevis %v", len(prava2), len(rezulto))
	}
	for i, vorto := range prava2 {
		if vorto != rezulto2[i] {
			t.Errorf("Malprava vorto: atendita %v, ricevis %v", vorto, rezulto2[i])
		}
	}
}