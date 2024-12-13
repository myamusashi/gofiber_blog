package markdown

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/myamusashi/go_blog/core/utils"
	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
)

func GetPostTag(tag string) ([]utils.PostData, error) {
	posts, err := LoadAllMarkdown("markdown/")
	if err != nil {
		return nil, err
	}

	var filteredPosts []utils.PostData
	for _, post := range posts {
		for _, t := range post.Tags {
			if t == tag {
				filteredPosts = append(filteredPosts, post)
				break
			}
		}
	}

	if len(filteredPosts) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Doesn't recognize tag name! or tag not found")
	}

	return filteredPosts, nil
}

func LoadAllMarkdown(dir string) ([]utils.PostData, error) {
	md := goldmark.New()
	var posts []utils.PostData

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var postData utils.PostData
			var buf bytes.Buffer

			split := strings.SplitN(string(content), "\n---\n", 2)
			if len(split) > 1 {
				err = yaml.Unmarshal([]byte(split[0]), &postData)
				if err != nil {
					return err
				}

				err = md.Convert([]byte(split[1]), &buf)
				if err != nil {
					return err
				}

				postData.Content = template.HTML(buf.String())
			} else {
				err = md.Convert(content, &buf)
				if err != nil {
					return err
				}

				postData.Content = template.HTML(buf.String())
			}

			posts = append(posts, postData)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return posts, nil
}
