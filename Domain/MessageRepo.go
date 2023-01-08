package Domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// NOTE: PULIC METHODS RE ALL CAPS
type MessageRepository interface {
	FindAll() ([]Message, error)
}

// We wrap the sql db stuff in  a struct CustomerRepoDB
type MessageRepoDB struct {
	db *sql.DB
}

// We implement the interface CustomerRepository
// By using the struct CustomerRepoDB (the real implementation)
func (ch MessageRepoDB) FindAll() ([]Message, error) {
	findall_sql := "SELECT * FROM message"
	rows, err := ch.db.Query(findall_sql)

	if err != nil {
		log.Println("error executing sql")
	}

	messages := make([]Message, 0)
	for rows.Next() {
		var message Message
		err = rows.Scan(&message.MessageID, &message.Content)
		if err != nil {
			log.Println("Error scanning rows" + err.Error())
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func NewMessageRepositoryDB() MessageRepoDB {

	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/message")

	if err != nil {
		panic(err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	return MessageRepoDB{db}
}
