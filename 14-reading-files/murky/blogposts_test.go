package murkyblogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/Tech-Book-Community/workshop-learn-go-with-tests/14-reading-files/murky"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)

		fs := fstest.MapFS{
			"hello world.md":  &fstest.MapFile{Data: []byte(firstBody)},
			"hello world2.md": &fstest.MapFile{Data: []byte(secondBody)},
		}

		// 直接讀取 MapFS 提供的檔案，不用真的建立檔案或資料夾
		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		// 驗證函式輸出 posts 的數量要和你的 MapFS 檔案數一致
		assertPostsLength(t, posts, fs)

		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}
		assertPost(t, posts[0], want)
	})

	t.Run("error", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Errorf("got nil error")
		}
	})
}

// 會丟失敗

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func assertPostsLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
