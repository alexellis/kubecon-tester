package function

import "testing"

const sep = " "

func Test_ValidHash_GivesTrue(t *testing.T) {
	req := "Alex" + sep + "$2a$10$2UEqmAEoGiif4KyOws.5.exi1XfWnpiK2XeQKQum4k9gLLkemJREa"
	decoded := decode([]byte(req))
	if decoded != "true" {
		t.Errorf("decode (%s), want: %s, got: %s", req, "true", decoded)
	}
}

func Test_InvalidHash_GivesFalse(t *testing.T) {
	req := "alex" + sep + "$2a$10$2UEqmAEoGiif4KyOws.5.exi1XfWnpiK2XeQKQum4k9gLLkemJREa"
	decoded := decode([]byte(req))
	if decoded != "false" {
		t.Errorf("decode (%s), want: %s, got: %s", req, "false", decoded)
	}
}

func Test_DefaultComplexity_GivesValidHash(t *testing.T) {

	input := "Alex"
	encoded := Handle([]byte(input))
	decodedValue := decode([]byte(input + sep + encoded))
	wantDecoded := "true"
	if decodedValue != wantDecoded {
		t.Errorf("decode, want: %s, got: %s", wantDecoded, decodedValue)
	}
}
