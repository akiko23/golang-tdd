package blogrender_test

import (
	"bytes"
	"testing"

	blogrender "github.com/akiko23/golang-tdd/blogrender"
)

func TestRender(t *testing.T) {
	var (
		somePost = blogrender.Post{
			Title:       "hello world",
			Body:        "This is a sample post",
			Description: "some description",
			Tags:        []string{"go", "tdd", "it"},
		}
	)

	buf := bytes.Buffer{}

	err := blogrender.Render(&buf, somePost)
	Handle(t, err)

	got := buf.String()
	want := "<h1>Hello world</h1>"

	if got != want {
		t.Errorf("Want %s but got %s", want, got)
	}
}

func Handle(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
