package sqrl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForUpdateSqlGeneration(t *testing.T) {
	sf := &selectFor{
		forType: ForUpdate,
	}
	sql, _, _ := sf.ToSql()
	assert.Equal(t, "FOR UPDATE ", sql)

	sf.forType = ForNoKeyUpdate
	sql, _, _ = sf.ToSql()
	assert.Equal(t, "FOR NO KEY UPDATE ", sql)

	sf.forType = ForShare
	sql, _, _ = sf.ToSql()
	assert.Equal(t, "FOR SHARE ", sql)

	sf.forType = ForKeyShare
	sql, _, _ = sf.ToSql()
	assert.Equal(t, "FOR KEY SHARE ", sql)

	sf.forType = ForUpdate
	sf.lockingType = NoWait
	sql, _, _ = sf.ToSql()
	assert.Equal(t, "FOR UPDATE NOWAIT", sql)

	sf.forType = ForUpdate
	sf.lockingType = SkipLocked
	sql, _, _ = sf.ToSql()
	assert.Equal(t, "FOR UPDATE SKIP LOCKED", sql)
}
