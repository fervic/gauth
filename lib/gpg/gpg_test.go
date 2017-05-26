package gpg_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/fervic/gauth/lib/gpg"
)

const (
	helloWorld = `-----BEGIN PGP MESSAGE-----

hQIMA5BMU610+Dk4ARAAlnZMHg8ls3LbvWFz8dGrHUxlUK7TCWrQekpCZ5V3b0Gi
vuAhnXMCQ3k8NvhHrO9YIw/vJXq5FC4+yLLSVpyuYDK9DEuaD5qhs9GwF/WH5kO9
b6CHf7xpKPIS6ISP5G/V4TTZEJ73QY6WIirWJbKZaxcLxrUn8ddkAZjQrrYnoCHy
mWChVygQ6D0St+dMIFhF3fZz/2llQcm2qU6M8RdOL6cQR3nzl2IAO4pkAcSDXa0V
BAQ9Ek62Fam9pA8ZXqbVoGm1aRgW0zvs6aKEnmSHypveQySNDD2LMpBffCNw3IHI
4+42j8+lMuHbe2KZf6LrMmlQ9DnJKoG9Xo3CdI7HzqtfIbGnVbeiQEXgCjphuI0j
L3Z/6zEcrTME+bjur9Sfg0rmWBZnZVtOCysorTGn4WWTgyA144Ox+TDYq+Jbm03e
+n5F7NQ9qr6J2oQeozee3sq7hcc7p7VcmOjYjpV1yYTz2vedTMwyFGXVSzCLsLmI
ku6apBR647Cx0tcaoCGP8MDNyEIeGREM07LezMBJvr9kVf9ABAVmghmlHJHILpV/
aVPYHcVqO5A9BlBdJsA3VMKU+O9oSQBR+x6vTzDVVR26jbk3f9zYf6lJB9DTFTFR
1aEqE/e14aigCKtGcZh7zNRlQlFwyPSBtFgMVzCr3tgJoTFcGaGe5l706iWrBpfS
SAHNDT8N9MG7yYlZOvecUyVVis7U0cmLhuipkOQ+AnHhak98Pp7JqXmCLlWjSm57
YFIGMmwpY0KhOoDBx/rkgKdFkGRBjwI57w==
=crmu
-----END PGP MESSAGE----- `

	empty = `-----BEGIN PGP MESSAGE-----

hQIMA5BMU610+Dk4AQ/9HvnmM79TPJUWLeBBw9fm+mMTlT61DEWomkgUQZxhjgRF
K18ifIkPuqyOj0fETj77kUnfVeJvsTDFjIjD30iAQ96J51Lka3PZPw6rQK+kyO16
rr6rhqKftucJwd3Z+HhO3fjAOuikjg9iD9hs4C2PxSLF9A6BuT7ajqKo9pXckioo
KQFOrQFd6XqqcaN2SNWK5NCNqCs+AQRR7FGDkKLaSH9j8mCg0+gpRN0u2gN6I/uw
kVK8/OaEti2nNGm3W8eF23Vi/OFwg/3HJ31wQzxZnYCAYXw8yjSNP0rrXmjcpf2P
vYIWjVw9qzuXMe0KFfMB4+TnL7IOdcyjsJzor+7rB9VrSzjRYdhwnHP3YyAf5PXy
8a6pc2F3EQFscY3A084j8KTfvBpyLdMBu3sKs+SKqYPEY9Ri15t4sJ7K6GIaZw7Y
RNYzEtgBuHIQPyHUGK8W3CQPXBch7DxJimE3GHjbAySKlgZs6C1reLoUzYr931Sc
hwXlez0oLsXg5cKWRnEfKZxDmrnU839vyzn1sEkDakrs1LeJtAR/fa631LCk68ca
GZoAIR7XLt7hrqL5swJImo42/EUasaofval/aAlVNHMcDmarxkkhktRq3rhoOrNH
TioqazVRNRvCWpZ5Pi3b9Mkby+BHvjn+vlcWZ696hw80xYt7OL3v1krmjBBDHaTS
PAHKGp0kNwtWskmxDWJr0QboHnoh3g2G4/h1yu7GlS8Lo/qFwtXDcMpmLbpR08+C
hDhXwNFBVCiWOfp32A==
=Hl8Y
-----END PGP MESSAGE-----`
)

func TestDecryptOfSomeText(t *testing.T) {
	expected := "Hello world!"
	result, err := gpg.Decrypt(helloWorld)
	if err != nil {
		t.Errorf("No errors were expected, got %s", err)
	}
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}

func TestDecryptOfEmpty(t *testing.T) {
	result, err := gpg.Decrypt(empty)
	if err != nil {
		t.Errorf("No errors were expected, got %s", err)
	}
	if result != "" {
		t.Errorf("Expected '', got '%s'", result)
	}
}

func TestDecryptFile(t *testing.T) {
	content := []byte(helloWorld)

	tmpfile, err := ioutil.TempFile("", "hello_world.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		panic(err)
	}

	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	expected := "Hello world!"
	reader, err := gpg.DecryptFile(tmpfile.Name())
	if err != nil {
		t.Errorf("No errors were expected, got %s", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	result := buf.String()
	if buf.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
