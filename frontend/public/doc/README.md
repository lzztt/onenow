# Documentation as Code

Generate web pages and slides from `Markdown` docs.


## How it Works

```mermaid
flowchart LR
    A(Markdown) -->|marked| B(HTML)
    B -->|mermaid.js| C(Diagram)
    C --> D(Web Page)
    C -->|reveal.js| E(Slides)
```


## Usage

- Write a doc: `README.md`
- Web page: [`/doc/page.html?file=README`](/doc/page.html?file=README)
- Slides: [`/doc/slides.html?file=README`](/doc/slides.html?file=README)


## Files

```console
$ ls doc/
page.html  README.md  slides.html

$ wc -l doc/*.html
  57 doc/page.html
  63 doc/slides.html
 120 total
```


## Dependencies

- [`marked`](https://marked.js.org/): convert Markdown to a web page
- [`mermaid.js`](https://mermaid.js.org/): render Mermaid diagrams
- [`reveal.js`](https://revealjs.com/): turn the web page into slides
- [`bootstrap`](https://getbootstrap.com/) and [`github-markdown-css`](https://github.com/sindresorhus/github-markdown-css): web page styling
