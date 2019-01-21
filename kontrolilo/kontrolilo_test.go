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
