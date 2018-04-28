package function

import "testing"

func Test_ValidHash_GivesTrue(t *testing.T) {
	decoded := decode([]byte("Alex $2a$10$2UEqmAEoGiif4KyOws.5.exi1XfWnpiK2XeQKQum4k9gLLkemJREa"))
	if decoded != "true" {
		t.Errorf("decode, want: %s, got: %s", "true", decoded)
	}
}

func Test_InvalidHash_GivesFalse(t *testing.T) {
	decoded := decode([]byte("alex $2a$10$2UEqmAEoGiif4KyOws.5.exi1XfWnpiK2XeQKQum4k9gLLkemJREa"))
	if decoded != "false" {
		t.Errorf("decode, want: %s, got: %s", "false", decoded)
	}
}
