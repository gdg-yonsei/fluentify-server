package model

type Scene struct {
	Id             string `firestore:"-"`
	Question       string `firestore:"question"`
	ImageUrl       string `firestore:"image_url"`
	Context        string `firestore:"context"`
	ExpectedAnswer string `firestore:"expected_answer"`
	TopicId        string `firestore:"topic_id"`
}
