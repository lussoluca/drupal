package drupal

// Release ...
type Release struct {
	Name         string `xml:"name"`
	Version      string `xml:"version"`
	Tag          string `xml:"tag"`
	VersionMajor int    `xml:"version_major"`
	VersionPatch int    `xml:"version_patch"`
	VersionExtra string `xml:"version_extra"`
	Status       string `xml:"status"`
	ReleaseLink  string `xml:"release_link"`
	DownloadLink string `xml:"download_link"`
	Date         string `xml:"date"`
	Mdhash       string `xml:"mdhash"`
	Filesize     int    `xml:"filesize"`
	Terms        []Term `xml:"terms>term"`
	Files        []File `xml:"files>file"`
}
