package navicat_test

import (
    "testing"

    "github.com/camry/navicat-crypt"
    "github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
    e := navicat.Encrypt("navicat123456")
    assert.Equal(t, e, "884CD1AE860020A982B9FA1F56EF635E")
}

func TestDecrypt(t *testing.T) {
    s := navicat.Decrypt("884CD1AE860020A982B9FA1F56EF635E")
    assert.Equal(t, s, "navicat123456")
}
