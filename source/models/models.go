package models

import (
	"github.com/Mimerel/go-logger-client"
	"time"
)

type Files struct {
	Name      string `yaml:"name,omitempty"`
	FullName  string `yaml:"fullname,omitempty"`
	FullPath  string `yaml:"fullpath,omitempty"`
	Path      string `yaml:"path,omitempty"`
	Extension string `yaml:"extension,omitempty"`
	FileType  string `yaml:"fileType,omitempty"`
}

type Elasticsearch struct {
	Url string `yaml:"url,omitempty"`
}

type Folders struct {
	Origin      string `yaml:"origin,omitempty"`
	Destination string `yaml:"destination,omitempty"`
}

type ConversionParams struct {
	TypeName string   `yaml:"typeName,omitempty"`
	Params   []string `yaml:"params,omitempty"`
}

type Configuration struct {
	Elasticsearch        Elasticsearch      `yaml:"elasticSearch,omitempty"`
	Host                 string             `yaml:"host,omitempty"`
	Folders              []Folders          `yaml:"folders,omitempty"`
	OriginExtensions     []string           `yaml:"originExtensions,omitempty"`
	Ignore               []string           `yaml:"ignore,omitempty"`
	DestinationExtension string             `yaml:"destinationExtensions,omitempty"`
	ConvertedFileFolder  string             `yaml:"convertedFileFolder,omitempty"`
	TemporaryFile        string             `yaml:"temporaryFile,omitempty"`
	ConversionParams     []ConversionParams `yaml:"conversionParams,omitempty"`
	Prowl                string             `yaml:"prowl,omitempty"`
	FromEnd              bool               `yaml:"fromEnd,omitempty"`
	Production           bool               `yaml:"production,omitempty"`
	MinimumFileAge       time.Duration      `yaml:"minimumFileAgeInHours,omitempty"`
	Logger               logs.LogParams
	Database             []Files
}
