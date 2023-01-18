package numberParser

import (
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	if str := normalize("cem" + CONJUNCTION); !strings.Contains(str, "cento") {
		t.Errorf("%q, not converted to %q!", "cem", "cento")
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
