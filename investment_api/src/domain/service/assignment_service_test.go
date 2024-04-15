package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignment3000(t *testing.T) {
	// Fixture
	r := NewCreditAssignmentService()
	var expectedx300 int32 = 6
	var expectedx500 int32 = 1
	var expectedx700 int32 = 1
	var investment int32 = 3000

	// Act
	x300, x500, x700, err := r.Assign(investment)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedx300, x300)
	assert.Equal(t, expectedx500, x500)
	assert.Equal(t, expectedx700, x700)
}

func TestAssignmentFail(t *testing.T) {
	// Fixture
	r := NewCreditAssignmentService()
	var expectedx300 int32 = 0
	var expectedx500 int32 = 0
	var expectedx700 int32 = 0
	var investment int32 = 4

	// Act
	x300, x500, x700, err := r.Assign(investment)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedx300, x300)
	assert.Equal(t, expectedx500, x500)
	assert.Equal(t, expectedx700, x700)
}
