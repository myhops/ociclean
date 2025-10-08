package oci

import (
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type Image struct {
	// Config
}

func ListTags(imageRef string) error {
	ref, err := name.ParseReference(imageRef)
	if err != nil {
		return err
	}
	img, err := remote.Image(ref)
	if err != nil {
		return err
	}
	configFile, err := img.ConfigFile()
	if err != nil {
		return err
	}
	_ = configFile
	return nil
}
