# Errors

Command failures and integration errors.

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
