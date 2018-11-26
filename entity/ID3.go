package entity

type ID3 struct {
	Title           string `json:"title" bson:"title"`
	Artist          string `json:"artist" bson:"artist"`
	Album           string `json:"album" bson:"album"`
	AlbumArtist     string `json:"albumArtist" bson:"albumArtist"`
	Composer        string `json:"composer" bson:"composer"`
	Year            int    `json:"year" bson:"year"`
	Genre           string `json:"genre"`
	Comment         string `json:"comments" bson:"comments"`
	Lyrics          string `json:"lyric" bson:"lyric"`
	TrackNumber     int    `json:"trackNumber" bson:"trackNumber"`
	TrackTotal      int    `json:"totalTrack" bson:"totalTrack"`
	PictureMIMEType string `json:"pictureMIMEType" bson:"pictureMIMEType"`
	PictureData     string `json:"pictureData" bson:"pictureData"`
}
