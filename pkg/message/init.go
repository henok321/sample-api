package message

import "database/sql"

func InitalizeMessageModule(db *sql.DB) MessageService {
	return newMessageService(newMessageRepository(db))
}
