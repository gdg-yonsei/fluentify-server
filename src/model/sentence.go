package model

type Sentence struct {
	Id        string `firestore:"-"`
	Text      string `firestore:"text"`
	Tip       string `firestore:"tip"`
	TopicId   string `firestore:"topic_id"`
	VideoPath string `firestore:"video_path"`
}
