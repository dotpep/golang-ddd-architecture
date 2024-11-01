package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a valueobject because
// it has no identifier and is immutable.
type Transaction struct {
	// camelCase because it can not be modified,
	// Immutable.
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
