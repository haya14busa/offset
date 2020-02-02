# offset - get file position from offset

[![Go status](https://github.com/haya14busa/offset/workflows/Go/badge.svg)](https://github.com/haya14busa/offset/actions)
[![GoDoc](https://godoc.org/github.com/haya14busa/offset?status.svg)](https://godoc.org/github.com/haya14busa/offset)

## Usage

```go
import "github.com/haya14busa/offset"
```

[![GoDoc](https://godoc.org/github.com/haya14busa/offset?status.svg)](https://godoc.org/github.com/haya14busa/offset)

## CLI

### Installation

```shell
# Go
$ go get github.com/haya14busa/offset/cmd/offset
```

### offset -h

```
Usage: offset [FLAGS] [File]
        Return position from given offset
Flags:
  -offset int
        byte offset
  -version
        print version

GitHub: https://github.com/haya14busa/offset
```

### Example

```shell
$ offset -offset 14 LICENSE
{"file":"LICENSE","offset":14,"line":3,"column":1}
```
