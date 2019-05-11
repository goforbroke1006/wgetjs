# wgetjs
Download html and allow JS do work (smart wget).

Based upon [chromedp/chromedp](https://github.com/chromedp/chromedp)

### Arguments

| Name          | Description                                |
|---------------|--------------------------------------------|
| target-url    | Page address                               |
| ready-timeout | How many time you allow for JS affect page |
| output        | Save result to file                        |

### Dependencies

* Installed chrome browser

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
