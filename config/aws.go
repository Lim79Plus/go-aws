package config

// Aws struct
type Aws struct {
	S3 S3s `yaml:"s3"`
}

// S3s struct
type S3s struct {
	URL    string `yaml:"url"`
	Bucket string `yaml:"bucket"`
}
