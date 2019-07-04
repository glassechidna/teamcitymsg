package teamcitymsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMsgBuildProblem(t *testing.T) {
	t.Run("with identity", func(t *testing.T) {
		sm := NewMsgBuildProblem("it's an escaped value", "problem.identity")
		assert.Equal(t, "##teamcity[buildProblem description='it|'s an escaped value' identity='problem.identity']", sm.String())
	})

	t.Run("without identity", func(t *testing.T) {
		sm := NewMsgBuildProblem("it's an escaped value", "")
		assert.Equal(t, "##teamcity[buildProblem description='it|'s an escaped value']", sm.String())
	})
}
