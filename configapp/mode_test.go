package configapp

import (
	"fmt"
	"testing"
)

func TestValidMode(t *testing.T) {

	t.Run("mode [dev] deve retornar nil", func(t *testing.T) {
		data := modeDev

		var expected error
		result := validMode(data)

		if expected != result {
			t.Errorf("\n[ERROR]\nExpected: %v\nResult: %v", expected, result)
		}
	})

	t.Run("mode [homolog] deve retornar nil", func(t *testing.T) {
		data := modeHomolog

		var expected error
		result := validMode(data)

		if expected != result {
			t.Errorf("\n[ERROR]\nExpected: %v\nResult: %v", expected, result)
		}
	})

	t.Run("mode [prod] deve retornar nil", func(t *testing.T) {
		data := modeProd

		var expected error
		result := validMode(data)

		if expected != result {
			t.Errorf("\n[ERROR]\nExpected: %v\nResult: %v", expected, result)
		}
	})

	t.Run("mode [xxx] deve retornar erro", func(t *testing.T) {
		data := "xxx"

		expected := fmt.Errorf("mode.properties invalid: xxx")
		result := validMode(data)

		if expected.Error() != result.Error() {
			t.Errorf("\n[ERROR]\nExpected: %v\nResult: %v", expected, result)
		}
	})
}
