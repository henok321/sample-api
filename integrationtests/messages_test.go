package integrationtests

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func TestScores(t *testing.T) {
	tests := map[string]testCase{
		"Get all messages": {
			method:             "GET",
			endpoint:           "/messages",
			expectedStatusCode: 200,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
			},
			expectedBody: readContentFromFile(t, "./test_data/json/get_all_messages.json"),
			setup: func(db *sql.DB) {
				executeSQLFile(t, db, "./test_data/sql/messages.sql")
			},
		},
		"Create new message": {
			method:             "POST",
			endpoint:           "/messages",
			requestBody:        `{"content": "Hello, World!"}`,
			expectedStatusCode: 201,
			expectedHeaders: map[string]string{
				"Content-Type": "application/json",
				"Location":     "/messages/1",
			},
		},
		"Update message": {
			method:             "PUT",
			endpoint:           "/messages/1",
			requestBody:        `{"id": 1, "content": "Hello, Universe!"}`,
			expectedStatusCode: 204,
			setup: func(db *sql.DB) {
				executeSQLFile(t, db, "./test_data/sql/messages.sql")
			},
		},
	}

	dbConn, teardownDatabase := setupTestDatabase(t)
	defer teardownDatabase()

	db, err := sql.Open("postgres", dbConn)

	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}

	defer db.Close()

	runGooseUp(t, db)

	server, teardown := setupTestServer()
	defer teardown(server)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.setup != nil {
				tc.setup(db)
			}
			defer executeSQLFile(t, db, "./test_data/sql/cleanup.sql")
			newTestRequest(t, tc, server)
		})
	}
}
