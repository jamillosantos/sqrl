package sqrl

type (
	SelectForType        int
	SelectForLockingType int
)

const (
	// ForUpdate marks a use a FOR UPDATE.
	ForUpdate SelectForType = iota
	// ForNoKeyUpdate marks a use a FOR NO KEY UPDATE.
	ForNoKeyUpdate
	// ForShare marks a use a FOR SHARE.
	ForShare
	// ForKeyShare marks a use a FOR KEY SHARE.
	ForKeyShare
)

const (
	// ForUpdateTypeNone will generate the FOR [OPTION] with no modifier.
	ForUpdateTypeNone SelectForLockingType = iota
	// SkipLocked will generate the FOR [OPTION] SKIP LOCKED modifier.
	SkipLocked
	// NoWait will generate the FOR [OPTION] NOWAIT modifier.
	NoWait
)

// IsValid returns whether or not the value is valid.
func (t SelectForType) IsValid() bool {
	switch t {
	case ForUpdate, ForNoKeyUpdate, ForShare, ForKeyShare:
		return true
	default:
		return false
	}
}

// String generate the string for the selectForType.
func (t SelectForType) String() string {
	switch t {
	case ForUpdate:
		return "UPDATE"
	case ForNoKeyUpdate:
		return "NO KEY UPDATE"
	case ForShare:
		return "SHARE"
	case ForKeyShare:
		return "KEY SHARE"
	default:
		return ""
	}
}

// IsValid returns whether or not the value is valid.
func (t SelectForLockingType) IsValid() bool {
	switch t {
	case SkipLocked, NoWait:
		return true
	default:
		return false
	}
}

// String generate the string for the selectForLockingType.
func (t SelectForLockingType) String() string {
	switch t {
	case SkipLocked:
		return "SKIP LOCKED"
	case NoWait:
		return "NOWAIT"
	default:
		return ""
	}
}

type selectFor struct {
	forType     SelectForType
	lockingType SelectForLockingType
}

// ToSql generates the FOR [UPDATE|NO KEY UPDATE|SHARE|KEYSHARE] [SKIP LOCKED|NOWAIT]
func (o *selectFor) ToSql() (string, []interface{}, error) {
	return "FOR " + o.forType.String() + " " + o.lockingType.String(), nil, nil
}
