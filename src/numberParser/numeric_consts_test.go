package numberParser

import (
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	if str := normalize("cem" + CONJUNCTION); !strings.Contains(str, "cento") {
		t.Error("\"cem\", not converted to \"cento\"!")
	}
	if str := normalize("um mil"); strings.Contains(str, "um") {
		t.Error("\"um\" in \"um mil\" should be ripped off!")
	}
	if str := normalize("cinquenta e um mil"); !strings.Contains(str, "um") {
		t.Error("\"um\" in \"cinquenta e um mil\" should stay!")
	}
	for key, value := range HUNDREDS.translate {
		if value == "" {
			continue
		}
		if str := normalize("mil " + value); !strings.Contains(str, CONJUNCTION) {
			t.Errorf("Conjunction missing!\nValue: %q", str)
		}
		if str := normalize("mil " + UNITS.Translate(key)); !strings.Contains(str, CONJUNCTION) {
			t.Errorf("Conjunction missing!\nValue: %q", str)
		}
	}

}
