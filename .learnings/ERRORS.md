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
## [ERR-20260430-001] python_html_parse

**Logged**: 2026-04-30T09:12:00-07:00
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
Tried to parse a local HTML page with BeautifulSoup, but `bs4` is not installed in this workspace environment.

### Error
```
ModuleNotFoundError: No module named 'bs4'
```

### Context
- Command/operation attempted: inline Python parse of `http://127.0.0.1:3104/projects`
- Input or parameters used: `python3` script importing `from bs4 import BeautifulSoup`
- Environment details: OpenClaw workspace on macOS, default Python environment

### Suggested Fix
Prefer stdlib-based HTML checks for quick heartbeat validations, or verify dependency availability before importing third-party parsing libraries.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260501-001] git_status_non_repo_obsidian_vault

**Logged**: 2026-05-01T08:24:00-07:00
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
Tried to run `git status` against the Obsidian vault path even though the vault is not a Git repository.

### Error
```
fatal: not a git repository (or any of the parent directories): .git
```

### Context
- Operation attempted: compare local repo changes in `github/rfcku/ditto-cli` and note-file changes in `/Users/rfcku/Documents/BishopVault`
- The code repo status check worked, but the second `git -C` assumed the vault itself was versioned
- No data loss occurred; the note files were still edited successfully

### Suggested Fix
When checking note updates under the Obsidian vault, use plain file reads or `ls`/`stat` unless the specific vault path is known to be a Git repo.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260501-002] mytzuko_preview_sigkill_false_negative

**Logged**: 2026-05-01T09:02:15-07:00
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
An async `npm start` for `github/nctto/mytzuko` surfaced as `signal SIGKILL` in exec-event even though the detached `next-server` stayed up and continued serving port `3102`.

### Error
```text
Exec failed (swift-tr, signal SIGKILL) :: > mytzuko@0.1.0 start
> next start
▲ Next.js 16.1.6 - Local: http://localhost:3102
✓ Starting...
✓ Ready in 224ms
```

### Context
- Operation attempted: `npm start` for the Mytzuko repo
- Verification after the async failure showed PID `37882` (`next-server v16.1.6`) still listening on `3102`
- `lsof -a -p 37882 -d cwd` resolved the process cwd to `/Users/rfcku/.openclaw/workspace/github/nctto/mytzuko`
- `curl -I http://127.0.0.1:3102` returned `HTTP/1.1 200 OK`
- The wrapper failure was therefore a false negative for preview health, not an app crash

### Suggested Fix
When a long-running preview start reports `SIGKILL` after logging `Ready`, immediately verify the target port and existing listener before retrying. Prefer explicit background/session handling for durable preview processes.

### Metadata
- Reproducible: unknown
- Related Files: github/nctto/mytzuko/package.json, .learnings/ERRORS.md
- See Also: ERR-20260428-002, ERR-20260429-005

### Resolution
- **Resolved**: 2026-05-01T09:02:15-07:00
- **Notes**: Confirmed the Mytzuko `next-server` remained healthy on port `3102`, so no restart or code change was needed.

---

## [ERR-20260501-003] cron_jobs_json_shape_and_croner_resolution

**Logged**: 2026-05-01T21:39:00Z
**Priority**: low
**Status**: pending
**Area**: config

### Summary
Two quick cron-maintenance probes failed because `~/.openclaw/cron/jobs.json` is an object with a `jobs` array, not a raw array, and because the bundled `croner` package is only resolvable from OpenClaw's installed node_modules path.

### Error
```
AttributeError: 'str' object has no attribute 'get'
Cannot find module 'croner'
```

### Context
- Operation attempted: summarize and reschedule OpenClaw cron jobs from the local JSON store
- The first parser assumed the JSON root was a list of jobs
- The first Node probe assumed `croner` was available from the workspace module path
- The working approach was to read `data.jobs` and require `/opt/homebrew/lib/node_modules/openclaw/node_modules/croner`

### Suggested Fix
When scripting against OpenClaw cron locally, treat `jobs.json` as `{ version, jobs[] }` and import bundled dependencies from the installed OpenClaw path instead of assuming workspace-local resolution.

### Metadata
- Reproducible: yes
- Related Files: /Users/rfcku/.openclaw/cron/jobs.json, .learnings/ERRORS.md

---
## [ERR-20260502-001] curl-path

**Logged**: 2026-05-02T15:36:10Z
**Priority**: low
**Status**: pending
**Area**: infra

### Summary
A Ditto smoke-test command failed because `curl` was not on PATH in this shell.

### Error
```
zsh:11: command not found: curl
zsh:12: command not found: awk
```

### Context
- Command attempted: shell script using `curl` and `awk` by bare command name inside OpenClaw exec
- Environment: macOS host shell under OpenClaw
- Follow-up approach: use absolute tool paths like `/usr/bin/curl` and `/usr/bin/awk` when PATH looks minimal

### Suggested Fix
Prefer absolute paths for common system binaries in OpenClaw exec scripts when PATH-dependent failures appear.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260502-002] playwright_waitforurl_timeout

**Logged**: 2026-05-02T16:22:00Z
**Priority**: high
**Status**: resolved
**Area**: frontend

### Summary
A Ditto browser smoke test timed out waiting for login to navigate back to `https://ditto.local/`, which confirmed the UI auth flow was still broken even after the backend `/v1` fix deployed.

### Error
```
page.waitForURL: Timeout 30000ms exceeded.
waiting for navigation to "https://ditto.local/" until "load"
```

### Context
- Operation attempted: Playwright smoke test for `login -> redirect/home` on `ditto.local`
- Relevant upstream state: `api.ditto.local/v1` had already been revalidated and was returning healthy auth and user responses with a valid token
- Observed frontend behavior around the same investigation: the bundle exposed `API URL: undefined` and auth requests were hitting same-origin frontend routes instead of `https://api.ditto.local/v1`
- This failure strengthened the conclusion that the remaining blocker is frontend build/runtime wiring, not the restored backend API

### Suggested Fix
Rebuild or rewire the Ditto UI with `NEXT_PUBLIC_API_URL=https://api.ditto.local/v1` available at build time, then rerun the full browser smoke test.

### Metadata
- Reproducible: yes
- Related Files: /Users/rfcku/Documents/BishopVault/Projects/Ditto/Hallazgo UI auth local 2026-05-02.md

### Resolution
- **Resolved**: 2026-05-04T15:58:00Z
- **Notes**: A fresh Playwright smoke check on `ditto.local` now reaches `https://api.ditto.local/v1/auth/register` and `https://api.ditto.local/v1/auth/authorize`, lands on `https://ditto.local/`, and preserves `token`, `user`, and `id` in localStorage after reload. See `/Users/rfcku/.openclaw/workspace/tmp/ditto-login-check.log` and `/Users/rfcku/.openclaw/workspace/tmp/ditto-login-verify.log`.

---
## [ERR-20260502-003] ditto_login_ui_exception

**Logged**: 2026-05-02T16:22:30Z
**Priority**: high
**Status**: resolved
**Area**: frontend

### Summary
The Ditto login page returned `200` responses for its RSC requests but still fell into an uncaught exception, leaving the UI unable to complete login.

### Error
```
response 200 https://ditto.local/auth/login?_rsc=...
response 200 https://ditto.local/auth/login?_rsc=...
node:internal/process/promises:332 triggerUncaughtException(err, true /* fromPromise */)
```

### Context
- Operation attempted: browser automation against `https://ditto.local/auth/login`
- The login page itself loaded, so the failure was not simple reachability or TLS
- Combined with `POST https://ditto.local/auth/authorize -> 404` and `GET https://ditto.local/users/me -> 404`, the evidence points to a broken frontend auth configuration or compiled bundle rather than a dead backend endpoint

### Suggested Fix
Inspect the real UI source or build pipeline for the active `vora-web` image, fix the API base URL wiring, and capture the browser console and network trace again after rebuilding.

### Metadata
- Reproducible: yes
- Related Files: /Users/rfcku/Documents/BishopVault/Projects/Ditto/Hallazgo UI auth local 2026-05-02.md

### Resolution
- **Resolved**: 2026-05-04T15:59:00Z
- **Notes**: Re-ran the flow with corrected Playwright scripts and confirmed the login page now posts to `https://api.ditto.local/v1/auth/authorize`, receives `200`, redirects to the homepage, and leaves the session in localStorage. The earlier uncaught-exception conclusion is no longer supported by the current local repro.

---
## [ERR-20260502-004] docker_buildx_imagetools_orphan_after_exec_sigkill

**Logged**: 2026-05-02T09:34:00-07:00
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
An async `docker buildx imagetools inspect ghcr.io/rfcku/vora-web:main` surfaced as `signal SIGKILL` with no captured output, but the orphaned `docker` and `docker-buildx` processes kept running until manually killed.

### Error
```text
Exec failed (gentle-f, signal SIGKILL)
(no output recorded)
```

### Context
- Operation attempted: inspect the remote `ghcr.io/rfcku/vora-web:main` image while debugging the local Ditto/Vora auth issue
- After the exec-event failure, `process list` showed `gentle-fjord` failed and `process poll` confirmed `SIGKILL`
- `ps` showed orphaned PID `64805` (`docker buildx imagetools inspect ...`) with child PID `64809` (`docker-buildx ...`) still sleeping under parent PID `1`
- Manual `kill -9 64805 64809` was required to clean up the stale inspection job

### Suggested Fix
When an async one-shot Docker inspection reports `SIGKILL` with no log output, immediately check for orphaned `docker` or `docker-buildx` processes and clean them up before retrying. Prefer a bounded foreground run for quick image inspection tasks so failures return real output instead of leaving background stragglers.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260428-002

---
## [ERR-20260503-001] lumenforge_preview_sigkill_false_negative

**Logged**: 2026-05-03T07:08:37-07:00
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
An async `npm start -- --hostname 0.0.0.0 --port 3104` for `github/bshop341-b/lumenforge` surfaced as `signal SIGKILL` after Next.js reported `Ready`, but the detached `next-server` kept serving port `3104`.

### Error
```text
Exec failed (mellow-l, signal SIGKILL) :: > nctto@0.1.0 start
> next start --hostname 0.0.0.0 --port 3104
▲ Next.js 16.1.6 - Local: http://localhost:3104
- Network: http://0.0.0.0:3104
✓ Starting...
✓ Ready in 468ms
```

### Context
- Operation attempted: start the local preview for the `nctto` package in `github/bshop341-b/lumenforge`
- Verification after the async failure showed PID `50494` (`next-server v16.1.6`) still listening on `3104`
- `lsof -a -p 50494 -d cwd` resolved the process cwd to `/Users/rfcku/.openclaw/workspace/github/bshop341-b/lumenforge`
- `curl -I http://127.0.0.1:3104` returned `HTTP/1.1 200 OK`
- The exec-event therefore reported a false negative for preview health, not an app crash

### Suggested Fix
When an async preview start reports `SIGKILL` after logging `Ready`, immediately verify the target port and process cwd before retrying. Prefer explicit background/session handling for durable Next.js preview processes.

### Metadata
- Reproducible: unknown
- Related Files: github/bshop341-b/lumenforge/package.json, .learnings/ERRORS.md
- See Also: ERR-20260429-005, ERR-20260501-002

### Resolution
- **Resolved**: 2026-05-03T07:08:37-07:00
- **Notes**: Confirmed the Lumenforge `next-server` remained healthy on port `3104`, so no restart or code change was needed.

---
## [ERR-20260503-002] medio_punto_preview_sigkill_false_negative

**Logged**: 2026-05-03T08:03:34-07:00
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
An async `npm start` for `github/rfcku/medio-punto` surfaced as `signal SIGKILL` after Next.js reported `Ready`, but the detached `next-server` kept serving port `3101`.

### Error
```text
Exec failed (mild-orb, signal SIGKILL) :: > medio-punto@0.1.0 start
> next start
▲ Next.js 14.1.0 - Local: http://localhost:3101
✓ Ready in 140ms
```

### Context
- Operation attempted: start the local preview for `github/rfcku/medio-punto`
- Verification after the async failure showed PID `51988` (`next-server`) still listening on `3101`
- `lsof -a -p 51988 -d cwd` resolved the process cwd to `/Users/rfcku/.openclaw/workspace/github/rfcku/medio-punto`
- `curl -I http://127.0.0.1:3101` returned `HTTP/1.1 200 OK`
- The exec-event therefore reported a false negative for preview health, not an app crash

### Suggested Fix
When an async preview start reports `SIGKILL` after logging `Ready`, immediately verify the target port and process cwd before retrying. Prefer explicit background/session handling for durable Next.js preview processes.

### Metadata
- Reproducible: unknown
- Related Files: github/rfcku/medio-punto/package.json, .learnings/ERRORS.md
- See Also: ERR-20260429-005, ERR-20260501-002, ERR-20260503-001

### Resolution
- **Resolved**: 2026-05-03T08:03:34-07:00
- **Notes**: Confirmed the Medio Punto `next-server` remained healthy on port `3101`, so no restart or code change was needed.

---
## [ERR-20260503-001] exec.elevated

**Logged**: 2026-05-03T15:38:51Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Elevated exec is unavailable in exec-event heartbeat sessions, so `/etc/hosts` updates cannot be performed from this runtime.

### Error
```
elevated is not available right now (runtime=direct).
Failing gates: allowFrom (tools.elevated.allowFrom.<provider> / agents.list[].tools.elevated.allowFrom.<provider>)
Context: provider=exec-event session=agent:main:cron:62c86454-1f85-4f4d-9c06-3be649c4b404:heartbeat
```

### Context
- Operation attempted: `./scripts/add-local-hosts.sh 172.18.255.200 medio-punto.local mytzuko.local stitching-octopus.local nctto.local`
- Workdir: `/Users/rfcku/.openclaw/workspace/github/argocd`
- Goal: unblock normal browser access for local preview sites without `--resolve`
- Relevant doc: `github/argocd/workloads/local-sites/README.md`

### Suggested Fix
Run the hosts update from a user-approved elevated session or outside exec-event heartbeat runtime, or enable elevated exec for the `exec-event` provider if that is an intended capability.

### Metadata
- Reproducible: yes
- Related Files: github/argocd/workloads/local-sites/README.md

---
## [ERR-20260503-002] exec.python

**Logged**: 2026-05-03T16:08:55Z
**Priority**: low
**Status**: pending
**Area**: infra

### Summary
The  binary is absent on this host; use  for quick local checks.

### Error


### Context
- Operation attempted: inline HTML inspection script against 
- Workdir: 
- Goal: confirm whether the Mytzuko local preview still exposed legacy booking form markers

### Suggested Fix
Default to  instead of  when running ad hoc scripts on this macOS host.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---
## [ERR-20260503-002] exec.python

**Logged**: 2026-05-03T16:10:00Z
**Priority**: low
**Status**: pending
**Area**: infra

### Summary
The `python` binary is absent on this host, use `python3` for quick local checks.

### Error
```
zsh:1: command not found: python
```

### Context
- Operation attempted: inline HTML inspection script against `http://127.0.0.1:3102`
- Workdir: `/Users/rfcku/.openclaw/workspace`
- Goal: confirm whether the Mytzuko local preview still exposed legacy booking form markers

### Suggested Fix
Default to `python3` instead of `python` when running ad hoc scripts on this macOS host.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md

---

## [ERR-20260503-003] exec.heredoc_backticks

**Logged**: 2026-05-03T16:11:00Z
**Priority**: low
**Status**: pending
**Area**: docs

### Summary
Backticks inside a double-quoted shell command were expanded before a heredoc append, which broke the logging command.

### Error
```
zsh:1: command not found: python
zsh:2: command not found: zsh:1:
zsh:1: no such file or directory: http://127.0.0.1:3102
zsh:1: permission denied: /Users/rfcku/.openclaw/workspace
zsh:1: command not found: python
```

### Context
- Operation attempted: append a markdown error entry to `.learnings/ERRORS.md`
- Cause: the surrounding shell command used double quotes, so backticks in markdown content triggered command substitution before the heredoc ran

### Suggested Fix
When appending markdown that contains backticks, use a single-quoted heredoc delimiter and avoid wrapping the full shell snippet in double quotes with raw backticks inside.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260426-005

---

## [ERR-20260503-004] docker-pull-hang

**Logged**: 2026-05-03T16:24:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
Docker image resolution for `node:22-alpine` hung with no progress output in this runtime, so container builds for Stitching Octopus could not be verified even though the app's local Next.js production build succeeded.

### Error
```text
#2 [internal] load metadata for docker.io/library/node:22-alpine
Exec failed (marine-n, signal SIGKILL)
docker pull node:22-alpine  -> no output until harness timeout / SIGKILL
```

### Context
- Operation attempted: `docker build -t stitchingoctopus:test .` in `github/rfcku/stitchingoctopus`
- A direct `npm run build` in the same repo completed successfully, so the app code itself was not the blocker
- Reproduction also hung for `docker pull node:22-alpine` and `docker buildx imagetools inspect node:22-alpine`
- Host-side network reachability to `https://registry-1.docker.io/v2/` was fine, which points more toward Docker Desktop / daemon image-resolution behavior than the project files

### Suggested Fix
Treat this as a Docker runtime issue first. If container verification is required, troubleshoot Docker Desktop image pull behavior or use a preloaded/local base image before spending more time on app code.

### Metadata
- Reproducible: yes
- Related Files: github/rfcku/stitchingoctopus/Dockerfile, github/rfcku/stitchingoctopus/.dockerignore
- See Also: ERR-20260428-002

---
## [ERR-20260503-005] kanban-backend-docker-metadata-hang

**Logged**: 2026-05-03T16:33:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
The `kanban-dashboard-backend` Docker build still hung at image metadata resolution in this runtime even after removing the external Dockerfile frontend dependency from the repo Dockerfile.

### Error
```text
#2 resolve image config for docker-image://docker.io/docker/dockerfile:1
Process exited with signal SIGKILL.

#2 [internal] load metadata for docker.io/library/node:20-bookworm-slim
Process exited with signal SIGKILL.
```

### Context
- Operation attempted: `docker build --progress=plain -t local/kanban-dashboard-backend:test .` in `github/bshop341-b/kanban-dashboard-backend`
- The original Dockerfile used `# syntax=docker/dockerfile:1`; removing that line avoided the initial frontend-image resolution step
- The Dockerfile was also tightened from `npm install --omit=dev` to `npm ci --omit=dev` because a lockfile is present
- Even after that repo-level cleanup, Docker still stalled while loading metadata for `node:20-bookworm-slim`, which points back to the local Docker runtime rather than the application files
- A fallback `DOCKER_BUILDKIT=0 docker build ...` also ended in `SIGKILL` after only printing the legacy-builder deprecation warning, so this is not limited to BuildKit frontend resolution
- Orphaned `docker build` / `docker buildx imagetools inspect` processes from earlier SIGKILLs had to be cleaned up manually

### Suggested Fix
Keep the Dockerfile simplification, but treat the remaining failure as a Docker daemon or registry-resolution problem. Before changing app code further, verify Docker Desktop can resolve/pull standard images cleanly or use a preloaded local base image.

### Metadata
- Reproducible: yes
- Related Files: github/bshop341-b/kanban-dashboard-backend/Dockerfile, .learnings/ERRORS.md
- See Also: ERR-20260428-002, ERR-20260502-004, ERR-20260503-004

---
## [ERR-20260503-006] async_docker_build_sigkill_cluster

**Logged**: 2026-05-03T16:36:00Z
**Priority**: medium
**Status**: pending
**Area**: infra

### Summary
A follow-up cluster of detached Docker build sessions also ended in `SIGKILL`, reinforcing that Docker verification is currently unreliable in this runtime and that async builds are surfacing too little context to diagnose precisely.

### Error
```text
Exec failed (amber-sa, signal SIGKILL)
Exec failed (good-sum, signal SIGKILL)
Exec failed (nova-sum, signal SIGKILL)
Exec failed (clear-sl, signal SIGKILL)

# clear-sl excerpt
DEPRECATED: The legacy builder is deprecated and will be removed in a future release.
BuildKit is currently disabled; enable it by removing the DOCKER_BUILDKIT=0 environment variable.
```

### Context
- Operation attempted: previously launched async Docker build checks after earlier metadata-resolution hangs
- Surfaced output showed multiple builds using the `desktop-linux` Docker driver, with at least two runs loading small Dockerfiles (`283B` and `248B`) before being killed
- One follow-up used the legacy builder path with `DOCKER_BUILDKIT=0` and still ended in `SIGKILL`, so the instability is not limited to BuildKit-only flows
- By the time this completion event arrived, no orphaned `docker` / `docker-buildx` processes were still running
- The exact original commands and working directories were not recoverable from the exec-event payload alone

### Suggested Fix
Avoid detached one-shot Docker builds when the goal is diagnosis. Run bounded foreground checks from the target repo so the cwd, full command, and failure point are preserved, and continue treating this as a host Docker runtime issue before changing more application code.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260502-004, ERR-20260503-004, ERR-20260503-005

---
## [ERR-20260504-001] async_node_curl_sigterm

**Logged**: 2026-05-04T15:01:00Z
**Priority**: low
**Status**: pending
**Area**: infra

### Summary
A detached exec against node `grand-pi` surfaced only as a `SIGTERM` on session `grand-pine`, with no captured output beyond a truncated `curl /dev/null` summary.

### Error
```text
Exec failed (grand-pi, signal SIGTERM)
Process exited with signal SIGTERM.
```

### Context
- Operation attempted: previously launched async exec session `grand-pine`
- `process list` showed the command summary only as `curl /dev/null`
- `process log` and `process poll` returned no command output, so the target URL, cwd, and reason for termination were not recoverable from the completion event alone
- The exec-event explicitly asked for internal handling only, so no user-facing follow-up was sent

### Suggested Fix
For detached node health checks, include the full target in the command string and emit one short startup line before blocking so later `SIGTERM` events preserve enough context to diagnose whether the process timed out, was canceled, or hit a remote-side failure.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260428-001

---
## [ERR-20260504-002] medio_punto_preview_sigkill_false_negative_repeat

**Logged**: 2026-05-04T15:15:30Z
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
Another async `npm start` for `github/rfcku/medio-punto` surfaced as `signal SIGKILL` after Next.js reported `Ready`, but the detached `next-server` remained healthy on port `3101`.

### Error
```text
Exec failed (rapid-sh, signal SIGKILL) :: > medio-punto@0.1.0 start
> next start
▲ Next.js 14.1.0 - Local: http://localhost:3101
✓ Ready in 171ms
```

### Context
- Operation attempted: start or keep alive the local preview for `github/rfcku/medio-punto`
- Verification after the async failure showed PID `11259` still listening on `3101`
- `lsof -a -p 11259 -d cwd` resolved the process cwd to `/Users/rfcku/.openclaw/workspace/github/rfcku/medio-punto`
- `curl -I http://127.0.0.1:3101` returned `HTTP/1.1 200 OK`
- This is the same false-negative pattern as ERR-20260503-002, so the preview did not need a restart

### Suggested Fix
Treat detached Next.js `SIGKILL` completion events that arrive after `✓ Ready` as suspicious but not definitive. Verify the port and process cwd before restarting or telling the user the app died.

### Metadata
- Reproducible: unknown
- Related Files: github/rfcku/medio-punto/package.json, .learnings/ERRORS.md
- See Also: ERR-20260503-002

### Resolution
- **Resolved**: 2026-05-04T15:15:30Z
- **Notes**: Confirmed the Medio Punto preview remained healthy on port `3101`, so no corrective action or user-facing alert was needed.

---
## [ERR-20260505-001] nctto_preview_sigkill_false_negative

**Logged**: 2026-05-05T15:03:48Z
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
Another async `npm start` for `github/bshop341-b/lumenforge` surfaced as `signal SIGKILL` after Next.js reported `Ready`, but the detached `next-server` remained healthy on port `3104`.

### Error
```text
Exec failed (rapid-cr, signal SIGKILL) :: > nctto@0.1.0 start
> next start
▲ Next.js 16.1.6 - Local: http://localhost:3104
- Network: http://172.18.255.200:3104
✓ Starting...
✓ Ready in 235ms
```

### Context
- Operation attempted: start or keep alive the local preview for `github/bshop341-b/lumenforge`
- Verification after the async failure showed PID `45074` (`node`) still listening on `3104`
- `lsof -a -p 45074 -d cwd` resolved the process cwd to `/Users/rfcku/.openclaw/workspace/github/bshop341-b/lumenforge`
- `curl -I http://127.0.0.1:3104` returned `HTTP/1.1 200 OK`
- This matches the earlier detached Next.js false-negative pattern, so the preview did not need a restart

### Suggested Fix
Treat detached Next.js `SIGKILL` completion events that arrive after `✓ Ready` as suspicious but not definitive. Verify the listening port and process cwd before restarting or reporting a crash.

### Metadata
- Reproducible: unknown
- Related Files: github/bshop341-b/lumenforge/package.json, .learnings/ERRORS.md
- See Also: ERR-20260503-001, ERR-20260503-002, ERR-20260504-001

### Resolution
- **Resolved**: 2026-05-05T15:03:48Z
- **Notes**: Confirmed the Lumenforge preview stayed healthy on port `3104`, so no corrective action or user-facing alert was needed.

---
## [ERR-20260504-003] playwright_ad_hoc_runner_mismatch

**Logged**: 2026-05-04T16:00:00Z
**Priority**: low
**Status**: resolved
**Area**: tests

### Summary
Several ad hoc Ditto browser-check commands failed because the temporary Playwright harness mixed incompatible runner imports and CLI flags.

### Error
```text
Error: No tests found.
error: unknown option '--test-dir=tmp'
Error: Cannot find module 'playwright/test'
```

### Context
- Operation attempted: run a quick local browser smoke test for `ditto.local` from the temp `pwcheck` sandbox
- Failed attempts included invoking `npx @playwright/test` with arguments that matched no tests, passing an unsupported `--test-dir=tmp` flag, and importing `playwright/test` instead of `@playwright/test`
- The reliable fallback in this sandbox was a plain Node script with `require('playwright')`, which then verified register, reload, and login successfully

### Suggested Fix
For one-off local smoke tests in the temp sandbox, prefer `node <script>.js` with `require('playwright')` unless a proper `@playwright/test` project and test file pattern are already in place.

### Metadata
- Reproducible: yes
- Related Files: tmp/ditto-login-check.spec.js, tmp/ditto-login-check.js, tmp/pwcheck/package.json
- See Also: ERR-20260502-002, ERR-20260502-003

### Resolution
- **Resolved**: 2026-05-04T16:00:00Z
- **Notes**: Switched the check to direct Node scripts under `tmp/pwcheck` and completed the auth verification successfully.

---
## [ERR-20260504-004] detached_home_scan_orphan_find

**Logged**: 2026-05-04T16:09:30Z
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
A small cluster of detached sessions surfaced only as `SIGTERM` or `SIGKILL`, and one broad `find /Users/rfcku` scan was still running orphaned under PID 1 after the exec-event failure.

### Error
```text
Exec failed (vivid-ri, signal SIGKILL)
Exec failed (amber-bl, signal SIGTERM)
Exec failed (amber-ot, signal SIGTERM)
Process exited with signal SIGKILL.
Process exited with signal SIGTERM.
```

### Context
- Operation attempted: previously launched async exec sessions `vivid-ridge`, `amber-bloom`, and `amber-otter`
- `process list` only preserved truncated summaries, showing two sessions as `find /Users/rfcku` and one as `python3 <<'PY'`
- `process log` produced no output for any of the three sessions, so the exact intent of the Python job and the reason for termination were not recoverable from the completion event alone
- Follow-up host inspection found an orphaned process: `find /Users/rfcku -maxdepth 5 \( -name go.mod -o -name .git \)` running with `PPID 1`
- The orphaned `find` was terminated manually with `kill 14964`
- The exec-event explicitly asked for internal handling only, so no user-facing follow-up was sent

### Suggested Fix
Avoid detached broad home-directory scans when a workspace-scoped search would do. For async discovery jobs, print one short startup line with the exact search root and query, and check for orphaned child processes when the wrapper reports only a signal with no logs.

### Metadata
- Reproducible: unknown
- Related Files: .learnings/ERRORS.md
- See Also: ERR-20260428-001, ERR-20260502-004, ERR-20260504-001

### Resolution
- **Resolved**: 2026-05-04T16:09:30Z
- **Notes**: Confirmed the detached sessions had no preserved logs, found the remaining orphaned `find` under PID 1, killed it, and kept the event internal as instructed.

---
## [ERR-20260504-001] ditto-heartbeat-api-smoke

**Logged**: 2026-05-04T17:04:00Z
**Priority**: medium
**Status**: pending
**Area**: backend

### Summary
Fresh Ditto heartbeat smoke setup hit an unexpected `/v1/auth/register` response shape and failed before creating test data.

### Error
```text
Traceback (most recent call last):
  File "<stdin>", line 39, in <module>
  File "<stdin>", line 37, in register
KeyError: 'data'
```

### Context
- Command/operation attempted: inline Python smoke setup against `https://api.ditto.local/v1`
- Expected register response shape: JSON object with `data.token` and `data.id`
- Actual issue: response body did not contain top-level `data`
- Summary output only, no tokens logged

### Suggested Fix
Capture and inspect the actual register response body before assuming the old response shape, then harden the smoke helper to tolerate validation or alternate success envelopes.

### Metadata
- Reproducible: unknown
- Related Files: /Users/rfcku/.openclaw/workspace/tmp/ditto-heartbeat-smoke.json

---
## [ERR-20260505-002] clustered_next_preview_sigkill_false_negative

**Logged**: 2026-05-05T17:27:30Z
**Priority**: low
**Status**: resolved
**Area**: infra

### Summary
Two async Next.js preview sessions surfaced as `signal SIGKILL` after `next start` reached `Ready`, but both detached servers stayed healthy on their target ports.

### Error
```text
Exec failed (quick-pr, signal SIGKILL) :: > mytzuko@0.1.0 build
... Finalizing page optimization ...
> mytzuko@0.1.0 start
> next start
▲ Next.js 16.1.6
- Local: http://localhost:3102
✓ Ready in 235ms

Exec failed (young-em, signal SIGKILL) :: > stiching-octopus@0.1.0 build
... Finalizing page optimization ...
> stiching-octopus@0.1.0 start
> next start
▲ Next.js 16.2.4
- Local: http://localhost:3103
✓ Ready in 67ms
```

### Context
- Operation attempted: previously launched async preview builds for `github/nctto/mytzuko` and `github/rfcku/stitchingoctopus`
- Follow-up verification showed PID `42494` still listening on `3102` with cwd `/Users/rfcku/.openclaw/workspace/github/nctto/mytzuko`
- Follow-up verification showed PID `42479` still listening on `3103` with cwd `/Users/rfcku/.openclaw/workspace/github/rfcku/stitchingoctopus`
- `curl -I http://127.0.0.1:3102/` returned `HTTP/1.1 200 OK`
- `curl -I http://127.0.0.1:3103/` returned `HTTP/1.1 200 OK`
- The exec-event explicitly asked for internal handling only, so no user-facing follow-up was sent

### Suggested Fix
Treat detached Next.js `SIGKILL` completion events that arrive after `✓ Ready` as likely wrapper false negatives. Verify the port and process cwd before restarting, and prefer explicit durable background/session handling when a preview server is meant to stay up.

### Metadata
- Reproducible: yes
- Related Files: .learnings/ERRORS.md, github/nctto/mytzuko/package.json, github/rfcku/stitchingoctopus/package.json
- See Also: ERR-20260505-001, ERR-20260504-003, ERR-20260504-002

### Resolution
- **Resolved**: 2026-05-05T17:27:30Z
- **Notes**: Confirmed both detached Next.js servers were still serving on ports `3102` and `3103`, so no restart or user interrupt was needed.

---
