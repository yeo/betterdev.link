package baja

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/russross/blackfriday"
	"gopkg.in/yaml.v2"
)

type Posts []Post

type Post struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
	Draft       bool   `yaml:"draft"`

	Slug    string        `yaml:"-"`
	PubTime time.Time     `yaml:"-"`
	Body    template.HTML `yaml:"-"`
	TOC     template.HTML `yaml:"-"`
}

func (p Post) FormatPubTime() string {
	return p.PubTime.Format("January 2, 2006")
}

func (p Post) ReadingTime() int {
	words := len(strings.Fields(string(p.Body)))
	minutes := words / 200
	if minutes < 1 {
		minutes = 1
	}
	return minutes
}

func (p Posts) Len() int           { return len(p) }
func (p Posts) Less(i, j int) bool { return p[i].PubTime.After(p[j].PubTime) }
func (p Posts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func loadPost(name string) (Post, error) {
	post := Post{
		Slug: strings.TrimSuffix(name, ".md"),
	}

	data, err := ioutil.ReadFile("./content/posts/" + name)
	if err != nil {
		log.Println("Fail to read post", name, err)
		return post, err
	}

	// Split YAML front matter from the markdown body
	body := string(data)
	if strings.HasPrefix(body, "---") {
		parts := strings.SplitN(body, "---", 3)
		if len(parts) == 3 {
			if err := yaml.Unmarshal([]byte(parts[1]), &post); err != nil {
				log.Println("Error unmarshal front matter of post", name, err)
			}
			body = parts[2]
		}
	}

	if post.PubTime, err = time.Parse("2006-01-02", post.Date); err != nil {
		log.Println("Cannot parse date on post", name, post.Date, err)
	}

	// [> note text <] becomes a margin side note. Split on ``` so
	// fenced code blocks (odd chunks) are left untouched.
	chunks := strings.Split(body, "```")
	for i := 0; i < len(chunks); i += 2 {
		chunks[i] = sidenoteRe.ReplaceAllString(chunks[i], `<span class="sidenote">$1</span>`)
	}
	body = strings.Join(chunks, "```")

	html, toc := buildTOC(string(blackfriday.Run([]byte(body))))
	post.Body = template.HTML(html)
	post.TOC = template.HTML(toc)

	return post, nil
}

var (
	sidenoteRe = regexp.MustCompile(`(?s)\[>\s*(.*?)\s*<\]`)
	headingRe  = regexp.MustCompile(`<h([23])(?: id="([^"]*)")?>(.*?)</h[23]>`)
	tagRe      = regexp.MustCompile(`<[^>]+>`)
)

func slugify(s string) string {
	s = strings.ToLower(tagRe.ReplaceAllString(s, ""))

	var b strings.Builder
	dash := true
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			dash = false
		} else if !dash {
			b.WriteRune('-')
			dash = true
		}
	}

	return strings.Trim(b.String(), "-")
}

type tocHeading struct {
	Level int
	ID    string
	Title string
}

// buildTOC gives every h2/h3 an anchor id and returns the html along a
// nested ul linking to them. The ul is empty when a post has too few
// headings to deserve a TOC.
func buildTOC(html string) (string, string) {
	var headings []tocHeading
	seen := map[string]int{}

	html = headingRe.ReplaceAllStringFunc(html, func(m string) string {
		parts := headingRe.FindStringSubmatch(m)

		level := 2
		if parts[1] == "3" {
			level = 3
		}

		id := parts[2]
		if id == "" {
			id = slugify(parts[3])
			if id == "" {
				id = "section"
			}
			if n := seen[id]; n > 0 {
				seen[id] = n + 1
				id = fmt.Sprintf("%s-%d", id, n)
			} else {
				seen[id] = 1
			}
		}

		headings = append(headings, tocHeading{level, id, parts[3]})
		return fmt.Sprintf(`<h%d id="%s">%s</h%d>`, level, id, parts[3], level)
	})

	if len(headings) < 2 {
		return html, ""
	}

	var b strings.Builder
	b.WriteString("<ul>")
	openItem, openSub := false, false
	for _, h := range headings {
		if h.Level == 2 {
			if openSub {
				b.WriteString("</ul>")
				openSub = false
			}
			if openItem {
				b.WriteString("</li>")
			}
			fmt.Fprintf(&b, `<li><a href="#%s">%s</a>`, h.ID, h.Title)
			openItem = true
		} else {
			if !openSub {
				b.WriteString("<ul>")
				openSub = true
			}
			fmt.Fprintf(&b, `<li><a href="#%s">%s</a></li>`, h.ID, h.Title)
		}
	}
	if openSub {
		b.WriteString("</ul>")
	}
	if openItem {
		b.WriteString("</li>")
	}
	b.WriteString("</ul>")

	return html, b.String()
}

func loadPosts() Posts {
	var posts Posts

	files, _ := ioutil.ReadDir("./content/posts/")
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".md") {
			if post, err := loadPost(f.Name()); err == nil {
				posts = append(posts, post)
			}
		}
	}

	sort.Sort(posts)

	return posts
}

func createBlog(page Page) {
	posts := loadPosts()

	for _, post := range posts {
		createPost(post)
	}

	// Drafts are rendered at their URL but kept off the index
	var publicable Posts
	for _, post := range posts {
		if post.Draft == false {
			publicable = append(publicable, post)
		}
	}
	page.Posts = publicable

	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/blog.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	os.MkdirAll("./public/blog", 0755)
	f, err := os.Create("./public/blog/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}

func createPost(post Post) {
	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/post.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	directory := "./public/blog/" + post.Slug
	os.MkdirAll(directory, 0755)

	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	page := &Page{
		Time: time.Now(),
		Post: post,
	}

	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}
