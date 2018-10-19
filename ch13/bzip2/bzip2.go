package bzip2

import (
	"io"
	"log"
	"os/exec"
)

type writer struct {
	w      io.Writer // underlying output stream
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2.exe")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	w := &writer{w: out, cmd: cmd, stdin: stdin, stdout: stdout, stderr: stderr}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	return w.stdin.Write(data)
}

// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	if err := w.stdin.Close(); err != nil {
		return err
	}
	if _, err := io.Copy(w.w, w.stdout); err != nil {
		return err
	}
	if err := w.cmd.Wait(); err != nil {
		return err
	}
	return nil
}
