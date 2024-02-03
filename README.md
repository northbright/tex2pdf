# tex2pdf
A Golang package to compile a tex to a PDF by running the xelatex command

## Requirements
* Install [TexLive](https://tug.org/texlive/)
  
  tex2pdf calls `xelatex` command which comes with installation of TexLive.
  Download and install [TexLive](https://tug.org/texlive/)("scheme-full" is recommended).

* Install [minted](https://www.ctan.org/pkg/minted) pacakge and [pygments](https://pygments.org) 

  [minted](https://www.ctan.org/pkg/minted) is used for code highlighting.
  Our test case includes compiling `.tex` file which uses minted for code highlighting("src/02-usage.tex").

  * Download and Install [pygments](https://pygments.org/download/) which is required by [minted](https://www.ctan.org/pkg/minted).
  * Install [minted](https://www.ctan.org/pkg/minted) package if need. TexLive Installation with "scheme-full" includes [minted](https://www.ctan.org/pkg/minted).

## Docs
* <https://pkg.go.dev/github.com/northbright/tex2pdf>

## Usage
```go
package main

import (
        "fmt"
        "log"
        "path/filepath"

        "github.com/northbright/tex2pdf"
)

func main() {
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
```
