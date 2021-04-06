package main

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func RunExperiment(e Experiment) (string, error) {
	if e.ErrorType == ExperimentErrors.TimeoutError {
		if e.ErrorRatio != 0 && rand.Intn(100) <= e.ErrorRatio {
			time.Sleep(time.Duration(e.TimeoutLengthInS) * time.Second)
		}

		return hash(e.StringToHash)
	}

	return hash(e.StringToHash)
}

func hash(s string) (hashed string, e error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 9)

	if err != nil {
		return s, err
	}

	return string(bytes), nil
}
