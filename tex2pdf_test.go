package tex2pdf_test

import (
	"log"

	"github.com/northbright/tex2pdf"
)

func ExampleTex2PDF() {
	texFile := "src/my_book.tex"
	pdfFile := ""
	if err := tex2pdf.Tex2PDF(texFile, pdfFile); err != nil {
		log.Printf("Tex2PDF() error: %v", err)
		return
	}

	// Output:
}
