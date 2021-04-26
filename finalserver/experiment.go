package finalserver

import (
	"errors"
	"github.com/amaro0/microservices-fault-tolerance-experiments/finalserver/config"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func runExperiment(e Experiment, config *config.ServerConfig) (string, error) {
	if e.ErrorRatio == 0 || rand.Intn(100) >= e.ErrorRatio {
		return hash(e.StringToHash, config.HashSalt)
	}

	if e.ErrorType == TimeoutError {
		timeoutErr := timeout(e)

		return "", timeoutErr
	}

	if e.ErrorType == UnhandledError {
		hash(e.StringToHash, config.HashSalt-2)

		return "", errors.New("unhandled error: experiment successful unhandled error")
	}

	return hash(e.StringToHash, config.HashSalt)
}

func timeout(e Experiment) error {
	time.Sleep(time.Duration(e.TimeoutLengthInS) * time.Second)

	return errors.New("timeout: experiment successful timeout")
}

func hash(s string, cost int) (hashed string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), cost)

	if err != nil {
		return s, err
	}

	return string(bytes), nil
}
