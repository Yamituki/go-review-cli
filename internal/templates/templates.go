package templates

import _ "embed"

//go:embed hooks/prepare-commit-msg
var PrepareCommitMsgHook string
