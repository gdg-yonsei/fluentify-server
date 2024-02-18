package model

type Sentence struct {
	Id      string `firestore:"-"`
	Text    string `firestore:"text"`
	TopicId string `firestore:"topic_id"`
}
