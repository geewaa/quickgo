package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

//go:embed template
var fileFS embed.FS

func usage() {
	fmt.Fprintf(os.Stderr, "usage: quickgo [dstmod [dir]]\n")
	os.Exit(2)
}
func main() {
	log.SetPrefix("quickgo: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 || len(args) > 2 {
		usage()
	}

	tmplpath := "template"
	var dir string
	if len(args) == 2 {
		dir = args[1]
	} else {
		dir = args[0]
	}

	dstMod := ""
	if len(args) >= 1 {
		dstMod = args[0]
		// if err := module.CheckPath(dstMod); err != nil {
		// 	log.Fatalf("invalid destination module name: %v", err)
		// }
		if strings.Contains(dstMod, "/") {
			log.Fatalf("invalid destination module name: %v", dstMod)
		}
	}
	fs.WalkDir(fileFS, "template", func(src string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		rel, err := filepath.Rel(tmplpath, src)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println("now:", rel)
		if strings.Contains(rel, "{appname}") {
			rel = strings.ReplaceAll(rel, "{appname}", dstMod)
		}
		dst := filepath.Join(dir, rel)
		if d.IsDir() {
			if err := os.MkdirAll(dst, 0777); err != nil {
				log.Fatal(err)
			}
			return nil
		}

		data, err := fileFS.ReadFile(src)
		if err != nil {
			log.Fatal(err)
		}

		// isRoot := !strings.Contains(rel, string(filepath.Separator))
		if strings.HasSuffix(rel, ".go") {
			data = fixGo(data, dstMod)
		}
		if rel == "go.mod.tmpl" {
			data = fixGoMod(data, dstMod)
			dst = strings.TrimRight(dst, ".tmpl")
		}

		if err := os.WriteFile(dst, data, 0666); err != nil {
			log.Fatal(err)
		}
		return nil
	})
	// filepath.WalkDir(tmplpath, func(src string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	rel, err := filepath.Rel(tmplpath, src)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println("now:", rel)
	// 	if strings.Contains(rel, "{appname}") {
	// 		rel = strings.ReplaceAll(rel, "{appname}", dstMod)
	// 	}
	// 	dst := filepath.Join(dir, rel)
	// 	if d.IsDir() {
	// 		if err := os.MkdirAll(dst, 0777); err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		return nil
	// 	}

	// 	data, err := os.ReadFile(src)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	isRoot := !strings.Contains(rel, string(filepath.Separator))
	// 	if strings.HasSuffix(rel, ".go") {
	// 		data = fixGo(data, rel, dstMod, isRoot)
	// 	}
	// 	if rel == "go.mod" {
	// 		data = fixGoMod(data, dstMod)
	// 	}

	// 	if err := os.WriteFile(dst, data, 0666); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return nil
	// })
	log.Println("begin go mod tidy ...")

	cmd := exec.Command("go", "mod", "tidy")
	cdir, _ := os.Getwd()

	cmd.Dir = path.Join(cdir, dir)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Println(string(stdoutStderr))

	// cmd := exec.Command("go", "install", config.WireCmd)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalf("go install %s error\n", err)
	// }
	log.Println("begin wire ...")

	cmd1 := exec.Command("go", "generate", fmt.Sprintf("./cmd/%sapp/", dstMod))
	cmd1.Dir = path.Join(cdir, dir)
	stdoutStderr1, err := cmd1.CombinedOutput()
	if err != nil {
		log.Println(err)
		log.Println("run this youself late:\n\tgo generate", fmt.Sprintf("./cmd/%sapp/", dstMod))
	}
	log.Println(string(stdoutStderr1))
	log.Printf("\n\tcd %s\n\tgo run ./...", dir)
}

// fixGo rewrites the Go source in data to replace srcMod with dstMod.
// isRoot indicates whether the file is in the root directory of the module,
// in which case we also update the package name.
func fixGo(data []byte, dstMod string) []byte {
	// fset := token.NewFileSet()
	// f, err := parser.ParseFile(fset, file, data, parser.ImportsOnly)
	// if err != nil {
	// 	log.Fatalf("parsing source module:\n%s", err)
	// }

	// buf := NewBuffer(data)
	// at := func(p token.Pos) int {
	// 	return fset.File(p).Offset(p)
	// }

	// srcName := path.Base("{appname}")
	// dstName := path.Base(dstMod)
	// if isRoot {
	// 	if name := f.Name.Name; name == srcName || name == srcName+"_test" {
	// 		dname := dstName + strings.TrimPrefix(name, srcName)
	// 		if !token.IsIdentifier(dname) {
	// 			log.Fatalf("%s: cannot rename package %s to package %s: invalid package name", file, name, dname)
	// 		}
	// 		buf.Replace(at(f.Name.Pos()), at(f.Name.End()), dname)
	// 	}
	// }

	// for _, spec := range f.Imports {
	// 	path, err := strconv.Unquote(spec.Path.Value)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	if path == "{appname}" {
	// 		if srcName != dstName && spec.Name == nil {
	// 			// Add package rename because source code uses original name.
	// 			// The renaming looks strange, but template authors are unlikely to
	// 			// create a template where the root package is imported by packages
	// 			// in subdirectories, and the renaming at least keeps the code working.
	// 			// A more sophisticated approach would be to rename the uses of
	// 			// the package identifier in the file too, but then you have to worry about
	// 			// name collisions, and given how unlikely this is, it doesn't seem worth
	// 			// trying to clean up the file that way.
	// 			buf.Insert(at(spec.Path.Pos()), srcName+" ")
	// 		}
	// 		// Change import path to dstMod
	// 		buf.Replace(at(spec.Path.Pos()), at(spec.Path.End()), strconv.Quote(dstMod))
	// 	}
	// 	if strings.HasPrefix(path, "{appname}"+"/") {
	// 		// Change import path to begin with dstMod
	// 		buf.Replace(at(spec.Path.Pos()), at(spec.Path.End()), strconv.Quote(strings.Replace(path, "{appname}", dstMod, 1)))
	// 	}
	// }

	buf := strings.ReplaceAll(string(data), "{appname}", dstMod)
	return []byte(buf)
}

// fixGoMod rewrites the go.mod content in data to replace srcMod with dstMod
// in the module path.
func fixGoMod(data []byte, dstMod string) []byte {

	// 	cmd := exec.Command("go", "mod", "edit", "-module", p.ProjectName)
	ndata := []byte(strings.ReplaceAll(string(data), "{appname}", dstMod))

	f, err := modfile.ParseLax("go.mod", ndata, nil)
	if err != nil {
		log.Fatalf("parsing source module:\n%s", err)
	}
	f.AddModuleStmt(dstMod)
	new, err := f.Format()
	if err != nil {
		return ndata
	}
	return new
}
