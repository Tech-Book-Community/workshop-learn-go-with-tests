package murkyblogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

// 最後要讀到這個
type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

// 從資料夾讀取
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	// fs.FS 就是有 Open

	// fs.ReadDir 是 Go 標準函式，會回傳該目錄下的所有檔案
	dirs, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0, len(dirs))

	for _, f := range dirs {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)

	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	post := Post{
		Title:       readLine(titleSeparator),
		Description: readLine(descriptionSeparator),
		Tags:        strings.Split(readLine(tagSeparator), ", "),
		Body:        readBody(scanner),
	}

	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	// 把 ----- 讀掉
	scanner.Scan()

	var sb strings.Builder

	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteString("\n")
	}

	return strings.TrimSuffix(sb.String(), "\n")
}
