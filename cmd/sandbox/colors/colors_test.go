package colors

import "testing"

func TestColorize(t *testing.T) {
	word, err := Colorize("Привет")
	if err != nil {
		t.Error("Error: ", err)
	}
	if word != "Привет is rainbow" {
		t.Errorf("Цвет не тот")
	}
}
