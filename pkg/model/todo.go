package model

// Todo autobind ==> type Todo {
//    id: ID!
//    text: String!
//    done: Boolean!
//    user: User!  need implement user method
// }
type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"user"`
}
