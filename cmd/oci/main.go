package main

import (
	"flag"
	"fmt"
	"io"
	"ociclean/oci"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type options struct {
	registry string
}

func getOptions(args []string) (*options, error) {
	opts := new(options)
	fs := flag.NewFlagSet(args[0], flag.ExitOnError)
	fs.StringVar(&opts.registry, "registry", "", "image registryname, tag is ignored")

	err := fs.Parse(args[1:])
	if err != nil {
		return nil, err
	}
	return opts, nil
}

func printImages(w io.Writer, images []oci.ImageInfo) {
	for _, img := range images {
		fmt.Fprintf(w, "%s:%s created: %s\n", img.RepoName, img.Tag, img.CreatedString)
	}
}

func run(args []string) error {
	opts, err := getOptions(args)
	if err != nil {
		return err
	}

	// create the key chain with docker file creds
	kc := authn.DefaultKeychain
	images, err := oci.ListImageInfo(opts.registry, oci.WithRemoteOption(remote.WithAuthFromKeychain(kc)))
	if err != nil {
		return err
	}

	// Print the image info
	printImages(os.Stdout, images)
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		fmt.Printf("error: %s\n ", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

