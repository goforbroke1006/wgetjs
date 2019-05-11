# wgetjs
Download html and allow JS do work (smart wget)

### Usage

```bash
./wgetjs \
    --target-url=https://www.some-site.com/some/page/what/modified-with-js-after-it-is-loaded/ \
    --ready-timeout=3500 \
    > ~/result.html
```
### Install

```bash
go get -u github.com/goforbroke1006/wgetjs/cmd/wgetjs
```
