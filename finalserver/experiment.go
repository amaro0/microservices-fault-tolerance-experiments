package finalserver

import (
	"errors"
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func runExperiment(e Experiment, c *config.ServerConfig) (string, error) {
	if c.ErrorRatio == 0 || rand.Intn(100) >= c.ErrorRatio {
		return hash(e.StringToHash, c.HashSalt)
	}

	if c.ErrorType == config.TimeoutError {
		timeoutErr := timeout(c)

		return "", timeoutErr
	}

	if c.ErrorType == config.UnhandledError {
		hash(e.StringToHash, c.HashSalt-2)

		return "", errors.New("unhandled error: experiment successful unhandled error")
	}

	return hash(e.StringToHash, c.HashSalt)
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
