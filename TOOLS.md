# TOOLS.md - Local Notes

Skills define _how_ tools work. This file is for _your_ specifics — the stuff that's unique to your setup.

## What Goes Here

Things like:

- Camera names and locations
- SSH hosts and aliases
- Preferred voices for TTS
- Speaker/room names
- Device nicknames
- Anything environment-specific

## Examples

```markdown
### Cameras

- living-room → Main area, 180° wide angle
- front-door → Entrance, motion-triggered

### SSH

- home-server → 192.168.1.100, user: admin

### TTS

- Preferred voice: "Nova" (warm, slightly British)
- Default speaker: Kitchen HomePod
```

## Why Separate?

Skills are shared. Your setup is yours. Keeping them apart means you can update skills without losing your notes, and share skills without leaking your infrastructure.

---

## Docker on this host

- Detached `docker build`, `docker pull`, and `docker buildx imagetools inspect` runs can surface as `SIGKILL` with little or no useful output.
- Prefer bounded foreground Docker checks when diagnosing image/build issues so the cwd, full command, and failure point are preserved.
- After an async Docker failure, check for orphaned `docker` or `docker-buildx` processes before retrying.

Add whatever helps you do your job. This is your cheat sheet.
