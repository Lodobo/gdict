# gdict
An offline CLI dictionary written in go, using data from wiktionary.

![screenshot.png](https://raw.githubusercontent.com/Lodobo/gdict/main/screenshot.png)

## Installation
### From release:
```sh
# Download latest binaries
curl -LO https://github.com/Lodobo/gdict/releases/latest/download/gdict.AMD64
curl -LO https://github.com/Lodobo/gdict/releases/latest/download/install.AMD64
```
```sh
mkdir -p ~/.local/bin # create folder
install -m 755 gdict ~/.local/bin # install program to ~/.local/bin
./install # run the install script
```
### From source:
#### Clone the Repository:
```bash
$ git clone https://github.com/Lodobo/gdict
$ cd gdict
```
#### Build and install:
```bash
$ make install
```
gdict will be installed in `~/.local/bin` and the
sqlite database will be created in `~/.local/share/gdict`
### Note:
Make sure `~/.local/bin` is in $PATH. Additionaly, your system may not have a default pager. You may have to add the following lines to your `~/.bashrc` or `~/.profile`:
```bash
export PATH="$HOME/.local/bin:$PATH" # add ~/.local/bin to path
export PAGER="less -R" # Set less as the default pager
export LESSCHARSET=utf-8 # Enable unicode
```
## Dependencies
The Go compiler is required to build this program. Make sure you have go version 1.18 or newer. You can install it from [Go's official website](https://go.dev/doc/install)

## Usage of gdict:

|options|Description|
|----|----|
|-w [WORD]|search word in dictionary|
|-p [part of speech]|specify part of speech|
|-l [lang]|specify language|

## Available languages:

|ISO 639‑1|Full name|
|----|----|
|en|English|
|ar|Arabic|
|da|Danish|
|de|German|
|es|Spanish|
|fi|Finnish|
|fr|French|
|hi|Hindi|
|is|Icelandic|
|it|Italian|
|ja|Japanese|
|la|Latin|
|no|Norwegian|
|nb|Norwegian bokmål|
|nn|Norwegian nynorsk|
|nl|Dutch|
|pl|Polish|
|pt|Portuguese|
|ru|Russian|
|se|Northern sami|
|sv|Swedish|
|ur|Urdu|
|te|Telugu|
|zh|Chinese|

## See also
- Tatu Ylonen's [Wiktextract](https://github.com/tatuylonen/wiktextract): A utility for extracting data from wiktionary. The lexical data used by gdict comes from dumps provided by Ylonen on [kaikki.org](https://kaikki.org/) 
- [rdict](https://github.com/Lodobo/rdict): Another version of this software written in rust
-  BetaPictoris`s [dict](https://github.com/BetaPictoris/dict): another command line dictionary
- [wordnet](https://wordnet.princeton.edu/): A large lexical database of glosses and synonyms
