# Go Project Structure

- `/cmd` — Main source files. For example, an application's `main.go` should live at `/cmd/foo/main.go`.
- `/internal` — Private code that other applications or libraries should not import.
- `/pkg` — Public code exposed for others to use.
- `/test` — External tests and test data. Go unit tests live in the same package as the source code; public API or integration tests, for example, should live in `/test`.
- `/configs` — Configuration files.
- `/docs` — Design and user documentation.
- `/examples` — Examples for the application and/or a public library.
- `/api` — API contract files, such as Swagger or Protocol Buffer definitions.
- `/web` — Web-application assets, such as static files.
- `/build` — Packaging and continuous-integration (CI) files.
- `/scripts` — Scripts for analysis, installation, and similar tasks.
- `/vendor` — Application dependencies, such as Go module dependencies.
