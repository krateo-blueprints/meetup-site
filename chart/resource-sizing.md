# Resource Sizing Guidelines for meetup-site Blueprint

When configuring resource requests and limits in the `meetup-site` blueprint, ensure container CPU and memory requests do not exceed their corresponding limits.

## Validation Rules
- `resources.requests.cpu` must be less than or equal to `resources.limits.cpu` (if limits are set).
- `resources.requests.memory` must be less than or equal to `resources.limits.memory` (if limits are set).

Failure to maintain valid ratios will cause Kubernetes admission webhooks to reject deployment updates with an `Invalid` field error, leaving the composition in `Synced=False`.
