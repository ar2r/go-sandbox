package words

import "testing"

func TestGetWord(t *testing.T) {
	word := GetWord()
	if word != "SecretWord" {
		t.Errorf("Не то слово")
	}
}
