# tex2pdf
Package tex2pdf provides functions to compile Tex Files into a PDF via XeLaTex engine.

## Requirements
* Install [TexLive](https://tug.org/texlive/)
  
  tex2pdf calls `xelatex` command which comes with installation of TexLive.
  Download and install [TexLive](https://tug.org/texlive/)("scheme-full" is recommended).

* Install [minted](https://www.ctan.org/pkg/minted) pacakge and [pygments](https://pygments.org) 
  [minted](https://www.ctan.org/pkg/minted) is used for code highlighting.
  Our test case includes compiling `.tex` file which uses minted for code highlighting("src/02-usage.tex").

  * Download and Install [pygments](https://pygments.org/download/) which is required by [minted](https://www.ctan.org/pkg/minted).
  * TexLive Installation with "scheme-full" includes [minted](https://www.ctan.org/pkg/minted).

## Docs
* <https://pkg.go.dev/github.com/northbright/tex2pdf>
