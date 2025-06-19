package storage

import (
	"database/sql"
	"go-chat/models"
)

var DB *sql.DB

func SaveMessage(groupID, senderID, receiverID int, message string) error {
	_, err := DB.Exec(
		context.Background(),
		`INSERT INTO messages (group_id, sender_id, receiver_id, message)
		 VALUES ($1, $2, $3, $4)`,
		groupID, senderID, receiverID, message,
	)
	return err
}


func GetMessagesByGroup(groupID, limit, offset int) ([]models.Message, error) {
	rows, err := DB.Query(
		context.Background(),
		`SELECT id, group_id, sender_id, receiver_id, message, created_at
		 FROM messages
		 WHERE group_id = $1
		 ORDER BY created_at ASC
		 LIMIT $2 OFFSET $3`,
		groupID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		err := rows.Scan(&msg.ID, &msg.GroupID, &msg.SenderID, &msg.ReceiverID, &msg.Content, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
