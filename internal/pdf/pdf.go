package pdf

import (
	"io"
	"os/exec"
)

func Generate(w io.Writer, r io.Reader) error {
	cmd := exec.Command("weasyprint", "-", "-")

	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	cmd.Start()

	// Write to stdin
	if _, err = io.Copy(in, r); err != nil {
		return err
	}

	// Close the stdin
	if err = in.Close(); err != nil {
		return err
	}

	// Copy stdout to writer
	if _, err = io.Copy(w, out); err != nil {
		return err
	}

	return nil
}
