package model

// FeedCard contains all elements in the cards
type FeedCard struct {
	CardHash string `json:"cardHash"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
