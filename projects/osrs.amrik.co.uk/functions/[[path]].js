// Catch-all Pages Function. NOTE: in Pages, Functions run BEFORE static
// assets, so this must serve the deployed files itself (env.ASSETS) and only
// then fall back to proxying. Anything file-like that isn't in dist/ lives on
// gim-hub.com (the bundles reference the same content-hashed /hashed/* paths
// production serves, so 30k icon files aren't vendored here); clean paths get
// the SPA's index.html.
const UPSTREAM = "https://gim-hub.com";

export async function onRequest({ request, env }) {
  const url = new URL(request.url);

  if (!url.pathname.startsWith("/api/")) {
    const asset = await env.ASSETS.fetch(request);
    if (asset.status < 400) {
      return asset;
    }
  }

  const lastSegment = url.pathname.split("/").pop();
  const looksLikeFile = lastSegment.includes(".");

  if (url.pathname.startsWith("/api/") || looksLikeFile) {
    const upstream = new URL(url.pathname + url.search, UPSTREAM);
    return fetch(new Request(upstream, request));
  }

  return env.ASSETS.fetch(new URL("/index.html", url.origin));
}
