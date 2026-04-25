# Errors

Command failures and integration errors.

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
