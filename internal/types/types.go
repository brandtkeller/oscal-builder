package types

import (
	"time"
)

type OscalComponentDefinition struct {
	UUID       string      `json:"uuid" yaml:"uuid"`
	Metadata   Metadata    `json:"metadata" yaml:"metadata"`
	Components []Component `json:"components" yaml:"components"`
	BackMatter BackMatter  `json:"back-matter" yaml:"back-matter"`
}

type Metadata struct {
	Title        string    `json:"title" yaml:"title"`
	LastModified time.Time `json:"last-modified" yaml:"last-modified"`
	Version      string    `json:"version" yaml:"version"`
	OscalVersion string    `json:"oscal-version" yaml:"oscal-version"`
	Parties      []struct {
		UUID  string `json:"uuid" yaml:"uuid"`
		Type  string `json:"type" yaml:"type"`
		Name  string `json:"name" yaml:"name"`
		Links []struct {
			Href string `json:"href" yaml:"href"`
			Rel  string `json:"rel" yaml:"rel"`
		} `json:"links" yaml:"links"`
	} `json:"parties" yaml:"parties"`
}

type Component struct {
	UUID                   string                  `json:"uuid" yaml:"uuid"`
	Type                   string                  `json:"type" yaml:"type"`
	Title                  string                  `json:"title" yaml:"title"`
	Description            string                  `json:"description" yaml:"description"`
	Purpose                string                  `json:"purpose" yaml:"purpose"`
	ResponsibleRoles       []ResponsibleRole       `json:"responsible-roles" yaml:"responsible-roles"`
	ControlImplementations []ControlImplementation `json:"control-implementations" yaml:"control-implementations"`
}

type ResponsibleRole struct {
	RoleID     string   `json:"role-id" yaml:"role-id"`
	PartyUUIDs []string `json:"party-uuids" yaml:"party-uuids"`
}

type ControlImplementation struct {
	UUID                    string                   `json:"uuid" yaml:"uuid"`
	Source                  string                   `json:"source" yaml:"source"`
	Description             string                   `json:"description" yaml:"description"`
	ImplementedRequirements []ImplementedRequirement `json:"implemented-requirements" yaml:"implemented-requirements"`
}

type ImplementedRequirement struct {
	UUID        string `json:"uuid" yaml:"uuid"`
	ControlID   string `json:"control-id" yaml:"control-id"`
	Description string `json:"description" yaml:"description"`
}

type BackMatter struct {
	Resources []struct {
		UUID   string `json:"uuid" yaml:"uuid"`
		Title  string `json:"title" yaml:"title"`
		Rlinks []struct {
			Href string `json:"href" yaml:"href"`
		} `json:"rlinks" yaml:"rlinks"`
	} `json:"resources" yaml:"resources"`
}
