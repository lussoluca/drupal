package project

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Term ...
type Term struct {
	Name  string
	Value string
}

// File ...
type File struct {
	URL         string
	ArchiveType string `xml:"archive_type"`
	Md5         string
	Size        int
	filedate    string
}

// Release ...
type Release struct {
	Name         string
	Version      string
	Tag          string
	VersionMajor int    `xml:"version_major"`
	VersionPatch int    `xml:"version_patch"`
	VersionExtra string `xml:"version_extra"`
	Status       string
	ReleaseLink  string `xml:"release_link"`
	DownloadLink string `xml:"download_link"`
	date         string
	mdhash       string
	filesize     int
	Files        []File
}

// Project is a Drupal project.
type Project struct {
	XMLName          xml.Name `xml:"project"`
	Name             string
	ShortName        string
	pack             string
	installedVersion string
	APIVersion       string `xml:"api_version"`
	RecommendedMajor int    `xml:"recommended_major"`
	SupportedMajors  int    `xml:"supported_majors"`
	DefaultMajor     int    `xml:"default_major"`
	ProjectStatus    string `xml:"project_status"`
	Link             string
	Terms            []Term
	Releases         []Release
}

// New return a new Project.
func New(pack string, name string, shortName string, installedVersion string) Project {
	project := Project{
		pack:             pack,
		Name:             name,
		ShortName:        shortName,
		installedVersion: installedVersion,
	}

	project.getReleaseHistory()

	return project
}

// getReleaseHistory retrieves versions from updates.drupal.org.
func (p *Project) getReleaseHistory() error {
	resp, err := http.Get("https://updates.drupal.org/release-history/" + p.ShortName + "/7.x")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	_ = xml.Unmarshal(data, p)

	return nil
}

// IsCore returns true if this is a Core project.
func (p *Project) IsCore() bool {
	return p.pack == "Core"
}
