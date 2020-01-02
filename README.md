# Curify

Curify is a set of command line tools to extract the main readable content from HTML pages and convert them into markdown.

It based on the work of [go-shiori/go-readability](https://github.com/go-shiori/go-readability) and [lunny/html2md](https://gitea.com/lunny/html2md).

## Curify

Curify extract the main readable content from HTML pages

### Installation

    $ go get github.com/squioc/curify/cmd/curify/

### Usage

    $ curify article.html

or

    $ curl https://example.org | curify


## Tomd

### Installation

    $ go get github.com/squioc/curify/cmd/tomd/

### Usage

    $ tomd curified.html

or

    $ cat article.html | tomd
