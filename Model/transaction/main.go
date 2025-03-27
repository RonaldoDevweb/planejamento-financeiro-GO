package transaction

import "time"

// Transaction : coment
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Transactions : coment
// Transaction : coment
type Transactions []Transaction
