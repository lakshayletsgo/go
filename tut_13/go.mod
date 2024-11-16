module tut_13

replace (
	github.com/lakshayletsgo/go/sample v1.0.0 => ../sample
	github.com/lakshayletsgo/go/sample2 v1.0.0 => ../sample2
)

require (
	github.com/lakshayletsgo/go/sample v1.0.0
	github.com/lakshayletsgo/go/sample2 v1.0.0
)

go 1.23.2
