package tex2pdf_test

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/northbright/tex2pdf"
)

func ExampleTex2PDF() {
	// Open DEBUG mode if need.
	//tex2pdf.DebugMode = true

	texFile := "src/my_book.tex"

	// Compile a tex file to PDF.
	pdf, err := tex2pdf.Tex2PDF(texFile)
	if err != nil {
		log.Printf("Tex2PDF() error: %v", err)
		return
	}

	fmt.Printf("Tex2PDF OK, output pdf: %v\n", filepath.Base(pdf))

	// Output:
	//Tex2PDF OK, output pdf: my_book.pdf
}
