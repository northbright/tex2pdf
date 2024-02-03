package tex2pdf_test

import (
	"fmt"
	"log"

	"github.com/northbright/tex2pdf"
)

func ExampleTex2PDF() {
	// Use DEBUG mode.
	tex2pdf.DebugMode = true

	texFile := "src/my_book.tex"

	// Compile a tex file to PDF.
	pdf, err := tex2pdf.Tex2PDF(texFile)
	if err != nil {
		log.Printf("Tex2PDF() error: %v", err)
		return
	}

	fmt.Printf("Tex2PDF OK, output pdf: %v\n", pdf)

	// Output:
	//Tex2PDF OK, output pdf: src/my_book.pdf
}
