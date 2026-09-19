// Pages Function fallback: static assets in dist/ are served before this
// runs, so anything reaching here is either an API call, a fingerprinted
// asset that lives on gim-hub.com (the bundles reference the same
// content-hashed /hashed/* paths production serves, so we proxy rather than
// vendor 30k icon files), or an SPA route that needs index.html.
const UPSTREAM = "https://gim-hub.com";

export async function onRequest({ request, env }) {
  const url = new URL(request.url);
  const lastSegment = url.pathname.split("/").pop();
  const looksLikeFile = lastSegment.includes(".");

  if (url.pathname.startsWith("/api/") || looksLikeFile) {
    const upstream = new URL(url.pathname + url.search, UPSTREAM);
    return fetch(new Request(upstream, request));
  }

  return env.ASSETS.fetch(new URL("/index.html", url.origin));
}
