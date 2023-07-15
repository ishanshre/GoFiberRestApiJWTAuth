package drivers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectSql(t *testing.T) {
	db, err := ConnectSql(dbString, dsn)
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.SQL.Close()
	assert.Equal(t, maxOpenDBConn, db.SQL.Stats().MaxOpenConnections)

	err = db.SQL.Ping()
	assert.NoError(t, err)
}

func TestConnectSql_Failure(t *testing.T) {
	db, err := ConnectSql("postgres", "noDatabase")
	assert.Error(t, err)
	assert.Nil(t, db)

	db, err = ConnectSql("unknown", "nodatabase")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "unknown driver")
	assert.Nil(t, db)
}
