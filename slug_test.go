package slug

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSlugify(t *testing.T) {
	assert.Equal(t, "foo-bar", Slugify("foo bar", "", ""))
	assert.Equal(t, "foo-bar", Slugify("foo ! bar", "", ""))
	assert.Equal(t, "foo-bar1232", Slugify("@!#!@#!@#foo bar1232", "", ""))
	assert.Equal(t, "f-bar", Slugify("foo bar", "o", ""))
	assert.Equal(t, "foo", Slugify("   foo    ---", "", ""))
}
