package reports

import (
	"testing"

	"github.com/abdullin/omni/core/spec"
)

func TestUseCases(t *testing.T) {
	ctx := spec.NewContext(Spec)
	mod := NewModule(ctx.Pub())
	ctx.Verify(mod).ToTesting(t)

}
