# sync-upstream branch

This branch exists solely to host the GitHub Actions workflow that syncs `main` with upstream.

The actual codebase lives on the `main` branch.

## How it works

- A daily cron (06:00 UTC) or manual trigger runs the sync workflow.
- The workflow fetches `main` from upstream (`shipwright-io/build`) and fast-forwards this repo's `main` to match.
- Uses a deploy key (via checkout `ssh-key`) because `GITHUB_TOKEN` cannot push changes to `.github/workflows/` files.
- `main` has no mirror-only commits, so the sync is always a fast-forward.

## Manual trigger

```bash
gh workflow run sync-upstream.yml --ref sync-upstream -R redhat-openshift-builds/shipwright-io-build
```
