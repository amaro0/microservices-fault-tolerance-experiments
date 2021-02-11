package main

import "golang.org/x/crypto/bcrypt"

func NewExperiment() {

}

func RunExperiment(e Experiment) (string, error) {
	return hash(e.StringToHash)
}

func hash(s string) (hashed string, e error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 9)

	if err != nil {
		return s, err
	}

	return string(bytes), nil
}
