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
