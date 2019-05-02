package model

type Track struct {
	Title    string   `json:"title"  bson:"title"`
	Artist   string   `json:"artist" bson:"artists"`
	Link     string   `json:"link" bson:"link"`
	Quality  string   `json:"quality" bson:"quality"`
	Duration int      `json:"duration" bson:"duration"`
	Sources  []Source `json:"sources" bson:"sources"`
}
