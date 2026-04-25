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
