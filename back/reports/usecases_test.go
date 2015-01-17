package reports

import (
	"testing"

	"bitbucket.org/abdullin/proto/back/spec"
)

func TestUseCases(t *testing.T) {
	ctx := spec.NewContext(Spec)
	mod := NewModule(ctx.Pub())
	ctx.Verify(mod).Report(t)

}
