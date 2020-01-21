package config

// Aws struct
type Aws struct {
	S3         S3s        `yaml:"s3"`
	Credential Credential `yaml:"credential"`
}

// Credential struct
type Credential struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

// S3s struct
type S3s struct {
	URL    string `yaml:"url"`
	Bucket string `yaml:"bucket"`
}
