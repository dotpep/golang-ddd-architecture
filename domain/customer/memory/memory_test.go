package memory

import (
	"errors"
	"testing"

	"github.com/dotpep/golang-ddd-architecture/aggregate"
	"github.com/dotpep/golang-ddd-architecture/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}

	cust, err := aggregate.NewCustomer("heimer")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryCustomerRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:          "no customer by id",
			id:            uuid.MustParse("4d95e388-cd13-49f2-867d-a886482c7d99"),
			expectedError: customer.ErrorCustomerNotFound,
		},
		{
			name:          "customer by id",
			id:            id,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
