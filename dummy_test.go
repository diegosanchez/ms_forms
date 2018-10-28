package ms_forms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_dummy_case(t *testing.T) {
	assert.Equal(t, NewDummy(), NewDummy())
}

func Test_dummy_case_is_abc(t *testing.T) {
	assert.Equal(t, true, NewDummy().isABC())
}
