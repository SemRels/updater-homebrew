# updater-homebrew

Updates a Homebrew formula for the new release.

This plugin is distributed as the standalone Go binary `semrel-plugin-updater-homebrew`. Semrel executes the binary as a subprocess, provides plugin configuration through `SEMREL_PLUGIN_*` environment variables, provides release context through `SEMREL_*` environment variables, reads standard output, and treats exit code `0` as success and any non-zero exit code as failure. Install the binary in `~/.semrel/plugins/` or anywhere on your `$PATH`.

## Installation

```bash
go install github.com/SemRels/updater-homebrew/cmd/plugin@latest
```

## Configuration

```yaml
plugins:
  - name: updater-homebrew
    path: ~/.semrel/plugins/semrel-plugin-updater-homebrew
    env:
      SEMREL_PLUGIN_FORMULA_FILE: "Formula/my-tool.rb"
      SEMREL_PLUGIN_URL_TEMPLATE: "https://github.com/acme/tool/archive/refs/tags/v{{ .Version }}.tar.gz"
      SEMREL_PLUGIN_SHA256: "${TARBALL_SHA256}"
```

## `SEMREL_PLUGIN_*` variables

| Name | Required | Description | Default |
| --- | --- | --- | --- |
| `SEMREL_PLUGIN_FORMULA_FILE` | Required | Path to the Homebrew formula file to update. | None |
| `SEMREL_PLUGIN_URL_TEMPLATE` | Optional | Template used to build the download URL for the new release. | None |
| `SEMREL_PLUGIN_SHA256` | Optional | SHA256 checksum to write into the formula. | None |

## `SEMREL_*` release context used

| Variable | Description |
| --- | --- |
| `SEMREL_VERSION` | Resolved release version for the current run. |
| `SEMREL_NEXT_VERSION` | Next version computed by semrel for the release. |
| `SEMREL_DRY_RUN` | Whether semrel is running in dry-run mode. |

## Example behavior

The plugin updates the formula version metadata and, when configured, also refreshes the tarball URL and SHA256 checksum.

## License

Apache-2.0
