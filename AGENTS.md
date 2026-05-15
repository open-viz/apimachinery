# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `go.openviz.dev/apimachinery` — the OpenViz project's Kubernetes API types, generated clients, OpenAPI definitions, and CRD manifests covering Grafana resources (dashboards, datasources, dashboard templates) plus UI-server API types. Library only; other OpenViz components (`grafana-tools`, `grafana-sdk`, the installer/UI server) depend on this module.

## Architecture

- `apis/` — Kubernetes API type definitions, split by group:
  - `apis/openviz/v1alpha1/` — group `openviz.dev` (`GrafanaDashboard`, `GrafanaDatasource`, `GrafanaDashboardTemplate`, etc.).
  - `apis/ui/v1alpha1/` — group `ui.openviz.dev` (UI-server projected views).
  - Each group has `register.go`, `install/` (scheme registration) and `fuzzer/` (round-trip fuzz helpers).
  - `*_types.go` is hand-written; `zz_generated.*.go` is produced by codegen.
- `client/` — generated typed clientsets, listers, informers (k8s.io/client-go style). Do not hand-edit.
- `crds/` — generated CRD YAML manifests plus `lib.go` exposing them via `go:embed`.
- `openapi/` — generated OpenAPI definitions.
- `hack/gencrd/`, `hack/scripts/`, `hack/build.sh`, `hack/test.sh`, `hack/fmt.sh` — codegen and build helpers invoked by the Makefile inside Docker.
- `vendor/` — vendored Go deps.
- API group/version pairs come from `Makefile`: `API_GROUPS := openviz:v1alpha1 ui:v1alpha1`; codegen targets fan out over this list.

## Common commands

All build/test/lint targets run inside the `ghcr.io/appscode/golang-dev` Docker image — Docker must be running.

- `make ci` — full CI pipeline: `verify check-license lint build unit-tests`. Run before opening a PR.
- `make gen` — regenerate everything: `clientset manifests openapi`. Run after any change to `apis/**/*_types.go` or `apis/**/groupversion_info.go`.
- `make manifests` — regenerate CRDs only (`gen-crds patch-crds label-crds`).
- `make clientset` — regenerate `client/` only.
- `make openapi` — regenerate OpenAPI definitions only.
- `make fmt` — gofmt + goimports across `apis client crds hack/gencrd`.
- `make lint` — golangci-lint.
- `make unit-tests` / `make test` — Go unit tests.
- `make verify` — `verify-gen` (re-run `gen fmt` and confirm tree is clean) + `verify-modules` (`go mod tidy && go mod vendor` clean).
- `make add-license` / `make check-license` — manage license headers.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/openviz/v1alpha1/... -run TestName -v
```

CI also applies CRDs against a Kind cluster matrix via `.github/workflows/ci.yml`; if you change a CRD, ensure it still applies cleanly there. `.github/workflows/update-crds.yml` automates pulling CRD updates from sibling repos.

## Conventions

- Module path is `go.openviz.dev/apimachinery` (vanity URL); imports must use that, not the GitHub URL (`open-viz/apimachinery`).
- Do not hand-edit any file starting with `zz_generated.` or anything under `client/`, `openapi/openapi_generated.go`, or `crds/*.yaml`. Change the source in `apis/**/*_types.go` and re-run `make gen`.
- License: Apache-2.0; new files need the standard AppsCode header (`make add-license`).
- All contributions must be signed off (`git commit -s`) per the DCO file. See `CONTRIBUTING.md`.
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
