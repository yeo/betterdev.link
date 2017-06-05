package baja

import (
	"log"

	//"bytes"
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

//TODO Move theme reference to config with DI
func Compile(source string) error {
	page := buildPage()

	createHome(page)
	createIssues(page)

	CopyDir("themes/yeo/assets", "public/assets")

	return nil
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

func createIssue(issue Issue) {
	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/issue.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./public/issues/"+issue.Name, 0755)
	f, err := os.Create("./public/issues/" + issue.Name + "/index.html")
	if err != nil {
		log.Println("Error creating file", err)
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
			createIssue(issue)
			issues = append(issues, issue)
		}
	}

	// Sort Issue
	sort.Sort(Issues(issues))

	page.Issues = issues
	page.Issue = page.Issues[len(page.Issues)-1]
	return page
}

func createIssues(page Page) {
	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/issues.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./public/issues/", 0755)
	f, err := os.Create("./public/issues/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	//var tpl bytes.Buffer
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
	log.Println(issue)
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
