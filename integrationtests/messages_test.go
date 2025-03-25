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
			assertDatabaseState: func(t *testing.T, db *sql.DB) {
				var content string
				err := db.QueryRow("SELECT content FROM messages WHERE id = 1").Scan(&content)
				if err != nil {
					t.Fatalf("Failed to query database: %v", err)
				}

				if content != "Hello, World!" {
					t.Fatalf("Expected content to be 'Hello, World!', got %s", content)
				}
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
			assertDatabaseState: func(t *testing.T, db *sql.DB) {
				var content string
				err := db.QueryRow("SELECT content FROM messages WHERE id = 1").Scan(&content)
				if err != nil {
					t.Fatalf("Failed to query database: %v", err)
				}

				if content != "Hello, Universe!" {
					t.Fatalf("Expected content to be 'Hello, Universe!', got %s", content)
				}
			},
		},
		"Delete message": {
			method:             "DELETE",
			endpoint:           "/messages/1",
			expectedStatusCode: 204,
			setup: func(db *sql.DB) {
				executeSQLFile(t, db, "./test_data/sql/messages.sql")
			},
			assertDatabaseState: func(t *testing.T, db *sql.DB) {
				var count int
				err := db.QueryRow("SELECT COUNT(*) FROM messages WHERE id = 1").Scan(&count)
				if err != nil {
					t.Fatalf("Failed to query database: %v", err)
				}

				if count != 0 {
					t.Fatalf("Expected 0 messages, got %d", count)
				}
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
			if tc.assertDatabaseState != nil {
				tc.assertDatabaseState(t, db)
			}
		})
	}
}
