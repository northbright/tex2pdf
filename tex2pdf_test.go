package tex2pdf_test

import (
	"log"

	"github.com/northbright/tex2pdf"
)

func ExampleCompiler_Compile() {
	texFile := "example/src/my_book.tex"
	outputPDF := "output/my_book.pdf"

	// Create a compiler with specified stdout and stderr options.
	// c := tex2pdf.New(texFile, outputPDF, tex2pdf.Stdout(os.Stdout), tex2pdf.Stderr(os.Stderr))
	// Create a compiler.
	c := tex2pdf.New(texFile, outputPDF)

	log.Printf("start compiling Tex File to PDF...\nTex file: %v\noutput PDF: %v", texFile, outputPDF)
	// Compile a Tex file to a PDF.
	if err := c.Compile(); err != nil {
		log.Printf("Compile() error: %v", err)
		return
	}

	log.Printf("compile successfully")

	// Output:
}
