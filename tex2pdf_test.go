package tex2pdf_test

import (
	"log"

	"github.com/northbright/tex2pdf"
)

func ExampleCompile() {
	// Open DEBUG mode if need.
	//tex2pdf.DebugMode = true

	texFile := "example/src/my_book.tex"
	outputPDF := "output/my_book.pdf"

	log.Printf("start to compile Tex File to PDF...\nTex file: %v\noutput PDF: %v", texFile, outputPDF)
	// Compile a Tex file to a PDF.
	if err := tex2pdf.Compile(texFile, outputPDF); err != nil {
		log.Printf("Compile() error: %v", err)
		return
	}

	log.Printf("compile successfully")

	// Output:
}
