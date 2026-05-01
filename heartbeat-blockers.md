# Heartbeat blockers

- 2026-04-30 07:40 America/Tijuana: Medio Punto public rollout is blocked.
  - `medio-punto.com` still points to Namecheap parking, not the cluster.
  - Argo CD `medio-punto` is healthy, but it only deploys a placeholder `nginx:stable-alpine` service with no public ingress.
  - The only working route is the local preview ingress `medio-punto.local -> host.docker.internal:3101`.
  - The `statics` Helm chart revision currently used for this app is too minimal for a real production deploy, and its ingress template has an API version typo (`networking.k8s.io/1`).
  - Missing decisions/items: where the production app should be hosted in GitOps (containerized Next.js app vs static export), how the production ingress/TLS should be modeled, and DNS changes at the registrar.

- 2026-04-22 07:03 America/Tijuana: Inbox scan is blocked.
  - Email check via Himalaya failed because no config was found at `/Users/rfcku/Library/Application Support/himalaya/config.toml`.
  - Messages check via `imsg` failed because access to `/Users/rfcku/Library/Messages/chat.db` was denied.
  - Missing items to ask Peter about next time: Himalaya account setup, and Messages database / Full Disk Access permissions.
