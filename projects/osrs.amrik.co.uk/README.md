# osrs.amrik.co.uk

PotatoBlood's year-in-review ("Wrapped") for gim-hub.com, served as a static
SPA with a Pages Function that proxies `/api/*` and asset requests to the
real gim-hub.com backend. Log in with the group name + token as usual; no
credentials are baked into the site.

`dist/` is a committed build artifact. To regenerate it, from the
`wrapped-poc` branch of [AmrikSD/gim-hub.com](https://github.com/AmrikSD/gim-hub.com):

```sh
STANDALONE=1 node_modules/.bin/vp build
cp -r dist/assets dist/index.html dist/favicon.png dist/apple-touch-icon.png \
  <this repo>/projects/osrs.amrik.co.uk/dist/
```

(Only those four entries — the heavyweight `dist/hashed/` icon/map assets are
deliberately not vendored; the function proxies them from gim-hub.com, which
serves identical content-hashed files.)

Deployed via Cloudflare Pages; infra in `infra/cloudflare/03-osrs.amrik.co.uk.tf`.
