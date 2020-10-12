// This file is automatically generated. Do not modify it manually.

package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.mattermost.agenda",
  "name": "Agenda",
  "description": "Plugin to handle meeting agendas for Mattermost channels.",
  "homepage_url": "https://github.com/mattermost/mattermost-plugin-agenda",
  "support_url": "https://github.com/mattermost/mattermost-plugin-agenda/issues",
  "release_notes_url": "https://github.com/mattermost/mattermost-plugin-agenda/releases/tag/v0.2.1",
  "version": "0.2.1",
  "min_server_version": "5.26.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": []
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
