package drupal

// File ...
type File struct {
	URL         string `xml:"url"`
	ArchiveType string `xml:"archive_type"`
	Md5         string `xml:"md5"`
	Size        int    `xml:"size"`
	Filedate    string `xml:"filedate"`
}
