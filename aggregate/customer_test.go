package aggregate_test

import (
	"errors"
	"testing"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "Empty name validation",
			name:          "",
			expectedError: aggregate.ErrorInvalidPerson,
		},
		{
			test:          "Valid name",
			name:          "Percy Bolmer",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
