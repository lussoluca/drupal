package drupal

import "fmt"

// Term ...
type Term struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

// String format a Term.
func (t Term) String() string {
	return fmt.Sprintf("%s : %s | ", t.Name, t.Value)
}
