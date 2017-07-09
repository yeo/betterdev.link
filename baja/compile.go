package baja

import (
	"log"

	//"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

//TODO Move theme reference to config with DI
func Compile(source string) error {
	page := buildPage()

	createRSS(page)
	createHome(page)
	createIssues(page)

	CopyDir("themes/yeo/assets", "public/assets")
	CopyDir("static/", "public/")

	return nil
}

func createRSS(page Page) {
	f, err := os.Create("./public/rss.xml")
	if err != nil {
		log.Println("Error creating file", err)
	}

	var publicableIssue Issues
	for i := len(page.Issues) - 1; i >= 0; i-- {
		issue := page.Issues[i]
		if issue.Draft == false {
			publicableIssue = append(publicableIssue, issue)
		}
	}
	page.Issues = publicableIssue

	funcMap := template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("themes/yeo/rss.tmpl"))
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}
func createHome(page Page) {
	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./public/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	//var tpl bytes.Buffer
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}

func createIssue(layout string, issue Issue) {
	t, err := template.ParseFiles("themes/yeo/"+layout+".tmpl", "themes/yeo/issue.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	directory := "./public/issues/" + issue.Name
	if issue.Draft == true {
		directory = "./public/issues/" + issue.Name + "/draft"
	}

	if layout == "email" {
		directory = "./public/issues/" + issue.Name + "/email"
	}
	os.MkdirAll(directory, 0755)

	f, err := os.Create(directory + "/index.html")
	if err != nil {
		log.Fatalf("Error creating file", err)
	}

	page := &Page{
		Time:  time.Now(),
		Issue: issue,
	}
	//var tpl bytes.Buffer
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}

}

func buildPage() Page {
	page := Page{
		Time: time.Now(),
	}

	var issues []Issue
	files, _ := ioutil.ReadDir("./content/issues/")
	for _, f := range files {
		if issue, err := loadIssue(f); err == nil {
			issues = append(issues, issue)
		}
	}

	// Sort Issue
	sort.Sort(Issues(issues))

	page.Issues = issues
	for k := len(page.Issues) - 1; k >= 0; k-- {
		page.Issue = page.Issues[k]
		if page.Issue.Draft == false {
			break
		}
	}
	return page
}

func createIssues(page Page) {
	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/issues.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./public/issues", 0755)
	f, err := os.Create("./public/issues/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	var publicableIssue Issues
	for _, issue := range page.Issues {
		if issue.Draft == false {
			publicableIssue = append(publicableIssue, issue)
		}

		createIssue("layout", issue)
		createIssue("email", issue)
	}

	//var tpl bytes.Buffer
	page.Issues = publicableIssue
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}

func loadIssue(f os.FileInfo) (Issue, error) {
	name := strings.Split(f.Name(), ".")
	issue := Issue{
		Name: name[0],
	}

	data, err := ioutil.ReadFile("./content/issues/" + f.Name())
	if err != nil {
		log.Fatal("Fail to parse issue", f.Name(), err)
		return issue, err
	}

	err = yaml.Unmarshal([]byte(data), &issue)
	if err != nil {
		log.Println("Error unmarshal yaml", err)
	}

	//Thu, 06 Jul 2017 17:42:26 +0000
	if issue.PubTime, err = time.Parse("Jan 2, 2006 15:04:05 -0700", issue.Time+" 05:19:00 -0700"); err != nil {
		log.Println("Cannot parse time on issue ", issue.Name, issue.Time, err)
	}
	return issue, nil
}

func CopyFile(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	return
}

// CopyDir recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
// Symlinks are ignored and skipped.
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	//if err == nil {
	//		return fmt.Errorf("destination already exists")
	//}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}
