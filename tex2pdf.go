package tex2pdf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/northbright/pathelper"
)

var (
	// Show xelatex output or not
	DebugMode = false
)

// Tex2PDF compiles a tex file into the PDF file by running xelatex.
// It outputs the pdf under the source tex file's dir and returns the compiled PDF path.
func Tex2PDF(texFile string) (string, error) {
	// Get absolute path of tex file.
	texFileAbsPath, err := filepath.Abs(texFile)
	if err != nil {
		return "", err
	}

	// Get source tex file's dir.
	srcDir := filepath.Dir(texFileAbsPath)

	// Check if xelatex command exists.
	if !pathelper.CommandExists("xelatex") {
		return "", fmt.Errorf("xelatex does not exists")
	}

	// Run "xelatex" command to compile a tex file into a PDF under src dir 2 times.
	// 1st time: create a PDF and .aux files(cross-references) and a .toc(Table of Content).
	// 2nd time: re-create the PDF with crosss-references and TOC.
	for i := 0; i < 2; i++ {
		// Run xelatex with options:
		// -synctex=1
		// -interaction=nonstopmode
		// -shell-escape
		cmd := exec.Command("xelatex", "-synctex", "1", "-interaction", "nonstopmode", "-shell-escape", texFileAbsPath)
		// Set work dir to source tex file's dir.
		cmd.Dir = srcDir

		// Show xelatex output for DEBUG.
		if DebugMode {
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
		}

		// Run xelatex
		if err := cmd.Run(); err != nil {
			return "", err
		}
	}

	// Get output PDF file path.
	baseFile := pathelper.BaseWithoutExt(texFile)
	pdf := filepath.Join(srcDir, baseFile+".pdf")

	// Check if PDF exists.
	if !pathelper.FileExists(pdf) {
		return "", fmt.Errorf("xelatex compiled successfully but no output pdf found")
	}

	return pdf, nil
}
