# The Charts · Go Leaderboard API Sketch

[简体中文](README.zh-CN.md)

A Go practice project that sketches an HTTP API for submitting scores and reading rankings. The intended design uses Redis for ranking data.

| Planned endpoint | Purpose |
| --- | --- |
| `POST /rank` | Submit a user's score. |
| `GET /rank/:username` | Look up one user's rank. |
| `GET /rank` | List the leaderboard. |

[`main.go`](main.go) sketches the server entry point. [`api/route.go`](api/route.go) declares the Gin routes and contains commented-out handler sketches. The checkout has no wired handlers or Redis implementation, and the module path in [`go.mod`](go.mod) differs from the import in `main.go`.
