package model

type Scene struct {
	Id       string `firestore:"-"`
	Question string `firestore:"question"`
	ImageUrl string `firestore:"image_url"`
	TopicId  string `firestore:"topic_id"`
}
