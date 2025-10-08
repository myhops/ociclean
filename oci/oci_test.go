package oci_test

import(
	"ociclean/oci"
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
			name: "postgres",
			imageRef: "postgres",
			wantErr: false,
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
