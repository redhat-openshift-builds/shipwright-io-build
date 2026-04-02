---
name: trigger-sync
description: >-
  Trigger the upstream sync workflow for shipwright mirror repos. Use when the
  user asks to sync, trigger sync, run sync workflow, or update main from upstream.
---

# Trigger Upstream Sync

## Usage

Run the sync workflow for the build mirror repo:

```bash
gh workflow run sync-upstream.yml --ref sync-upstream -R redhat-openshift-builds/shipwright-io-build
```

Run the sync workflow for the CLI mirror repo:

```bash
gh workflow run sync-upstream.yml --ref sync-upstream -R redhat-openshift-builds/shipwright-cli-mirror
```

## Check status

After triggering, check the run status:

```bash
gh run list --workflow=sync-upstream.yml -R redhat-openshift-builds/shipwright-io-build --limit 1
```

If a run failed, view the logs:

```bash
gh run view <RUN_ID> --log-failed -R redhat-openshift-builds/shipwright-io-build
```
