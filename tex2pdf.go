package tex2pdf

import (
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/northbright/cp"
	"github.com/northbright/pathelper"
)

var (
	ErrXelatexNotExist = errors.New("xelatex does not exists")
	ErrNoOutputPDF     = errors.New("xelatex compiled successfully but no output pdf found")
)

// Compiler reads main LaTex file and compiles all LaTex files to a PDF.
type Compiler struct {
	texFile   string
	outputPDF string
	stdout    io.Writer
	stderr    io.Writer
}

// Option represents the option of compiler.
type Option func(c *Compiler)

// Stdout returns option to set stdout of the cmd to run xelatex.
func Stdout(stdout io.Writer) Option {
	return func(c *Compiler) {
		c.stdout = stdout
	}
}

// Stderr returns option to set stderr of the cmd to run xelatex.
func Stderr(stderr io.Writer) Option {
	return func(c *Compiler) {
		c.stderr = stderr
	}
}

// New creates a new compiler.
func New(texFile, outputPDF string, options ...Option) *Compiler {
	c := &Compiler{
		texFile:   texFile,
		outputPDF: outputPDF,
	}

	for _, option := range options {
		option(c)
	}

	return c
}

// Compile compiles all LaTex files to a PDF.
func (c *Compiler) Compile() error {
	// Check if xelatex command exists.
	if !pathelper.CommandExists("xelatex") {
		return ErrXelatexNotExist
	}

	// Convert LaTex file and output PDF to absolute paths.
	texFile, err := filepath.Abs(c.texFile)
	if err != nil {
		return err
	}

	outputPDF, err := filepath.Abs(c.outputPDF)
	if err != nil {
		return err
	}

	// Get tex file's dir.
	srcDir := filepath.Dir(texFile)

	// Copy the source dir contains tex files to a temp dir.
	tmpDir := filepath.Join(os.TempDir(), filepath.Base(srcDir))
	_, err = cp.CopyDir(context.Background(), srcDir, tmpDir)
	if err != nil {
		return err
	}

	// Use base file name because we'll set work dir to the temp dir.
	tmpTexFile := filepath.Base(texFile)

	// Run "xelatex" command to compile a tex file into a PDF under temp dir 2 times.
	// 1st time: create a PDF and .aux files(cross-references) and a .toc(Table of Content).
	// 2nd time: re-create the PDF with crosss-references and TOC.
	for i := 0; i < 2; i++ {
		// Run xelatex with options:
		// -synctex=1
		// -interaction=nonstopmode
		// -8bit
		// -shell-escape
		cmd := exec.Command(
			"xelatex",
			"-synctex",
			"1",
			"-interaction",
			"nonstopmode",
			"-8bit",
			"-shell-escape",
			tmpTexFile,
		)

		// Set work dir to the temp dir.
		cmd.Dir = tmpDir

		// Set stdout and stderr for xelatex command.
		if c.stdout != nil {
			cmd.Stdout = c.stdout
		}
		if c.stderr != nil {
			cmd.Stderr = c.stderr
		}

		// Run xelatex
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Get output PDF file path.
	baseFile := pathelper.BaseWithoutExt(c.texFile)
	pdf := filepath.Join(tmpDir, baseFile+".pdf")

	// Check if PDF exists.
	if !pathelper.FileExists(pdf) {
		return ErrNoOutputPDF
	}

	// Copy the PDF from temp dir to dst.
	_, err = cp.CopyFile(context.Background(), pdf, outputPDF)
	if err != nil {
		return err
	}

	// Remove temp dir.
	return os.RemoveAll(tmpDir)
}
