package kontrolilo

import (
  "errors"
	"menteia/aŭtomato"
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
