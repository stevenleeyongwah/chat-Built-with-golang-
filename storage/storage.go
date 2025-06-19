package storage

import (
	"database/sql"
	"go-chat/models"
)

var DB *sql.DB

func SaveMessage(senderID, receiverID int, message string) error {
	_, err := DB.Exec(
		`INSERT INTO messages (sender_id, receiver_id, message) VALUES ($1, $2, $3)`,
		senderID, receiverID, message,
	)
	return err
}

func GetMessages(senderID, receiverID int) ([]models.Message, error) {
	rows, err := DB.Query(
		context.Background(),
		`SELECT id, sender_id, receiver_id, message, created_at
		 FROM messages
		 WHERE (sender_id = $1 AND receiver_id = $2)
		    OR (sender_id = $2 AND receiver_id = $1)
		 ORDER BY created_at ASC
		 LIMIT 100`, senderID, receiverID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}