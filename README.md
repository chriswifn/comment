# 🌳 Go Bonzai Comment Filter Branch

Comment tool.

## Installation

This comment command can be installed as a standalone program or composed into a Bonzai command tree.

Standalone
```
go install github.com/chriswifn/comment/cmd/comment@latest
```

Composed

```go
package z

import (
    Z "github.com/rwxrob/bonzai/z"
    "github.com/chriswifn/comment"
)

var Cmd = &Z.Cmd{
    Name: `z`,
    Commands: []*Z.Cmd{help.Cmd, comment.Cmd},
}
```

## Tab Completion
To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C comment comment
```

