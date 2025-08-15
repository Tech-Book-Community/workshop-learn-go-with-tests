package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post

	// 1 遍歷每個檔案
	for _, f := range dir {

		// 2 呼叫 getPost 函式，取得 Post 結構
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}

		// 6 將 Post 結構加入 posts
		posts = append(posts, post)
	}

	// 7 回傳 posts
	return posts, nil
}

func getPost(fileSystem fs.FS, filename string) (Post, error) {
	postFile, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	// 用 bufio.NewScanner 逐行讀取檔案內容
	scanner := bufio.NewScanner(postFile)

	// 將 Scan() 跟 Text() 包裝成一個函式，方便重用
	// 使用 strings.TrimPrefix 去除前綴，這樣更有彈性，未來任何欄位都能重用
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		// 用 readBody() 解析 Body
		Body: readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	// 忽略 --- 分隔線那一行
	scanner.Scan()

	// 用 fmt.Fprintln() 把每一行寫進 buffer，這樣能自動補換行字元
	// 因為 scanner.Text() 拿到的是每一行內容，沒有換行
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	// 最後把最後多出來的換行去掉
	return strings.TrimSuffix(buf.String(), "\n")
}
