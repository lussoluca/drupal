package project

// Project a Drupal project.
type Project struct {
	name    string
	project string
	pack    string
	version string
}

// New return a new project.
func New(name string, project string, pack string, version string) Project {
	return Project{
		pack:    pack,
		name:    name,
		project: project,
		version: version,
	}
}

// GetVersions retrieves versions from server.
func (p *Project) GetVersions() string {
	return p.name
}
