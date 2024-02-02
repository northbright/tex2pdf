package tex2pdf

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/northbright/iocopy"
	"github.com/northbright/pathelper"
)

// Tex2PDF compiles the Tex file into the PDF file by running xelatex.
func Tex2PDF(texFile, pdfFile string) error {
	// If pdfFile is empty, set it to the same name as Tex file and put it to current dir.
	if pdfFile == "" {
		pdfFile = pathelper.BaseWithoutExt(texFile) + ".pdf"
	}

	// Get absolute path of tex file.
	texFileAbsPath, err := filepath.Abs(texFile)
	if err != nil {
		return err
	}

	// Get Source Tex File's Dir.
	srcDir := filepath.Dir(texFileAbsPath)

	// Check if xelatex command exists.
	if !pathelper.CommandExists("xelatex") {
		return fmt.Errorf("xelatex does not exists")
	}

	// Run "xelatex" command to compile Tex file into a PDF under src dir 2 times.
	// 1st time: create a PDF and .aux files(cross-references) and a .toc(Table of Content).
	// 2nd time: re-create the PDF with crosss-references and TOC.
	for i := 0; i < 2; i++ {
		cmd := exec.Command("xelatex", "-shell-escape", texFileAbsPath)
		// Set work dir to tex file's dir.
		cmd.Dir = srcDir

		// Show xelatex output for DEBUG.
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin

		// Run xelatex
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Source PDF: TMPDIR/TEX_FILE_NAME.pdf
	src := filepath.Join(srcDir, pathelper.BaseWithoutExt(texFile)+".pdf")
	dst := pdfFile

	ctx := context.Background()
	bufSize := (int64)(64 * 1024)

	// Copy PDF from src to dst.
	if _, err := iocopy.CopyFile(ctx, dst, src, bufSize); err != nil {
		return err
	}

	return nil
}
