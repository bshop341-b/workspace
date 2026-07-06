### Operational Flow: Ditto CLI Steward Agent

This agent runs a scheduled, highly structured workflow to triage, develop, and submit GitHub issues for the `ditto-cli` repository.

**Objective:** To autonomously pull open, non-redundant issues from GitHub, implement solutions, and submit them as clean Pull Requests against `main`. The agent must maintain a clean working state after every run.

**Current Detailed Workflow (MUST FOLLOW THIS ORDER):**
1. **Setup & Sync:** Navigate to the repo (`~/.openclaw/workspace/github/rfcku/ditto-cli`) and ensure the local branch is set to `main` and fully synchronized against `origin/main`.
2. **Triage:** Check open GitHub issues and PRs. The selection criteria must prioritize:
    *   Avoiding any issue that already has an active PR or clear evidence of ongoing work.
    *   **Priority Order for Bugs (TUI):** #135 > #136 > #138 > #137.
    *   Generally: Bugs over Enhancements; Small/Complete fixes over large features.
3. **Issue Selection:** Select exactly one actionable open issue based on the defined priorities. If no issues are found or none are actionable, stop and report this status clearly to `#bot`.
4. **Development:** Create a new feature branch from `main`, ensuring it is updated, and give it a clear name related to the issue/fix. Implement the full solution with small, validating commits, run relevant validations, commit changes, and push the new branch.
5. **Review & PR:** Open a Pull Request (PR) targeting `main`. Immediately comment on the opened PR with `@copilot review`.
6. **Cleanup:** Switch back to local `main`, perform `pull/ff-only` to ensure the repository is clean and ready for the next cycle.
7. **Reporting:** Conclude by posting a single, brief status message specifically for `#bot`, stating exactly which issue was addressed, what branch/PR was opened, or if no issues were found.

**Critical Rules & Constraints:**
*   Must *always* leave the local repository in a clean state (no uncommitted changes).
*   Avoid creating stacked PRs unless strictly necessary and justified by the flow.
*   If an issue is blocked by external factors not resolvable within this agent's scope, report the block briefly to `#bot` and proceed immediately to checking the next highest-priority/most profitable issue.

**Note:** This agent handles complex git state management and external API interactions (GitHub). Ensure it has all necessary credentials and permissions.