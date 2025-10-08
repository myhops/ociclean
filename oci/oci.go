package oci

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type options struct {
	nameOptions []name.Option
	remoteOptions []remote.Option
}

type option func(*options)

func newOptions(opts []option) *options {
	o := &options{}

	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (opts *options) NameOptions() []name.Option {
	return opts.nameOptions
}

func (opts *options) RemoteOptions() []remote.Option {
	return opts.remoteOptions
}

func WithNameOption(opt name.Option) option  {
	return func(o *options) {
		o.nameOptions = append(o.nameOptions, opt)
	}
}

func WithRemoteOption(opt remote.Option) option  {
	return func(o *options) {
		o.remoteOptions = append(o.remoteOptions, opt)
	}
}

type ImageInfo struct {
	RepoName      string
	Tag           string
	Created       time.Time
	CreatedString string
}

func ImageName(name string) string {
	parts := strings.Split(name, ":")

	if len(parts) == 0 {
		return ""
	}
	return parts[0]
}

func List(imageRef string) error {
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

// ListTags returns the tags of the image.
func ListTags(src string, opts ...option) ([]string, error) {
	oo := newOptions(opts)

	imgName := ImageName(src)
	if imgName == "" {
		return nil, errors.New("invalid image name")
	}
	// Get the repo for the image source
	repo, err := name.NewRepository(src, oo.NameOptions()...)
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %w", err)
	}
	return remote.List(repo)
}

func ListImageInfo(src string, opts ...option) ([]ImageInfo, error) {
	oo := newOptions(opts)

	imageName := ImageName(src)
	// Get the tags
	tags, err := ListTags(src, opts...)
	if err != nil {
		return nil, fmt.Errorf("error getting tags: %w", err)
	}
	res := make([]ImageInfo, 0, len(tags))
	for _, tag := range tags {
		// Get the image ref
		ref, err := name.ParseReference(imageName + ":" + tag, oo.NameOptions()...)
		if err != nil {
			return nil, fmt.Errorf("error parsing %s:%s: %w", imageName, tag, err)
		}
		// Get the image metadata
		img, err := remote.Image(ref, oo.RemoteOptions()...)
		if err != nil {
			continue
			// return nil, fmt.Errorf("error getting image: %w", err)
		}
		// Get the creation date
		cf, err := img.ConfigFile()
		if err != nil {
			continue
			// return nil, fmt.Errorf("error getting config file: %w", err)
		}
		res = append(res, ImageInfo{
			RepoName:      imageName,
			Tag:           tag,
			Created:       cf.Created.Time,
			CreatedString: cf.Created.Time.String(),
		})
	}
	return res, nil
}
