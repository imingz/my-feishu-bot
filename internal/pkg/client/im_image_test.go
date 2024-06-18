package client

import (
	"fmt"
	"os"
	"testing"
)

func TestIm_Image_Upload(t *testing.T) {
	file, _ := os.Open("avatar.png")
	imageKey, err := Get().Im_Image_Upload(ImageType_Message, file)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("imageKey: %v\n", imageKey)
}
