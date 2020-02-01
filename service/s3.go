package service

import (
	"log"
	"os"

	"github.com/Lim79Plus/go-aws/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// svc s3 access instance
var svc *s3.S3

// S3Access entryPoint
func S3Access() {
	log.Println("s3.access()")
	createAccessInstance()
	// bucketList()
	// createBucket()
	// getBucketItems("test")
	uplocadItemToBucket()
}

func createNewS3Session() *session.Session {
	profileName := "go-aws"
	creds := credentials.NewSharedCredentials("", profileName)
	sess, err := session.NewSession(&aws.Config{
		Credentials:      creds,
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("ap-northeast-1"),
		Endpoint:         aws.String(config.Conf.Aws.S3.URL),
	},
	)
	if err != nil {
		log.Fatal(err)
	}
	return sess
}

func createAccessInstance() {
	sess := createNewS3Session()
	svc = s3.New(sess)
}

func bucketList() {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Buckets result: ", result.Buckets[0].Name)

	// for _, b := range result.Buckets {
	// 	log.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	// }
}

func createBucket() {
	bucket := "test"
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Fatalf("Unable to create bucket %q, %v", bucket, err)
	}
}

func getBucketItems(bucket string) {
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
	if err != nil {
		log.Fatalf("Unable to list items in bucket %q, %v", bucket, err)
	}

	for _, item := range resp.Contents {
		log.Println("Name:", *item.Key)
	}
}

func openFicture() func() (*os.File, string) {
	return func() (*os.File, string) {
		filename := "gogo.png"
		// Uploadするファイルの読み込み
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Unable to open file %s, %v", filename, err)
		}
		defer file.Close()
		return file, filename
	}
}

func uplocadItemToBucket() {
	bucket := "s3-local"
	filename := "gogo.png"
	// Uploadするファイルの読み込み
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %s, %v", filename, err)
	}
	defer file.Close()

	// S3のアップロードインスタンスの取得
	sess := createNewS3Session()
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		log.Fatalf("Unable to upload %s to %s, %v", filename, bucket, err)
	}

	log.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}
