# Heartbeat blockers

- 2026-05-01 07:14 America/Tijuana: Medio Punto public rollout is still blocked.
  - `medio-punto.com` still forwards through Namecheap (`302` to `www`) and apex HTTPS still times out.
  - `www.medio-punto.com` DNS still points at Namecheap parking (`parkingpage.namecheap.com` / `91.195.240.19`); the latest HTTPS probe timed out instead of returning the earlier `Parking/1.0` response, so public routing is still broken.
  - `medio-punto.local` still does not resolve normally on this Mac, but the ingress route works with `--resolve` and serves the current first-drop preview.
  - The local preview behind `127.0.0.1:3101` and `medio-punto.local` is now aligned again with the honest first-drop storefront (`Tote tejido hero`, `Roadmap del drop`), so the stale mock-route issue is no longer the main blocker.
  - The storefront rewrite is now committed locally in `github/rfcku/medio-punto` as `1fbef60` (`feat: reframe storefront as phase 1 drop preview`), but `main` is still `ahead 1` of `origin/main`, so it is not pushed upstream yet.
  - Argo CD `medio-punto` is healthy, but it still deploys only a placeholder `nginx:stable-alpine` service with no real public app ingress.
  - The `statics` Helm chart revision currently used for this app is too minimal for a real production deploy, and its ingress template has an API version typo (`networking.k8s.io/1`).
  - Missing decisions/items: push/preserve the storefront rewrite upstream, choose the temporary canonical operating route (`.local` or another preview), decide where the real production app should live in GitOps, and make the registrar/DNS changes.

- 2026-04-22 07:03 America/Tijuana: Inbox scan is blocked.
  - Email check via Himalaya failed because no config was found at `/Users/rfcku/Library/Application Support/himalaya/config.toml`.
  - Messages check via `imsg` failed because access to `/Users/rfcku/Library/Messages/chat.db` was denied.
  - Missing items to ask Peter about next time: Himalaya account setup, and Messages database / Full Disk Access permissions.
