package tex2pdf

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/northbright/copy/copydir"
	"github.com/northbright/copy/copyfile"
	"github.com/northbright/pathelper"
)

var (
	// Show xelatex output or not
	DebugMode = false

	ErrXelatexNotExist = errors.New("xelatex does not exists")
	ErrNoOutputPDF     = errors.New("xelatex compiled successfully but no output pdf found")
)

// Compile compiles a tex file into a PDF file by running xelatex.
func Compile(texFile, outputPDF string) error {
	// Check if xelatex command exists.
	if !pathelper.CommandExists("xelatex") {
		return ErrXelatexNotExist
	}

	// Get tex file's dir.
	srcDir := filepath.Dir(texFile)

	// Copy the source dir contains tex files to a temp dir.
	tmpDir := filepath.Join(os.TempDir(), filepath.Base(srcDir))
	if err := copydir.Do(context.Background(), srcDir, tmpDir); err != nil {
		return err
	}

	tmpTexFile := filepath.Join(tmpDir, filepath.Base(texFile))

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

		// Show xelatex output for DEBUG.
		if DebugMode {
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
		}

		// Run xelatex
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Get output PDF file path.
	baseFile := pathelper.BaseWithoutExt(texFile)
	pdf := filepath.Join(tmpDir, baseFile+".pdf")

	// Check if PDF exists.
	if !pathelper.FileExists(pdf) {
		return ErrNoOutputPDF
	}

	// Copy the PDF from temp dir to dst.
	if err := copyfile.Do(context.Background(), pdf, outputPDF); err != nil {
		return err
	}

	// Remove temp dir.
	return os.RemoveAll(tmpDir)
}
