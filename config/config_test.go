package config

import (
	"os"
	"testing"

	"github.com/go-test/deep"
)

func Test_getWorkDir(t *testing.T) {

	tests := []struct {
		name, env, folderName string
	}{
		{
			"Test_getWorkDir_開発環境",
			"",
			"/go-aws",
		},
		{
			"Test_getWorkDir_テスト環境",
			"test",
			"/go-aws",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			os.Setenv("GO_ENV", test.env)
			setEnv()
			readConfig()
			// When
			got := getWorkDir()
			want := os.Getenv("WRK_DIR") + test.folderName
			// Then
			if got != want {
				t.Errorf("got is not match want. got=%v, want=%s", got, want)
			}
		})
	}
}

func Test_isDev(t *testing.T) {

	tests := []struct {
		name, env string
		want      bool
	}{
		{
			"Test_isDev_開発環境",
			"",
			true,
		},
		{
			"Test_isDev_テスト環境",
			"test",
			true,
		},
		{
			"Test_isDev_開発及びテスト環境以外",
			"staging",
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Given
			env := test.env
			os.Setenv("GO_ENV", env)
			setEnv()
			// When
			got := isDev()
			want := test.want
			// Then
			if got != want {
				t.Errorf("got is not match want. got=%v, want=%v", got, want)
			}
		})
	}
}

func Test_init(t *testing.T) {

	t.Run("環境変数のテスト", func(t *testing.T) {
		// Given
		want := "test_env"
		env := "GO_TEST_ENV"
		os.Setenv(env, want)
		// When
		got := os.Getenv(env)
		// Then
		if got != want {
			t.Errorf("got is not match want. got=%s, want=%s", got, want)
		}
	})

	tests := []struct {
		name, env, want string
	}{
		{
			"init_開発環境",
			"",
			"develop",
		},
		{
			"init_検証環境",
			"staging",
			"staging",
		},
		{
			"init_商用環境",
			"production",
			"production",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			// Given
			os.Setenv("GO_ENV", test.env)
			// When
			got := setEnv()
			want := test.want
			// Then
			if got != want {
				t.Errorf("got is not match want. got=%v, want=%s", got, want)
			}
		})
	}
}

func Test_readConfig(t *testing.T) {

	t.Run("Test_readConfig", func(t *testing.T) {
		// Given
		want := Config{
			App: App{
				Name: "test-go-aws",
				Env:  "test",
			},
			Aws: Aws{
				S3: S3s{
					URL:    "test_s3_url",
					Bucket: "test_s3_bucket",
				},
			},
		}
		Environment = "test"
		// When
		got := readConfig()
		// Then

		if diff := deep.Equal(got, want); diff != nil {
			t.Error(diff)
		}
	})
}
