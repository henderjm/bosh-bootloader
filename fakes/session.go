package fakes

import "github.com/aws/aws-sdk-go/service/ec2"

type Session struct{}

func (s *Session) ImportKeyPair(*ec2.ImportKeyPairInput) (*ec2.ImportKeyPairOutput, error) {
	return nil, nil
}
