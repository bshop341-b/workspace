# Heartbeat blockers

- 2026-05-04 10:37 America/Tijuana: Ditto frontend closure is still blocked.
  - The real product smoke now passes `register/login -> refresh -> create community -> create post -> vote`, and the post is visible on home plus its community.
  - `https://ditto.local/popular` and `https://ditto.local/trending` are no longer hard-broken because local exact redirects now return `308` to `https://ditto.local/`.
  - `https://ditto.local/explore` loads `Explore Communities | Vora`, so exploration is alive.
  - The remaining blocker is frontend ownership and rebuild path: there is still no canonical trending view, and the community page still throws CORS on `GET /v1/subs/check/:id?type=1`.
  - Missing items to ask or unblock next: identify the real source/pipeline for `ghcr.io/rfcku/vora-web:main`, rebuild the UI with a proper trending route instead of redirects, and fix the `subs/check` CORS wiring from a real frontend source.

- 2026-05-05 08:42 America/Tijuana: Medio Punto public rollout is still blocked.
  - `npm run build` passed again in `github/rfcku/medio-punto`, and the local preview on `127.0.0.1:3101` is back to `HTTP 200`.
  - `https://medio-punto.local/` is healthy again when smoke-tested with `curl -k --resolve medio-punto.local:443:127.0.0.1`, so the host-run preview path is working.
  - `medio-punto.com` still forwards through Namecheap (`302` to `www`) and apex HTTPS still times out.
  - `www.medio-punto.com` still points at Namecheap parking and is not serving the app publicly.
  - The storefront rewrite is already synced upstream at `1fbef60` (`feat: reframe storefront as phase 1 drop preview`), so the repo state is no longer the blocker.
  - Argo CD `medio-punto` still deploys the placeholder `charts/statics` app with empty values (`values/statics/medio-punto/values.yaml` is `{}`), while the only working ingress in GitOps is the local-only `medio-punto.local` route in `workloads/local-sites/medio-punto`.
  - `github/rfcku/medio-punto` still has no `Dockerfile` or deployment workflow, so there is no production image path for GitOps to deploy yet.
  - Even the fallback `charts/statics` route needs chart work before it can host a public site: `charts/statics/templates/ingress.yaml` still declares the invalid API version `networking.k8s.io/1`.
  - Missing decisions/items: pick the real production hosting/GitOps path for the Next.js app, add the build/deploy path, and make the registrar/DNS changes for `medio-punto.com`.

- 2026-04-22 07:03 America/Tijuana: Inbox scan is blocked.
  - Email check via Himalaya failed because no config was found at `/Users/rfcku/Library/Application Support/himalaya/config.toml`.
  - Messages check via `imsg` failed because access to `/Users/rfcku/Library/Messages/chat.db` was denied.
  - Missing items to ask Peter about next time: Himalaya account setup, and Messages database / Full Disk Access permissions.
