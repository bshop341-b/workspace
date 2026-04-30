# Errors

Command failures and integration errors.

---

## [ERR-20260428-001] background_exec_sigterm_hosts

**Logged**: 2026-04-28T16:00:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
An async background exec against `nctto.local` and `mytzuko.local` terminated with `SIGTERM` before producing actionable output.

### Error
```
Process exited with signal SIGTERM.
```

### Context
- Operation attempted: previously launched async exec session `crisp-sable`
- Visible log output only showed host headers: `## nctto.local` and `## mytzuko.local`
- The command summary in the process list was truncated to `for host`, so the exact original command was not recoverable from the surfaced event alone
- No user-facing follow-up was sent because the completion event explicitly said to handle the result internally

### Suggested Fix
When launching long-running or multi-host background execs, emit the full command and enough structured logging to diagnose early termination, and check whether the process was intentionally canceled, timed out, or killed by the host.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md

---

## [ERR-20260427-001] git-push

**Logged**: 2026-04-27T01:28:43Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
Pushing `rfcku/ditto` from this session failed because GitHub access is authenticated as `bshop341-b`, which does not have write access to `rfcku/ditto`.

### Error
```
remote: Permission to rfcku/ditto.git denied to bshop341-b.
fatal: unable to access 'https://github.com/rfcku/ditto.git/': The requested URL returned error: 403
```

### Context
- Command attempted: `git -C /Users/rfcku/.openclaw/workspace/github/rfcku/ditto push origin main`
- Local branch `main` is ahead of `origin/main` and contains commit `6c65e0c`
- The session can read local mirrors but cannot push to the upstream repo with the current GitHub identity

### Suggested Fix
Use a GitHub auth context with write access to `rfcku/ditto`, or retarget the remote to an account that owns the deployment pipeline.

### Metadata
- Reproducible: yes
- Related Files: github/rfcku/ditto

---

## [ERR-20260423-001] cron_whatsapp_delivery_target

**Logged**: 2026-04-23T08:05:00-07:00
**Priority**: high
**Status**: resolved
**Area**: config

### Summary
Morning and night report cron jobs generated successfully but failed to deliver because WhatsApp cron delivery used `to: default` instead of an E.164 target.

### Error
```
Delivering to WhatsApp requires target <E.164|group JID>
```

### Context
- Operation attempted: cron announce delivery for morning and night reports
- Jobs affected: `reporte matutino ejecutivo`, `reporte nocturno ejecutivo`
- Report generation and Obsidian save succeeded; only final WhatsApp delivery failed
- Corrective action taken: updated both cron jobs to use `+5216642800707`

### Suggested Fix
When configuring cron delivery for WhatsApp, always set `delivery.to` to a real E.164 phone number or group JID, not `default`.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

### Resolution
- **Resolved**: 2026-04-23T08:05:30-07:00
- **Notes**: Updated both WhatsApp cron jobs to use the direct E.164 target `+5216642800707` instead of `default`.

---
## [ERR-20260426-004] gog_sheets_find_replace_flag

**Logged**: 2026-04-26T16:04:00-07:00
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
Attempted to use a non-existent `--all-sheets` flag with `gog sheets find-replace`.

### Error
```
unknown flag --all-sheets
Run with --help to see available flags
```

### Context
- Operation attempted: rename `Gas` to `Gasolina` across a spreadsheet
- Assumed `find-replace` needed an explicit all-sheets flag
- The command failed before any sheet changes were made

### Suggested Fix
Check `gog sheets find-replace --help` first. If no tab-scoping flag exists, treat the command as spreadsheet-wide by default or run it per target tab.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260426-007] gog_sheets_range_column_mismatch

**Logged**: 2026-04-26T16:58:00-07:00
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
A `gog sheets update` failed because the target range did not include enough columns for the provided row values.

### Error
```
Google API error (400 badRequest): Requested writing within range [BALANCE!A13:D15], but tried writing to column [E]
```

### Context
- Operation attempted: add a new Infonavit asset row and refresh totals in `BALANCE`
- The payload for `Total Neto` contained 6 cells, but the range only allowed 4 columns
- Earlier row insertion likely succeeded; only the write failed

### Suggested Fix
When rows have different widths, either widen the range to fit the longest row or write narrower blocks separately.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260426-005

---
## [ERR-20260426-006] gog_sheets_merge_rows_flag

**Logged**: 2026-04-26T16:30:00-07:00
**Priority**: low
**Status**: resolved
**Area**: docs

### Summary
Tried to merge the annual header cells in Google Sheets using a non-existent `--rows` flag.

### Error
```
unknown flag --rows
Run with --help to see available flags
```

### Context
- Operation attempted: merge `GASTOS!H1:I1` after adding annual columns
- The sheet content update succeeded; only the cosmetic merge step failed
- Need to check the exact `gog sheets merge --help` syntax before using orientation flags

### Suggested Fix
Consult `gog sheets merge --help` and only pass supported flags, or skip merging when the header still reads clearly without it.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260426-004

### Resolution
- **Resolved**: 2026-04-26T16:31:00-07:00
- **Notes**: Confirmed `gog sheets merge` uses `--type` (`MERGE_ALL|MERGE_COLUMNS|MERGE_ROWS`), then merged `GASTOS!H1:I1` successfully with `--type MERGE_ALL`.

---
## [ERR-20260426-005] gog_sheets_values_json_quoting

**Logged**: 2026-04-26T16:10:00-07:00
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
A multi-range `gog sheets update` bootstrap script failed because the JSON payload for `--values-json` broke when formulas contained embedded quotes.

### Error
```
invalid JSON values: invalid character '"' after array element
```

### Context
- Operation attempted: seed a new Google Sheet for Casa 373 with formulas and dashboard rows
- The issue came from hand-built JSON inside a shell heredoc, especially formulas using `MATCH("Netflix", ...)`
- The sheet creation likely succeeded before the update failure

### Suggested Fix
Generate the values payload with Python `json.dumps()` or avoid embedded double-quote formulas inside raw shell JSON.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260426-004

---
## [ERR-20260426-003] github_repo_lookup_ditto_ui

**Logged**: 2026-04-26T14:15:00-07:00
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Trying to clone the expected frontend repos `rfcku/ditto-ui` and `rfcku/vora-web` failed because neither repository was accessible by those names, even though the deployed UI image clearly identifies itself as `vora-web`.

### Error
```
remote: Repository not found.
fatal: repository 'https://github.com/rfcku/ditto-ui.git/' not found
remote: Repository not found.
fatal: repository 'https://github.com/rfcku/vora-web.git/' not found
```

### Context
- Operation attempted: locate and clone the Ditto frontend source after confirming it was missing from `~/.openclaw/workspace/github/rfcku`
- `gh repo list rfcku` with current auth showed no accessible repo matching `ditto-ui` or `vora-web`
- The running pod `/app/package.json` identifies the deployed app as `vora-web`, so the source likely lives under a different owner/name or is not accessible from this environment

### Suggested Fix
Before assuming the repo name from a deployment or package name, confirm the canonical source repository via image labels, CI metadata, or an explicit repo URL from the user.

### Metadata
- Reproducible: yes
- Related Files: github/argocd/applications/ditto/ditto-ui.yaml, github/argocd/charts/ditto/ui/values.yaml
- See Also: ERR-20260425-001

---
## [ERR-20260425-001] github-push-permissions

**Logged**: 2026-04-25T19:00:00Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
Pushing changes to rfcku/helm-charts failed because the configured GitHub identity only had write access to bshop341-b/argocd.

### Error
```
remote: Permission to rfcku/helm-charts.git denied to bshop341-b.
fatal: unable to access 'https://github.com/rfcku/helm-charts.git/': The requested URL returned error: 403
```

### Context
- Operation attempted: push chart changes adding imagePullSecrets support to rfcku/helm-charts
- The same session could push successfully to bshop341-b/argocd
- This indicates cross-repo permission mismatch, not a generic git or network failure

### Suggested Fix
Check which GitHub account/token `gh` is using before attempting repo writes, and verify it has write access to the destination repository.

### Metadata
- Reproducible: yes
- Related Files: charts/ditto/charts/apis/templates/deployment.yaml, charts/ditto/charts/apis/values.yaml

---

## [ERR-20260426-001] mongodb_atlas_cutover

**Logged**: 2026-04-26T18:07:25Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
Switching Ditto from local MongoDB to Atlas failed because Atlas rejected cluster connections during TLS handshake due to Network Access List restrictions.

### Error
```
server selection timeout, current topology: ReplicaSetNoPrimary ... remote error: tls: internal error
MongoServerSelectionError ... Please ensure that your Network Access List allows connections from your IP.
```

### Context
- Operation attempted: patch `app-secrets` MONGO_URI to Atlas and restart Ditto API deployments
- Cluster test from inside `ditto-mongodb` with `mongosh` hit the same Atlas rejection
- The local MongoDB URI was restored to keep Ditto healthy

### Suggested Fix
Allow the cluster egress/public IP in MongoDB Atlas before cutting Ditto over from the in-cluster MongoDB service.

### Metadata
- Reproducible: yes
- Related Files: github/argocd/applications/ditto/kustomization.yaml, github/argocd/values/ditto/apis/values.yaml

---

## [ERR-20260426-002] shell_sed_inplace_on_nonwritable_dir

**Logged**: 2026-04-26T12:44:00-07:00
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Using `sed -i` inside the `ditto-ui` container failed even though the target file was writable, because the parent directory was not writable for temporary file creation.

### Error
```
sed: can't create temp file 'server.jsXXXXXX': Permission denied
```

### Context
- Operation attempted: patch `/app/server.js` at container startup to inject `NEXT_PUBLIC_API_URL`
- The container runs as `nextjs`, the file was writable, but `/app` itself was not writable by that user
- `sed -i` needs directory write access to create a temporary file before replacing the original

### Suggested Fix
When patching files in containers with read-only or non-writable directories, write to `mktemp`, then overwrite the existing file with `cat > target` instead of relying on `sed -i`.

### Metadata
- Reproducible: yes
- Related Files: github/argocd/values/ditto/ui/values.yaml
- See Also: ERR-20260425-001

---

## [ERR-20260427-002] vercel-cli-auth

**Logged**: 2026-04-27T16:39:00Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
A Vercel CLI attempt failed because this environment is not logged into Vercel and no token was provided.

### Error
```
Error: No existing credentials found. Please run `vercel login` or pass "--token"
Learn More: https://err.sh/vercel/no-credentials-found
```

### Context
- Command attempted: `npx vercel`
- The CLI installed and started normally, then exited after checking credentials
- `npm` also warned about unsupported engines under Node `v25.9.0`, but the hard blocker was missing Vercel auth

### Suggested Fix
Authenticate Vercel in this environment with `vercel login` or provide a valid `--token` before retrying deployment or project inspection commands.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260427-001] musicbrainz_ingest_connectivity

**Logged**: 2026-04-27T21:40:00Z
**Priority**: medium
**Status**: pending
**Area**: backend

### Summary
MusicBrainz batch ingestion failed on first live fetch with TLS connection reset.

### Error
```
TypeError: fetch failed
cause: Client network socket disconnected before secure TLS connection was established
code: ECONNRESET
host: musicbrainz.org
port: 443
```

### Context
- Command attempted: `INGEST_TARGET_COUNT=10 MUSICBRAINZ_USER_AGENT='metal-api/0.1.0 (openclaw@example.com)' npm run ingest:musicbrainz`
- Repo: `~/.openclaw/workspace/github/rfcku/metal-api`
- The new ingestion script wrote failure metadata but could not complete a live batch.

### Suggested Fix
Retry with a confirmed User-Agent and add retry/backoff around MusicBrainz fetches, or verify outbound connectivity/TLS from this host before depending on the script for the first real dataset run.

### Metadata
- Reproducible: unknown
- Related Files: github/rfcku/metal-api/scripts/ingest-musicbrainz.mjs

---
## [ERR-20260428-001] sed

**Logged**: 2026-04-28T14:09:00Z
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
Used an invalid sed substitution while listing repos under the workspace.

### Error
```text
sed: 1: "s#^#/##": bad flag in substitute command: '#'
```

### Context
- Command attempted: `find /Users/rfcku/.openclaw/workspace/github -maxdepth 2 -mindepth 2 -type d -name .git -prune -exec dirname {} \\; | sed 's#^#/##' | sed 's#^/Users/rfcku/.openclaw/workspace/github/##' | sort`
- The first `sed` expression was malformed and unnecessary.

### Suggested Fix
Remove the bad `sed` step and trim the prefix with a single valid substitution.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260428-002] docker-buildx

**Logged**: 2026-04-28T14:48:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Attempting to containerize the Mytzuko and Stitching Octopus Next.js apps via Docker hung silently long enough that the async exec sessions were later killed by the harness.

### Error
```text
Exec failed (calm-roo, signal SIGTERM) :: HOSTS CHECK ## mytzuko.local ## stitching-octopus.local
Exec failed (ember-ha, signal SIGKILL)
```

### Context
- Operations attempted: `docker build -t local/mytzuko:dev .` and `docker build -t local/stitching-octopus:dev .` after adding Dockerfiles for both repos
- A direct `docker pull node:22-alpine` also produced no progress output for an extended period in this runtime
- The actual goal was to restore `.local` routes on the existing kind cluster, not necessarily to finish image packaging first
- A faster workaround succeeded: keep the Next.js apps running on fixed host ports and route cluster ingress to `host.docker.internal` with `ExternalName` services

### Suggested Fix
When local `.local` recovery is urgent, probe Docker feasibility quickly. If image pulls or builds stall, pivot to host-backed previews through ingress instead of waiting on long silent Docker operations.

### Metadata
- Reproducible: unknown
- Related Files: github/nctto/mytzuko/Dockerfile, github/rfcku/stitchingoctopus/Dockerfile, github/argocd/workloads/local-sites

---
## [ERR-20260428-003] exec-event-elevation

**Logged**: 2026-04-28T15:12:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
An exec-event heartbeat session could not perform the final `/etc/hosts` edit because elevated exec is disabled for this provider.

### Error
```text
elevated is not available right now (runtime=direct).
Failing gates: allowFrom (tools.elevated.allowFrom.<provider> / agents.list[].tools.elevated.allowFrom.<provider>)
Context: provider=exec-event session=agent:main:cron:296ab27d-a8a5-4d0f-a693-2df1a8369b0a:heartbeat
```

### Context
- Operation attempted: append `mytzuko.local` and `stitching-octopus.local` to the existing `172.18.255.200 ...` line in `/etc/hosts`
- Verification before the failure showed both hosts already return `HTTP/2 200` when resolved to `172.18.255.200`
- The remaining blocker was only the privileged hosts-file edit, not cluster, TLS, or app readiness

### Suggested Fix
For exec-event or heartbeat runs, do not assume sudo/host-file edits are available. Either hand off the exact manual command to the user, switch to a runtime that supports elevation, or defer the final privileged step to a human-run shell.

### Metadata
- Reproducible: yes
- Related Files: /etc/hosts, HEARTBEAT.md, .learnings/ERRORS.md

---
## [ERR-20260428-004] recursive-grep-sigkill

**Logged**: 2026-04-28T15:16:40Z
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
A broad `grep -R` over the workspace plus BishopVault was killed by the harness before it returned any matches.

### Error
```text
Exec failed (nimble-p, signal SIGKILL)
```

### Context
- Command attempted: `grep -R "6642800707\|whatsapp main\|main channel" -n /Users/rfcku/.openclaw/workspace /Users/rfcku/Documents/BishopVault 2>/dev/null | head -100`
- The search was meant to locate the WhatsApp main-channel reference and related blocker notes
- A follow-up `rg -n --hidden -S` with targeted excludes returned the needed matches immediately, including the blocker report and task notes

### Suggested Fix
Prefer `rg` with explicit exclude globs for wide vault/workspace searches. Avoid `grep -R` over both trees in heartbeat or exec-event runs because it can be slow enough to get killed before producing output.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md, HEARTBEAT.md, /Users/rfcku/Documents/BishopVault/Reports/Cron seguimiento bloqueos 2026-04-28.md

---
## [ERR-20260428-005] atlas_tls_handshake_rejected

**Logged**: 2026-04-28T15:55:00Z
**Priority**: high
**Status**: pending
**Area**: infra

### Summary
Trying to use the Atlas URI from `~/.openclaw/.env` for the kanban backend failed because Atlas rejected the TLS handshake before the driver could select a server.

### Error
```text
MongoServerSelectionError: ... tlsv1 alert internal error ... SSL alert number 80
```

### Context
- Operation attempted: verify and use the `MONGODB_URI` from `~/.openclaw/.env` with `kanban-dashboard-backend`
- Direct Node driver connection and `openssl s_client` to the Atlas shard host both failed with the same TLS alert
- Current public IP during the test was `189.223.36.251`
- This pattern matches Atlas access-list or cluster-side network rejection more than an app-code issue

### Suggested Fix
Check the Atlas cluster status and Network Access allowlist, then add the current public IP (or the correct egress IP/range) before retrying the backend connection.

### Metadata
- Reproducible: yes
- Related Files: /Users/rfcku/.openclaw/.env, github/bshop341-b/kanban-dashboard-backend/.env.example, .learnings/ERRORS.md

---
## [ERR-20260428-006] lumenforge_local_start_authjs_mongo

**Logged**: 2026-04-28T16:26:30Z
**Priority**: high
**Status**: resolved
**Area**: config

### Summary
A production-mode local start of `lumenforge` failed because Auth.js rejected `https://nctto.local` as an untrusted host, while the app was also pointed at a non-running local MongoDB instance.

### Error
```text
[auth][error] UntrustedHost: Host must be trusted. URL was: https://nctto.local/api/auth/session
MongoServerSelectionError: connect ECONNREFUSED ::1:27017, connect ECONNREFUSED 127.0.0.1:27017
```

### Context
- Commands involved: `npm run build` and `MONGODB_URI=mongodb://localhost:27017/test npm run start -- --hostname 0.0.0.0 --port 3104`
- Repo: `github/bshop341-b/lumenforge`
- `next build` completed, but page data collection and runtime requests emitted repeated Mongo connection failures
- `next start` then surfaced the Auth.js host-trust error when the app served `nctto.local` in production mode without `AUTH_URL`, `AUTH_TRUST_HOST`, or explicit `trustHost: true`
- There was no `.env` file in the repo root, only `.env.example`

### Suggested Fix
Set Auth.js host trust explicitly for local/preview environments (`trustHost: true` or `AUTH_TRUST_HOST=true` / `AUTH_URL=https://nctto.local`) and provide a reachable MongoDB instance before relying on authenticated routes or session polling.

### Metadata
- Reproducible: yes
- Related Files: github/bshop341-b/lumenforge/src/auth.ts, github/bshop341-b/lumenforge/src/lib/mongodb.ts, github/bshop341-b/lumenforge/.env.example

### Resolution
- **Resolved**: 2026-04-28T16:33:40Z
- **Notes**: Patched `src/auth.ts` to pass `secret: env.AUTH_SECRET` and `trustHost: true`, documented `AUTH_URL` / `AUTH_TRUST_HOST` in `.env.example` and `README.md`, rebuilt the app, and verified `GET /api/auth/session` on `Host: nctto.local` now returns `200 null` instead of the Auth.js configuration error. Mongo is still required for real sign-in and DB-backed flows.

---
## [ERR-20260428-001] exec-elevated-direct-runtime

**Logged**: 2026-04-28T17:28:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Elevated exec is unavailable in exec-event direct runtime, blocking /etc/hosts updates from heartbeat tasks.

### Error
```
elevated is not available right now (runtime=direct).
Failing gates: allowFrom (tools.elevated.allowFrom.<provider> / agents.list[].tools.elevated.allowFrom.<provider>)
Context: provider=exec-event session=agent:main:main:heartbeat
```

### Context
- Command attempted: `sudo ./scripts/add-local-hosts.sh 172.18.255.200 nctto.local mytzuko.local stitching-octopus.local`
- Goal: unblock local browser verification for three `.local` sites already live behind ingress.
- Environment: OpenClaw exec-event heartbeat on main session.

### Suggested Fix
Use a runtime/provider that permits elevated exec for exec-event heartbeats, or route host-file changes through a user-approved/manual path.

### Metadata
- Reproducible: yes
- Related Files: github/argocd/scripts/add-local-hosts.sh, HEARTBEAT.md

---
## [ERR-20260429-002] obsidian-cli-search-path-only

**Logged**: 2026-04-29T16:04:30Z
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
Tried to use a non-existent `--path-only` flag with `obsidian-cli search`.

### Error
```
Error: unknown flag: --path-only
```

### Context
- Command attempted: `obsidian-cli search "Medio Punto" --path-only`
- Goal: quickly list matching note paths inside the active vault
- The CLI only supports plain `search`, plus `--editor` and `--vault`

### Suggested Fix
Use `find` or read the vault directly when exact note paths are needed, and check `obsidian-cli search --help` before assuming path-output flags exist.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260429-001] kubectl-ditto-react-ssr-log

**Logged**: 2026-04-29T14:27:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Background `kubectl ditto` exec labeled `ember-sable` exited with SIGTERM while dumping a large minified React DOM server stack, which obscured the actionable app error.

### Error
```
Exec failed (ember-sable, signal SIGTERM)
/app/node_modules/react-dom/cjs/react-dom-server.browser.production.js:2019
? resumableState.styleResources[href$jscomp$0]
Process exited with signal SIGTERM.
```

### Context
- Background process label/session: `ember-sable`
- Command family shown by process list: `kubectl ditto`
- Captured output was dominated by minified `.next` / `react-dom-server.browser.production.js` frames around style resource handling, making the originating application error unclear.
- Follow-up inspection worked better with `process poll` on the process label than with recursive `rg`, which exploded on bundled assets.

### Suggested Fix
Rerun the failing `kubectl ditto` workflow with narrower log capture, ideally the originating request/error lines before the React stack flood, and prefer source-mapped or debug-prerender output if this is a Next.js prerender/SSR failure.

### Metadata
- Reproducible: unknown
- Related Files: /app/node_modules/react-dom/cjs/react-dom-server.browser.production.js

---
## [ERR-20260429-003] exec-elevated-heartbeat-hosts

**Logged**: 2026-04-29T16:35:10Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
A heartbeat-triggered attempt to add local site hostnames to `/etc/hosts` could not use `exec` with `elevated=true` because elevated tools are disabled for the `heartbeat` provider in this runtime.

### Error
```
elevated is not available right now (runtime=direct).
Failing gates: allowFrom (tools.elevated.allowFrom.<provider> / agents.list[].tools.elevated.allowFrom.<provider>)
Context: provider=heartbeat session=agent:main:main:heartbeat
Fix-it keys:
- tools.elevated.enabled
- tools.elevated.allowFrom.<provider>
- agents.list[].tools.elevated.enabled
- agents.list[].tools.elevated.allowFrom.<provider>
```

### Context
- Command attempted: Python append to `/etc/hosts` with `elevated=true`
- Goal: unblock local hostname resolution for `medio-punto.local`, `mytzuko.local`, `stitching-octopus.local`, and `nctto.local`
- Verification before the failure showed the ingress-backed sites already answer `HTTP/2 200` when forced with `curl --resolve ... 172.18.255.200`
- The blocker is capability gating in heartbeat sessions, not site availability or ingress routing

### Suggested Fix
Use a non-heartbeat session/provider with elevated tool access, or adjust OpenClaw policy to allow `tools.elevated` for the `heartbeat` provider before trying to modify `/etc/hosts` from heartbeat automation.

### Metadata
- Reproducible: yes
- Related Files: /etc/hosts, .learnings/ERRORS.md

---
## [ERR-20260429-004] lumenforge_local_mongo_missing_listener

**Logged**: 2026-04-29T16:56:44Z
**Priority**: high
**Status**: resolved
**Area**: infra

### Summary
Recurring background local preview starts for `lumenforge` crashed because `MONGODB_URI=mongodb://localhost:27017/test` was set while nothing was listening on port `27017`.

### Error
```text
MongoServerSelectionError: connect ECONNREFUSED ::1:27017, connect ECONNREFUSED 127.0.0.1:27017
```

### Context
- Failed background sessions included `delta-coral` on port `3104` and `cool-willow` on port `3116`
- Process logs showed `next start` came up, then died on first Mongo server selection attempt
- No host MongoDB service or existing Docker Mongo container was listening on `27017`
- The workspace copy at `github/bshop341-b/lumenforge` is the active repo with installed dependencies

### Suggested Fix
For local preview workflows that pin `MONGODB_URI` to `mongodb://localhost:27017/test`, ensure a local MongoDB listener exists first, for example by starting a lightweight Docker `mongo:7` container mapped to `27017`.

### Metadata
- Reproducible: yes
- Related Files: github/bshop341-b/lumenforge/src/lib/mongodb.ts, github/bshop341-b/lumenforge/.env.example
- See Also: ERR-20260428-006

### Resolution
- **Resolved**: 2026-04-29T16:56:44Z
- **Notes**: Started `openclaw-local-mongo` with Docker on `127.0.0.1:27017`, verified `db.adminCommand({ ping: 1 })` returned `{ ok: 1 }`, and confirmed both local preview ports `3104` and `3116` were serving `HTTP 200` again.

---
## [ERR-20260429-005] lumenforge_preview_restart_false_negative

**Logged**: 2026-04-29T16:58:50Z
**Priority**: medium
**Status**: resolved
**Area**: infra

### Summary
A follow-up `npm start` attempt for `github/bshop341-b/lumenforge` reported failure because port `3104` was already occupied, but the existing `next-server` in that repo was already healthy.

### Error
```text
sh: next: command not found
Error: listen EADDRINUSE: address already in use :::3104
```

### Context
- The active repo is `github/bshop341-b/lumenforge` where `package.json` defines `"start": "next start"`
- A `next-server (v16.1.6)` process with cwd `/Users/rfcku/.openclaw/workspace/github/bshop341-b/lumenforge` was already listening on `3104`
- `curl -I http://127.0.0.1:3104` returned `HTTP/1.1 200 OK`
- The retry surfaced as a failed async command even though the preview target was already up

### Suggested Fix
Before rerunning `npm start` for this local preview, check whether `3104` is already serving the app. Treat `EADDRINUSE` as a prompt to verify the existing listener instead of blindly retrying.

### Metadata
- Reproducible: yes
- Related Files: github/bshop341-b/lumenforge/package.json
- See Also: ERR-20260429-004

### Resolution
- **Resolved**: 2026-04-29T16:58:50Z
- **Notes**: Verified PID `65965` was listening on `3104` from the `lumenforge` repo and confirmed the app returned `HTTP 200`, so no restart was needed.

---
## [ERR-20260429-006] lumenforge_build_missing_mongodb_uri

**Logged**: 2026-04-29T17:35:57Z
**Priority**: medium
**Status**: resolved
**Area**: config

### Summary
A background `npm run build` for `github/bshop341-b/lumenforge` failed because `src/lib/env.ts` requires `MONGODB_URI` during build-time page data collection for `/bootcamp`.

### Error
```text
❌ Invalid environment variables: {
  MONGODB_URI: [ 'Invalid input: expected string, received undefined' ]
}
Error: Failed to collect configuration for /bootcamp
Error: Failed to collect page data for /bootcamp
```

### Context
- Failed async session: `gentle-nudibranch`
- Command attempted: `npm run build`
- Repo: `github/bshop341-b/lumenforge`
- `src/lib/env.ts` validates `MONGODB_URI` eagerly with `z.string().url()` and throws on missing values
- A follow-up build in the same repo succeeded with `MONGODB_URI=mongodb://localhost:27017/test npm run build`
- The successful build still emitted `Notion API key or Services DB ID is missing.` warnings, but those did not stop static generation

### Suggested Fix
For local or CI build verification, load `.env.local` from `.env.example` before `npm run build`, or provide a safe temporary `MONGODB_URI` when only validating the build. If `/bootcamp` can be built without a live database, consider deferring or relaxing that env validation path.

### Metadata
- Reproducible: yes
- Related Files: github/bshop341-b/lumenforge/src/lib/env.ts, github/bshop341-b/lumenforge/.env.example, github/bshop341-b/lumenforge/README.md
- See Also: ERR-20260429-004

### Resolution
- **Resolved**: 2026-04-29T17:35:57Z
- **Notes**: Confirmed the failure was missing build-time env, then reran the build with a temporary local Mongo URI and got a successful Next.js production build.

---
