package model

type Topic struct {
	Id           string   `firestore:"-"`
	Title        string   `firestore:"title"`
	ThumbnailUrl string   `firestore:"thumbnail_url"`
	SentenceIds  []string `firestore:"-"`
	SceneIds     []string `firestore:"-"`
}
