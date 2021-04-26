package finalserver

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func runExperiment(e Experiment) (string, error) {
	if e.ErrorRatio == 0 || rand.Intn(100) >= e.ErrorRatio {
		return hash(e)
	}

	if e.ErrorType == TimeoutError {
		timeoutErr := timeout(e)

		return "", timeoutErr
	}

	if e.ErrorType == UnhandledError {
		hash(e)

		return "", errors.New("unhandled error: experiment successful unhandled error")
	}

	return hash(e)
}

func timeout(e Experiment) error {
	time.Sleep(time.Duration(e.TimeoutLengthInS) * time.Second)

	return errors.New("timeout: experiment successful timeout")
}

func hash(e Experiment) (hashed string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(e.StringToHash), 10)

	if err != nil {
		return e.StringToHash, err
	}

	return string(bytes), nil
}
