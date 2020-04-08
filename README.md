loggr
===============

This is a library for mix Zap logger to Go "context" package.

## Installing

### *go get*

    $ go get -u github.com/zhs/loggr

## Example

### Mix logger to context

```golang
import (
	"fmt"

	"github.com/zhs/loggr"
)

func main() {
	logger := New("version", "0.1")
	logger = logger.With("addition", "x")
	ctx := ToContext(context.TODO(), logger)
	WithContext(ctx).Info("Test context")
}
```

## Contributing

You are more than welcome to contribute to this project.  Fork and
make a Pull Request, or create an Issue if you see any problem.

## License

BSD 2 Clause license