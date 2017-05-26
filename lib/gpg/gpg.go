// Package gpg has utility methods for interacting with the gpg command line
// tool
package gpg

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

// Decrypt receives an input string and decrypts it using the gpg command
// line tool. It returns the decrypted string.
func Decrypt(input string) (string, error) {
	in := bytes.NewBufferString(input)
	out := new(bytes.Buffer)
	err := runGPGPipe([]string{"--decrypt", "--quiet"}, in, out)
	if err != nil {
		return "", err
	}

	if out.Len() > 0 {
		// Remove the 0xa that is present at the end of the buffer
		out.Truncate(out.Len() - 1)
	}

	return out.String(), nil
}

// DecryptFile receives a filename and decrypts it using the gpg command
// line tool. It returns the decrypted content as an io.Reader.
func DecryptFile(filename string) (io.Reader, error) {
	out := new(bytes.Buffer)
	err := runGPGPipe([]string{"--decrypt", "--quiet", filename}, nil, out)
	if err != nil {
		return out, err
	}

	if out.Len() > 0 {
		// Remove the 0xa that is present at the end of the buffer
		out.Truncate(out.Len() - 1)
	}

	return out, nil
}

func runGPGPipe(args []string, in io.Reader, out io.Writer) error {
	cmd := exec.Command("gpg", args...)
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
