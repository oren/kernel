// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package s3iface_test

import (
	"testing"

	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/s3"
	"github.com/convox/kernel/Godeps/_workspace/src/github.com/awslabs/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*s3iface.S3API)(nil), s3.New(nil))
}
