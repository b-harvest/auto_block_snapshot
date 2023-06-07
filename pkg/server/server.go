package server

import (
	"os/exec"
	"time"

	"auto_block_snapshot/pkg/aws"
	"auto_block_snapshot/pkg/config"

	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run() {
	for {
		// Step 1: Run full node binary
		cmd := exec.Command(s.cfg.FullNode.Path, "start")
		err := cmd.Start()
		if err != nil {
			log.Fatal().Err(err).Msg("An error occurred while starting the full node")
			return
		}
		time.Sleep(6 * time.Hour)

		// Step 3: Kill full node
		cmd.Process.Kill()
		if err != nil {
			log.Fatal().Err(err).Msg("An error occurred while killing the full node")
			return
		}
		// Step 4: Run pruning job
		PrunerCmd := exec.Command(s.cfg.Pruner.Path, "prune", s.cfg.FullNode.Data_Path, "--cosmos-sdk=false") // Assuming cosmprund.Prune() is a function that does the pruning job.
		err = PrunerCmd.Run()
		if err != nil {
			log.Fatal().Err(err).Msg("An error occurred while pruning the full node")
			return
		}
		// Step 5: Compress output
		tarCmd := exec.Command("tar", "-czf", "data.tar.gz", s.cfg.FullNode.Data_Path)
		err = tarCmd.Run()
		if err != nil {
			log.Fatal().Err(err).Msg("An error occurred while compressing the data")
			return
		}

		go func() {
			s3 := aws.NewS3(s.cfg)
			s3.Upload()
		}()

		// Step 6: Restart fullnode
		cmd.Start()
	}
}
