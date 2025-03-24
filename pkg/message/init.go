package message

import "database/sql"

func InitalizeMessageModule(db *sql.DB) Service {
	return newMessageService(newMessageRepository(db))
}
