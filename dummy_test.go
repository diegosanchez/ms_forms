package ms_forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dummy_case(t *testing.T) {
	assert.Equal(t, NewDummy(), NewDummy())
}
