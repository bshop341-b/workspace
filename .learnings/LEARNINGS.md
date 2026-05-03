# Learnings

Corrections, insights, and knowledge gaps captured during development.

**Categories**: correction | insight | knowledge_gap | best_practice

---
## [LRN-20260426-003] correction

**Logged**: 2026-04-26T17:00:00-07:00
**Priority**: medium
**Status**: pending
**Area**: docs

### Summary
When a user reports an additional asset, confirm whether it is liquid/deployable before folding it into cash or retirement-readiness analysis.

### Details
I initially added the user's `238,518 MXN` Infonavit amount as an asset and let it affect the broad liquidity framing. The user clarified it is separate and non-liquid. Future finance updates should distinguish net-worth assets from investable/liquid assets immediately.

### Suggested Action
For personal finance sheets and analysis, label non-liquid assets explicitly and exclude them from liquid-assets commentary unless the user says they are available.

### Metadata
- Source: user_feedback
- Related Files: Projects/Rafa/Presupuesto 2026.md, Projects/Rafa/Perfil Rafa.md
- Tags: finance, liquidity, net-worth, correction

---
## [LRN-20260425-001] correction

**Logged**: 2026-04-25T22:18:04Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
GitHub repo mirrors for this workspace should live under ~/.openclaw/workspace/github, not ~/.openclaw/github.

### Details
The user corrected the expected mirror path after I cloned repos into ~/.openclaw/github. Future repo mirror/sync work for this workspace should target ~/.openclaw/workspace/github first and keep Obsidian notes aligned to that root.

### Suggested Action
Use ~/.openclaw/workspace/github as the canonical GitHub mirror root for this workspace and reflect that exact path in Obsidian GitHub notes.

### Metadata
- Source: user_feedback
- Related Files: USER.md, MEMORY.md
- Tags: github, workspace, obsidian, path

---

## [LRN-20260426-001] best_practice

**Logged**: 2026-04-26T18:07:25Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
For MongoDB Atlas cutovers, verify Atlas network allowlist access from inside the cluster before removing the old in-cluster database.

### Details
A direct cutover attempt caused new Ditto API pods to fail. Testing from inside the cluster isolated the issue to Atlas Network Access List restrictions rather than missing CA certificates or Ditto config.

### Suggested Action
Before deleting the old database application, test the Atlas URI from a pod in the target namespace and only proceed with GitOps removal after a successful ping.

### Metadata
- Source: error
- Related Files: github/argocd/applications/ditto/kustomization.yaml
- Tags: mongodb, atlas, cutover, argocd, ditto

---

## [LRN-20260426-002] best_practice

**Logged**: 2026-04-26T12:44:00-07:00
**Priority**: high
**Status**: pending
**Area**: frontend

### Summary
When a Next.js standalone UI image is already built with missing `NEXT_PUBLIC_*` client config and the source repo is unavailable, patching `server.js` plus the emitted static chunks at container startup is a workable GitOps fallback.

### Details
The deployed `ghcr.io/rfcku/vora-web:main` image had runtime env set correctly in Kubernetes, but the browser bundle still resolved API calls relatively because the standalone output carried `nextConfig.env = {}` and the client chunk created axios with an empty base URL. Since the real frontend source/build pipeline was unavailable and the upstream chart repo was not writable, the practical fix was to vendor the UI chart into the writable Argo repo and inject a startup patch that rewrites `/app/server.js`, the axios client chunk, and the media URL helper before `node /app/server.js` starts.

### Suggested Action
If the proper frontend source remains unavailable, prefer a startup patch on the deployed standalone artifact over risky ingress same-origin hacks, and vendor the chart into the writable GitOps repo when upstream chart permissions block the change.

### Metadata
- Source: error
- Related Files: github/argocd/applications/ditto/ditto-ui.yaml, github/argocd/charts/ditto/ui/templates/deployment.yaml, github/argocd/values/ditto/ui/values.yaml
- Tags: nextjs, standalone, argocd, ditto-ui, runtime-patch, gitops
- See Also: ERR-20260425-001, ERR-20260426-002

---
## [LRN-20260428-001] best_practice

**Logged**: 2026-04-28T14:48:30Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
For kind-based `.local` recovery, the fastest path is often ingress plus `ExternalName` services to `host.docker.internal` on fixed preview ports, with container images deferred until later.

### Details
The local cluster originally lacked namespaces, services, and ingress for `mytzuko.local` and `stitching-octopus.local`. Building fresh app images through Docker turned out to be slow and unreliable in this runtime, but both Next.js apps already built and served correctly on the host. Switching the cluster manifests to cert-manager certificates plus ingress rules backed by `ExternalName` services targeting `host.docker.internal:3102` and `host.docker.internal:3103` restored HTTPS routing immediately without waiting on registry pulls or kind image loading.

### Suggested Action
When the goal is to recover local-only hostnames quickly, prefer host-run previews behind ingress first. Use fixed ports, `ExternalName` services, and only circle back to fully containerized deployments once routing is proven and time pressure is lower.

### Metadata
- Source: error
- Related Files: github/argocd/workloads/local-sites/mytzuko/service.yaml, github/argocd/workloads/local-sites/stitching-octopus/service.yaml
- Tags: kind, ingress, externalname, nextjs, local-preview, cert-manager

---
## [LRN-20260428-002] correction

**Logged**: 2026-04-28T23:28:30Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
Having manifests in the Argo repo worktree is not the same as the sites being visible in Argo CD. They must also be represented by a tracked `Application` path in the remote repo.

### Details
I told the user the local-sites work was "in argocd/gitops" because the manifests existed under `github/argocd/workloads/local-sites` and had been applied manually with `kubectl apply -k`. The user correctly pointed out the sites were not visible in Argo CD. Inspection showed there was no `Application` for `workloads/local-sites`, and the whole folder plus `applications/ingress/local-sites.yaml` were still untracked locally, so Argo's remote repo view could not render or sync them.

### Suggested Action
When claiming something is in Argo/GitOps, verify both layers: the manifests exist in the repo path and an Argo `Application` or ApplicationSet actually points at that path in the remote tracked revision.

### Metadata
- Source: user_feedback
- Related Files: github/argocd/workloads/local-sites/kustomization.yaml, github/argocd/applications/ingress/local-sites.yaml
- Tags: argocd, gitops, correction, local-sites

---
## [LRN-20260501-001] correction

**Logged**: 2026-05-01T21:13:24Z
**Priority**: critical
**Status**: pending
**Area**: infra

### Summary
When a user asks to create a secret for GitOps, do not commit plaintext credentials to the repo without an explicit confirmation that secrets may live in Git history.

### Details
I created a Kubernetes secret definition from local environment values and committed it to the GitOps repository. The user then corrected that the secret should not have gone into the commit. The safer default is to keep secrets out of Git, prefer SealedSecrets/SOPS/ExternalSecrets, and if a temporary manual cluster secret is required, label it clearly as temporary.

### Suggested Action
Before committing any secret material, pause and confirm the storage method. If a secret was already pushed, immediately remove it from the live tree, restore service with a temporary out-of-band secret if needed, and recommend credential rotation plus optional history rewrite.

### Metadata
- Source: user_feedback
- Related Files: github/argocd/values/openclaw/cleo/values.yaml
- Tags: secrets, gitops, security, correction

---
## [LRN-20260502-002] best_practice

**Logged**: 2026-05-02T16:22:00Z
**Priority**: high
**Status**: pending
**Area**: tests

### Summary
When production config starts enforcing required auth/session secrets, Go unit tests should inject test-only defaults instead of requiring real env vars.

### Details
A Ditto worktree began failing broad `go test -race -coverprofile=coverage.out ./...` runs after `GetConfig()` started hard-failing on missing `JWT_SIGNING_SECRET` and `SESSION_COOKIE_SECRET`. The fix was to detect Go test binaries and apply non-production placeholder secrets before validation, so local and CI test runs stay hermetic while runtime validation remains strict for real server processes.

### Suggested Action
If config validation tightens in backend services, add explicit test-mode defaults or dedicated test env setup before enforcing secrets globally.

### Metadata
- Source: error
- Related Files: github/rfcku/ditto/config/config.go
- Tags: go, tests, config, env, ditto

---
## [LRN-20260502-001] best_practice

**Logged**: 2026-05-02T16:15:52Z
**Priority**: high
**Status**: pending
**Area**: frontend

### Summary
When the browser tool is policy-blocked for local `.local` URLs, a reliable fallback for local UI smoke tests is Playwright with the host Chrome executable in a temp npm sandbox.

### Details
The browser tool could not navigate to `https://ditto.local`, but local HTTP access still worked. Installing `@playwright/test` into a temp directory and launching Playwright against `/Applications/Google Chrome.app/Contents/MacOS/Google Chrome` made it possible to run a real browser smoke test anyway. That fallback immediately exposed the real issue in Ditto: the frontend bundle had `API URL: undefined` and was posting auth requests to same-origin frontend routes instead of `api.ditto.local/v1`.

### Suggested Action
For future local UI checks on this Mac, try Playwright plus the local Chrome binary before giving up on browser verification. Use that path especially when `.local` hostnames are reachable via curl but the first-class browser tool is blocked by policy.

### Metadata
- Source: error
- Related Files: /Users/rfcku/Documents/BishopVault/Projects/Ditto/Hallazgo UI auth local 2026-05-02.md
- Tags: playwright, chrome, browser, local-ui, ditto, smoke-test
- Pattern-Key: local-ui.playwright.chrome
- Recurrence-Count: 1
- First-Seen: 2026-05-02
- Last-Seen: 2026-05-02

---
## [LRN-20260502-003] best_practice

**Logged**: 2026-05-02T16:25:00Z
**Priority**: high
**Status**: pending
**Area**: frontend

### Summary
A runtime sed patch for Next.js env wiring must inspect compiled client chunks too, because browser bundles may reference `c.default.env.NEXT_PUBLIC_API_URL` instead of `process.env.NEXT_PUBLIC_API_URL`.

### Details
In the live Ditto UI pod, `server.js` and SSR chunks already carried `https://api.ditto.local/v1`, but the client chunk `.next/static/chunks/15e769aa6e2bc104.js` still contained `let tf=c.default.env.NEXT_PUBLIC_API_URL;`. That left the browser with `API URL: undefined` even though the container env var and server-side patch were present. The existing sed patch only replaced `process.env.NEXT_PUBLIC_API_URL`, so it missed the actual compiled client pattern.

### Suggested Action
When hot-patching a built Next.js app, inspect both `.next/static/chunks` and `.next/server/chunks` for the real compiled env access pattern before assuming a `process.env.*` replacement is sufficient.

### Metadata
- Source: error
- Related Files: github/argocd/values/ditto/ui/values.yaml, /Users/rfcku/Documents/BishopVault/Projects/Ditto/Hallazgo UI auth local 2026-05-02.md
- Tags: nextjs, runtime-patch, env, frontend, ditto
- Pattern-Key: nextjs.runtime_patch.client_env
- Recurrence-Count: 1
- First-Seen: 2026-05-02
- Last-Seen: 2026-05-02

---
