package finalserver

import (
	"errors"
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func runExperiment(e Experiment, c *config.ServerConfig, experimentStartTime time.Time) (string, error) {
	if c.Randomized && (c.ErrorRatio > 0 && rand.Intn(100) < c.ErrorRatio) {
		return "", forceError(e, c)
	}

	durationSinceStart := time.Now().Sub(experimentStartTime)
	if !c.Randomized && c.ErrorRatio > 0 {
		if durationSinceStart.Seconds() > float64(c.FailAfterTimeInS) {
			return "", forceError(e, c)
		}
		return hash(e.StringToHash, c.HashSalt)
	}

	return hash(e.StringToHash, c.HashSalt)
}

func forceError(e Experiment, c *config.ServerConfig) error {
	if c.ErrorType == config.TimeoutError {
		timeoutErr := timeout(c)

		return timeoutErr
	}

	hash(e.StringToHash, c.HashSalt-2)

	return errors.New("unhandled error: experiment successful unhandled error")
}

func timeout(c *config.ServerConfig) error {
	time.Sleep(time.Duration(c.TimeoutLengthInS) * time.Second)

	return errors.New("timeout: experiment successful timeout")
}

func hash(s string, cost int) (hashed string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), cost)

	if err != nil {
		return s, err
	}

	return string(bytes), nil
}
