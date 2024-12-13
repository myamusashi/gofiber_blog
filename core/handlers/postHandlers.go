package handlers

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/gofiber/fiber/v2"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"

	"github.com/myamusashi/go_blog/core/utils"
)

// Handler untuk route '/posts/{slug}' dan akan render slug berdasarkan parser YAML markdown.
// PostHandler akan mengembalikan sebuah fiber handler function yang akan memproses request
// Untuk single post berdasarkan slug, dan akan parsing markdown frontmatter, mengubah markdown menjadi HTML, dan
// Render post dengan fiber template
func PostHandler(sl utils.SlugRender) fiber.Handler {
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)

	return func(c *fiber.Ctx) error {
		slug := c.Params("slug")

		postMarkdown, err := sl.Read(slug)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "No post here")
		}

		var post utils.PostData
		remainingMd, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error parsing frontmatter")
		}

		var buf bytes.Buffer
		err = mdRenderer.Convert([]byte(remainingMd), &buf)
		if err != nil {
			fiber.NewError(fiber.StatusInternalServerError, "Error rendering markdown to html")
		}

		post.Content = template.HTML(buf.String())

		return c.Render("post", fiber.Map{
			"Title":                   post.Title,
			"Author":                  post.Author,
			"Date":                    post.Date,
			"Content":                 post.Content,
			"Description":             post.Description,
			"Tags":                    post.Tags,
			"MetaDescription":         post.MetaDescription,
			"MetaPropertyTitle":       post.MetaPropertyTitle,
			"MetaPropertyDescription": post.MetaPropertyDescription,
			"MetaOgURL":               post.MetaOgURL,
		})
	}
}
