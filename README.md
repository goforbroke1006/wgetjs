# wgetjs
Download html and allow JS do work (smart wget)

### Usage

```bash
./wgetjs \
    --target-url=https://www.some-site.com/some/page/what/modified-with-js-after-it-is-loaded/ \
    --ready-timeout=3500 \
    > ~/result.html
```

