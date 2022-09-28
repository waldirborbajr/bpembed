package bpembed_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/waldirborbajr/bpembed"
)

func TestNotExpectedContent(t *testing.T) {

	expectedContent := "v1.0.1"

	outputContent := bpembed.BPEmbed("test.txt")

	assert.NotEqual(t, expectedContent, outputContent, "The two contents should not be the same.")

}

func TestExpectedContent(t *testing.T) {

	expectedContent := "v0.0.1"

	outputContent := bpembed.BPEmbed("test.txt")

	assert.Equal(t, expectedContent, outputContent, "The two contents should be the same.")

}

func TestFileNotFound(t *testing.T) {

	expectedContent := "NotExists"

	outputContent := bpembed.BPEmbed("nonexists.yaml")

	assert.Contains(t, expectedContent, outputContent, "The file should not exist.")

}
