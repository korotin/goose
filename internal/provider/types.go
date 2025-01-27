package provider

import (
	"fmt"
	"time"
)

// MigrationType is the type of migration.
type MigrationType int

const (
	TypeGo MigrationType = iota + 1
	TypeSQL
)

func (t MigrationType) String() string {
	switch t {
	case TypeGo:
		return "go"
	case TypeSQL:
		return "sql"
	default:
		// This should never happen.
		return fmt.Sprintf("unknown (%d)", t)
	}
}

// Source represents a single migration source.
//
// The Path field may be empty if the migration was registered manually. This is typically the case
// for Go migrations registered using the [WithGoMigration] option.
type Source struct {
	Type    MigrationType
	Path    string
	Version int64
}

// MigrationResult is the result of a single migration operation.
type MigrationResult struct {
	Source    Source
	Duration  time.Duration
	Direction string
	// Empty indicates no action was taken during the migration, but it was still versioned. For
	// SQL, it means no statements; for Go, it's a nil function.
	Empty bool
	// Error is only set if the migration failed.
	Error error
}

// State represents the state of a migration.
type State string

const (
	// StatePending is a migration that exists on the filesystem, but not in the database.
	StatePending State = "pending"
	// StateApplied is a migration that has been applied to the database and exists on the
	// filesystem.
	StateApplied State = "applied"

	// TODO(mf): we could also add a third state for untracked migrations. This would be useful for
	// migrations that were manually applied to the database, but not versioned. Or the Source was
	// deleted, but the migration still exists in the database. StateUntracked State = "untracked"
)

// MigrationStatus represents the status of a single migration.
type MigrationStatus struct {
	Source    Source
	State     State
	AppliedAt time.Time
}
