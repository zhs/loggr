package loggr

import (
	"context"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	logger := New("version", "0.1")
	logger = logger.With("addition", "x")
	logger.Info("Test")
}

func TestToContext(t *testing.T) {
	logger := New("version", "0.1")
	logger = logger.With("addition", "x")
	ctx := ToContext(context.TODO(), logger)
	WithContext(ctx).Info("Test context")
}
