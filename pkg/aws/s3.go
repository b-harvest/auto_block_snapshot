package aws

import (
	"os"
	"time"

	"auto_block_snapshot/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/rs/zerolog/log"
)

type S3 struct {
	cfg *config.Config
}

func NewS3(cfg *config.Config) *S3 {
	return &S3{cfg: cfg}
}

func (s *S3) Upload() {
	// Step 7: Upload to AWS S3
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.cfg.Aws.Region),
	})
	if err != nil {
		log.Error().Err(err).Msg("An error occurred while creating a new session")
	}
	file, err := os.Open("data.tar.gz")
	if err != nil {
		log.Error().Err(err).Msg("An error occurred while Open a data.tar.gz")
	}
	currentTime := time.Now()

	filepath_name := currentTime.String() + s.cfg.FullNode.Chain_Name + "_" + "data.tar.gz"
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.cfg.Aws.Bucket),
		Key:    aws.String(filepath_name),
		Body:   file,
	})
	if err != nil {
		log.Error().Err(err).Msg("An error occurred while Upload a data.tar.gz")
	}
}
