package bpembed_test

import (
	"testing"

	"github.com/waldirborbajr/bpembed"
)

func TestNotExpectedContent(t *testing.T) {

	expectedContent := "v1.0.1"

	outputContnt := bpembed.BPEmbed("test.txt")

	if expectedContent != outputContnt {
		t.Logf("Success !")
	} else {
		t.Errorf("Failed! got %s want %s", outputContnt, expectedContent)
	}
}

func TestExpectedContent(t *testing.T) {

	expectedContent := "v0.0.1"

	outputContent := bpembed.BPEmbed("test.txt")

	if expectedContent != outputContent {
		t.Errorf("Failed! got %s want %s", outputContent, expectedContent)
	} else {
		t.Logf("Success !")
	}
}
