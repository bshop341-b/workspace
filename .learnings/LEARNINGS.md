# Learnings

Corrections, insights, and knowledge gaps captured during development.

**Categories**: correction | insight | knowledge_gap | best_practice

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
