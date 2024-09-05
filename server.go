/*
Dokumentasi Aplikasi Go dengan framework Fiber dan menggunakan Goldmark untuk konversi markdown ke HTML.
Aplikasi ini menggunakan beberapa middleware seperti caching dan compression untuk meningkatkan performa
dari website 

Fitur yang diguanakan disini adalah:

1. Serving static assets, seperti css, js, gambar, dll
2. Caching dari website ini akan menyimpan data user selama 15 menit
3. Kompresi dari website ini menggunakan built-in fiber framework dengan level kompresi
'best speed' untuk mempercepat delivery website ke client
4. Loading dan parsing markdown , konversi markdown ke HTML, dan rendering semua HTML menggunakan Fiber template
5. Posts yang akan ditampilkan di route '/' di urutkan berdasarkan waktu, dari yang paling baru (paling atas) sampai
paling lama (paling bawah).
*/
package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/html/v2"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"gopkg.in/yaml.v2"
)

func main() {
    // Render semua template '.html' yang ada di dir 'templates'
	Engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: Engine,
	})
    
    // Cache file selama 15 menit
	app.Use(cache.New(cache.Config{
		Expiration:   15 * time.Minute,
		CacheControl: true,
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

    // Include semua file yang ada di dir 'static' ke global
	app.Static("/static", "./static")
    
    // Render konten markdown berdasarkan slug di markdown tersebut
	app.Get("/posts/:slug", PostHandler(FileReader{}))
    
    // Render tampilan utama dari website
	app.Get("/", func(c *fiber.Ctx) error {
		posts, err := loadMarkdownPosts("./markdown")
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "No markdown here boss!")
		}

		// Mengurutkan postingan dari yang paling baru sampai paling lama
		sort.Slice(posts, func(i, j int) bool {
			dateFormat := "2006-01-02 15:04"
			dateI, _ := time.Parse(dateFormat, posts[i].Date)
			dateJ, _ := time.Parse(dateFormat, posts[j].Date)
			return dateI.After(dateJ)
		})

		return c.Render("index", fiber.Map{
			"Posts": posts,
		})
	})
	log.Fatal(app.Listen(":8000"))
}

// Struktur data untuk postingan yang akan ditampilkan berisi metadata, informasi author, dan 
// konten markdown yang sudah diubah menjadi HTML
type PostData struct {
	Title                   string `yaml:"Title"`
	Slug                    string `yaml:"Slug"`
	Date                    string `yaml:"Date"`
	Description             string `yaml:"Description"`
	MetaDescription         string `yaml:"MetaDescription"`
	MetaPropertyTitle       string `yaml:"MetaPropertyTitle"`
	MetaPropertyDescription string `yaml:"MetaPropertyDescription"`
	MetaOgURL               string `yaml:"MetaOgURL"`
	Author                  Author `yaml:"author"`
	Content                 template.HTML
}

type Author struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

// Handler untuk route '/posts/' berdasarkan slug.
// PostHandler akan mengembalikan sebuah fiber handler function yang akan memproses request
// untuk single post berdasarkan slug, dan akan parsing markdown frontmatter, mengubah markdown menjadi HTML, dan
// Merender post dengan fiber template
func PostHandler(sl SlugRender) fiber.Handler {
	// Konfigurasi markdown dengan extension yang sama dengan markdown
    // yang ada di Github, dan highlighting dracula
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
		
        // Mengecek slug 'markdown' dan jika tidak ada akan return error 'not found'
        postMarkdown, err := sl.Read(slug)
		if err != nil {
			return fiber.NewError(fiber.StatusNotFound, "No post here")
		}
        
		var post PostData
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
			"MetaDescription":         post.MetaDescription,
			"MetaPropertyTitle":       post.MetaPropertyTitle,
			"MetaPropertyDescription": post.MetaPropertyDescription,
			"MetaOgURL":               post.MetaOgURL,
		})
	}
}

// Interface untuk membaca isi konten markdown file 
// tergantung pada slug dari markdown tersebut
type SlugRender interface {
	Read(slug string) (string, error)
}

// FileReader mengimplematasikan SlugRender interface 
// Ini akan membaca files dari filesystem dan tergantung slug yang diberikan 
type FileReader struct{}

// Function yang akan membuka file markdown dari path yang dikasih dari slug.
// Function ini akan me-returns file content sebagai type string, atau memberikan error 
// jika file yang dikasih tidak bisa dibuka(corrupt, not found, dll)
func (fRead FileReader) Read(slug string) (string, error) {
	fileRead, err := os.Open("markdown/" + slug + ".md")
	if err != nil {
		return "", err
	}
	defer fileRead.Close()
	b, err := io.ReadAll(fileRead)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Function ini akan membaca semua markdown files dari folder 'markdowns' 
// Function ini akan parse frontmatter dan content dan konversi file markdown tersebut ke HTML
// Dan akan mengembalikannya beberapa bagian PostData atau error jika ada yang salah dengan file markdown 
// atau folder markdown tidak ditemukan
func loadMarkdownPosts(dir string) ([]PostData, error) {
	md := goldmark.New()
	var posts []PostData

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check file .md
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var postData PostData
			var buf bytes.Buffer

			// Memisahkan konten dan extract YAML frontmatter dan bagian body dari file markdown 
			split := strings.SplitN(string(content), "\n---\n", 2)
			if len(split) > 1 {
				// Parse YAML front matter -> Convert Markdown to HTML -> Assign HTML content to PostData
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
				// Menangani masalah jika frontmatter tidak ditemukan pada markdown files
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
