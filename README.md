# gdict
An offline CLI dictionary written in go, using data from wiktionary.

![screenshot.png](https://raw.githubusercontent.com/Lodobo/gdict/main/screenshot.png)

## Installation
#### Clone the Repository:
```bash
$ git clone https://github.com/Lodobo/gdict
$ cd gdict
```

#### build and install:
```bash
$ make install
```

gdict will be installed in `~/.local/bin` and the
sqlite database will be created in `~/.local/share/gdict`

## Dependencies
The Go compiler is required to build this program. Make sure you have go version 1.18 or newer.

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
- [dict.py](https://github.com/Lodobo/dict.py): The first version of this software written in python
-  BetaPictoris`s [dict](https://github.com/BetaPictoris/dict): another command line dictionary
- [wordnet](https://wordnet.princeton.edu/): A large lexical database of glosses and synonyms
