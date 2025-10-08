package oci_test

import (
	"ociclean/oci"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		imageRef string
		wantErr  bool
	}{
		{
			name:     "postgres",
			imageRef: "postgres",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := oci.List(tt.imageRef)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("List() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("List() succeeded unexpectedly")
			}
		})
	}
}

func TestListTags(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		src     string
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "postgres tags",
			src:     "postgres",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := oci.ListTags(tt.src)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ListTags() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ListTags() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			// res := slices.Equal(tt.want, got)
			res := reflect.DeepEqual(tt.want, got)
			if !res {
				t.Errorf("ListTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListImageInfo(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		src     string
		want    []oci.ImageInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "envprops",
			src: "ghcr.io/myhops/f12",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := oci.ListImageInfo(tt.src)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ListImageInfo() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ListImageInfo() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if false {
				t.Errorf("ListImageInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
