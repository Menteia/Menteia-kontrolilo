package kontrolilo

import (
	"errors"
	"fmt"
	"menteia/aŭtomato"
	"strings"
)

func KontroliVorton(vorto string) error {
	aŭtomato := aŭtomato.Krei()
	for _, litero := range vorto {
		err := aŭtomato.Movi(litero)
		if err != nil {
			return err
		}
	}

	if !aŭtomato.ĈuFinita() {
		return errors.New("Ne finita")
	}

	return nil
}

func IgiIPA(eniro string) ([]string, error) {
	vortoj := strings.Split(eniro, " ")
	rezulto := make([]string, 0, len(vortoj))
	aŭtomato := aŭtomato.Krei()

	for _, vorto := range vortoj {
		silaboj, err := aŭtomato.Dividi(vorto)
		if (err != nil) {
			return []string{}, err
		}
		var ipa string
		switch len(silaboj) {
		case 1:
			ipa = silaboj[0]
		case 2:
			ipa = fmt.Sprintf("%v%v", silaboj[0], strings.Replace(silaboj[1], "a", "ə", 1))
		case 3:
			ipa = fmt.Sprintf(
				"%v'%v%v",
				strings.Replace(silaboj[0], "a", "ə", 1),
				strings.Replace(silaboj[1], "e", "eɪ", 1),
				strings.Replace(silaboj[2], "a", "ə", 1),
			)
		}
		rezulto = append(rezulto, ipa)
	}

	return rezulto, nil
}
