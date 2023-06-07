package aws

import (
	"context"
	"os"
	"time"

	"auto_block_snapshot/pkg/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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
	awsCfg, err := aws_config.LoadDefaultConfig(context.Background(), aws_config.WithRegion(s.cfg.Aws.Region))
	if err != nil {
		log.Error().Err(err).Msg("unable to load SDK config")
	}

	client := s3.NewFromConfig(awsCfg)

	f, err := os.Open("data.tar.gz")
	if err != nil {
		log.Error().Err(err).Msg("failed to open file")
	}
	defer f.Close()
	currentTime := time.Now()

	filepath_name := currentTime.String() + s.cfg.FullNode.Chain_Name + "_" + "data.tar.gz"
	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(s.cfg.Aws.Bucket),
		Key:    aws.String(filepath_name),
		Body:   f,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to upload file")
	}
}
