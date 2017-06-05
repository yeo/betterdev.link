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
	"strings"
	"time"
)

//TODO Move theme reference to config with DI
func Compile(source string) error {
	createHome()
	createIssues()

	CopyDir("themes/yeo/assets", "public/assets")

	return nil
}

func createHome() {
	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./public/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	lastIssue := Issue{}
	page := &Page{time.Now(), lastIssue}
	//var tpl bytes.Buffer
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}

func createIssue(issue *Issue) {
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
		Issue: Issue{},
	}
	//var tpl bytes.Buffer
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}

}

func createIssues() {
	issues := make([]*Issue, 10, 100)

	files, _ := ioutil.ReadDir("./content/issues/")
	for _, f := range files {
		if issue := loadIssue(f); issue != nil {
			createIssue(issue)
			issues = append(issues, issue)
		}
	}

	t, err := template.ParseFiles("themes/yeo/layout.tmpl", "themes/yeo/issues.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir("./public/issues/", 0755)
	f, err := os.Create("./public/issues/index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}

	page := &Page{
		Time: time.Now(),
	}
	//var tpl bytes.Buffer
	if err := t.ExecuteTemplate(f, "base", &page); err != nil {
		log.Fatal(err)
	}
}

func loadIssue(f os.FileInfo) *Issue {
	fmt.Println(f.Name())
	name := strings.Split(f.Name(), ".")
	issue := Issue{
		Name: name[0],
	}
	return &issue
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
