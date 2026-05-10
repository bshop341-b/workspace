### Operational Flow: Ditto CLI Steward Agent

This agent runs a scheduled, highly structured workflow to triage, develop, and submit GitHub issues for the `ditto-cli` repository.

**Objective:** To autonomously pull open, non-redundant issues from GitHub, implement solutions, and submit them as clean Pull Requests against `main`.

**Critical Steps:**
1. **Setup:** Sync local repo to `main` and `origin/main`.
2. **Triage:** Identify the highest priority issue from open GitHub issues/PRs, avoiding duplicated effort.
3. **Development:** Create a feature branch, implement the fix with small, validating commits, and push.
4. **Review:** Open a PR against `main` and automatically comment `@copilot review`.
5. **Cleanup:** Switch back to local `main`, pull/ff-only, ensuring a clean working state for the next run.
6. **Reporting:** Provide a brief, dedicated status message for the bot.

**Target Repository:** `github/ditto-cli`
**Priority Order for Bugs:** 1. #135, 2. #136, 3. #138, 4. #137. (Then generally: Bugs > Enhancements; Small/Complete > Big Features)

**Constraint:** Must only report the action taken (Issue ID/PR link, or lack of action). Never leave the local repo in a dirty state.

**Note:** This agent handles complex git state management and external API interactions (GitHub). Ensure it has all necessary credentials and permissions.