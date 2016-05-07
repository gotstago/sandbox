package services

import (
	. "os"
	"path/filepath"
	"runtime"
	"testing"
)

var isReadonlyError = func(error) bool { return false }

func TestGetGopath(t *testing.T){
    _,err := GetGopath()
    if err != nil{
        t.Fatal("we need a GOPATH",err)
    }
    //t.Log("GOPATH is ",goPath)
}

func TestGetGopathNoGopath(t *testing.T){
    //no GOPATH, should succeed
    current := Getenv("GOPATH")
    defer Setenv("GOPATH",current)
    Setenv("GOPATH","")
    goPath,err := GetGopath()
    if err != nil{
        t.FailNow()
    }
    if goPath != (Getenv("HOME") + "/go"){
        t.Fatal("Expecting GOPATH to default to $HOME/go, but found",goPath)
    }
    //no HOME, should fail
    Setenv("GOPATH","")
    Setenv("HOME","")
    goPath,err = GetGopath()
    if err == nil{
        t.FailNow()
    }
    // if goPath != (Getenv("HOME") + "/go"){
    //     t.Fatal("Expecting GOPATH to default to $HOME/go, but found",goPath)
    // }
    
}

func TestCreateProjectFolders(t *testing.T){
    goPath,err := GetGopath()
    if err != nil {
        t.Fatal("should be a GOPATH")
    }
    projectName := "myProject"
    
    targetRoot := goPath + "/src/github.com/gotstago/qor-projects/" + projectName
    targetController := targetRoot + "/app/controllers"
    targetModels := targetRoot + "/app/models"
    targetViews := targetRoot + "/app/views"
    targetAdmin := targetRoot + "/config/admin"
    targetRoutes := targetRoot + "/config/routes"
    targetDb := targetRoot + "/db"
    targetDist := targetRoot + "/public/dist"
    targetFonts := targetRoot + "/public/fonts"
    targetImages := targetRoot + "/public/images"
    targetJavascripts := targetRoot + "/public/javascripts"
    targetStylesheets := targetRoot + "/public/stylesheets"
    targetVendors := targetRoot + "/public/vendors"
    err = CreateProjectFolders(targetRoot,targetController,targetModels,targetViews,targetAdmin,targetRoutes,targetDb,targetDist,targetFonts,targetImages,targetJavascripts,targetStylesheets,targetVendors)
    if err != nil{
        t.Fatal(err)
    }
}

func TestRemoveProjectFolder(t *testing.T){
    goPath,err := GetGopath()
    if err != nil {
        t.Fatal("should be a GOPATH")
    }
    projectName := "myProject"
    
    targetRoot := goPath + "/src/github.com/gotstago/qor-projects/" + projectName
    err = RemoveProject(targetRoot)
    if err != nil{
        t.Fatal(err)
    }
}

func TestMkdirAll(t *testing.T) {
	tmpDir := TempDir()
	path := tmpDir + "/_TestMkdirAll_/dir/./dir2"
	err := MkdirAll(path, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q: %s", path, err)
	}
	defer RemoveAll(tmpDir + "/_TestMkdirAll_")

	// Already exists, should succeed.
	err = MkdirAll(path, 0777)
	if err != nil {
		t.Fatalf("MkdirAll %q (second time): %s", path, err)
	}

	// Make file.
	fpath := path + "/file"
	f, err := Create(fpath)
	if err != nil {
		t.Fatalf("create %q: %s", fpath, err)
	}
	defer f.Close()

	// Can't make directory named after file.
	err = MkdirAll(fpath, 0777)
	if err == nil {
		t.Fatalf("MkdirAll %q: no error", fpath)
	}
	perr, ok := err.(*PathError)
	if !ok {
		t.Fatalf("MkdirAll %q returned %T, not *PathError", fpath, err)
	}
	if filepath.Clean(perr.Path) != filepath.Clean(fpath) {
		t.Fatalf("MkdirAll %q returned wrong error path: %q not %q", fpath, filepath.Clean(perr.Path), filepath.Clean(fpath))
	}

	// Can't make subdirectory of file.
	ffpath := fpath + "/subdir"
	err = MkdirAll(ffpath, 0777)
	if err == nil {
		t.Fatalf("MkdirAll %q: no error", ffpath)
	}
	perr, ok = err.(*PathError)
	if !ok {
		t.Fatalf("MkdirAll %q returned %T, not *PathError", ffpath, err)
	}
	if filepath.Clean(perr.Path) != filepath.Clean(fpath) {
		t.Fatalf("MkdirAll %q returned wrong error path: %q not %q", ffpath, filepath.Clean(perr.Path), filepath.Clean(fpath))
	}

	if runtime.GOOS == "windows" {
		path := tmpDir + `\_TestMkdirAll_\dir\.\dir2\`
		err := MkdirAll(path, 0777)
		if err != nil {
			t.Fatalf("MkdirAll %q: %s", path, err)
		}
	}
}

// func TestRemoveAll(t *testing.T) {
// 	tmpDir := TempDir()
// 	// Work directory.
// 	path := tmpDir + "/_TestRemoveAll_"
// 	fpath := path + "/file"
// 	dpath := path + "/dir"

// 	// Make directory with 1 file and remove.
// 	if err := MkdirAll(path, 0777); err != nil {
// 		t.Fatalf("MkdirAll %q: %s", path, err)
// 	}
// 	fd, err := Create(fpath)
// 	if err != nil {
// 		t.Fatalf("create %q: %s", fpath, err)
// 	}
// 	fd.Close()
// 	if err = RemoveAll(path); err != nil {
// 		t.Fatalf("RemoveAll %q (first): %s", path, err)
// 	}
// 	if _, err = Lstat(path); err == nil {
// 		t.Fatalf("Lstat %q succeeded after RemoveAll (first)", path)
// 	}

// 	// Make directory with file and subdirectory and remove.
// 	if err = MkdirAll(dpath, 0777); err != nil {
// 		t.Fatalf("MkdirAll %q: %s", dpath, err)
// 	}
// 	fd, err = Create(fpath)
// 	if err != nil {
// 		t.Fatalf("create %q: %s", fpath, err)
// 	}
// 	fd.Close()
// 	fd, err = Create(dpath + "/file")
// 	if err != nil {
// 		t.Fatalf("create %q: %s", fpath, err)
// 	}
// 	fd.Close()
// 	if err = RemoveAll(path); err != nil {
// 		t.Fatalf("RemoveAll %q (second): %s", path, err)
// 	}
// 	if _, err := Lstat(path); err == nil {
// 		t.Fatalf("Lstat %q succeeded after RemoveAll (second)", path)
// 	}

// 	// Determine if we should run the following test.
// 	testit := true
// 	if runtime.GOOS == "windows" {
// 		// Chmod is not supported under windows.
// 		testit = false
// 	} else {
// 		// Test fails as root.
// 		testit = Getuid() != 0
// 	}
// 	if testit {
// 		// Make directory with file and subdirectory and trigger error.
// 		if err = MkdirAll(dpath, 0777); err != nil {
// 			t.Fatalf("MkdirAll %q: %s", dpath, err)
// 		}

// 		for _, s := range []string{fpath, dpath + "/file1", path + "/zzz"} {
// 			fd, err = Create(s)
// 			if err != nil {
// 				t.Fatalf("create %q: %s", s, err)
// 			}
// 			fd.Close()
// 		}
// 		if err = Chmod(dpath, 0); err != nil {
// 			t.Fatalf("Chmod %q 0: %s", dpath, err)
// 		}

// 		// No error checking here: either RemoveAll
// 		// will or won't be able to remove dpath;
// 		// either way we want to see if it removes fpath
// 		// and path/zzz.  Reasons why RemoveAll might
// 		// succeed in removing dpath as well include:
// 		//	* running as root
// 		//	* running on a file system without permissions (FAT)
// 		RemoveAll(path)
// 		Chmod(dpath, 0777)

// 		for _, s := range []string{fpath, path + "/zzz"} {
// 			if _, err = Lstat(s); err == nil {
// 				t.Fatalf("Lstat %q succeeded after partial RemoveAll", s)
// 			}
// 		}
// 	}
// 	if err = RemoveAll(path); err != nil {
// 		t.Fatalf("RemoveAll %q after partial RemoveAll: %s", path, err)
// 	}
// 	if _, err = Lstat(path); err == nil {
// 		t.Fatalf("Lstat %q succeeded after RemoveAll (final)", path)
// 	}
// }

// func TestMkdirAllWithSymlink(t *testing.T) {
// 	switch runtime.GOOS {
// 	case "android", "nacl", "plan9":
// 		t.Skipf("skipping on %s", runtime.GOOS)
// 	case "windows":
// 		if !supportsSymlinks {
// 			t.Skipf("skipping on %s", runtime.GOOS)
// 		}
// 	}

// 	tmpDir, err := ioutil.TempDir("", "TestMkdirAllWithSymlink-")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer RemoveAll(tmpDir)

// 	dir := tmpDir + "/dir"
// 	err = Mkdir(dir, 0755)
// 	if err != nil {
// 		t.Fatalf("Mkdir %s: %s", dir, err)
// 	}

// 	link := tmpDir + "/link"
// 	err = Symlink("dir", link)
// 	if err != nil {
// 		t.Fatalf("Symlink %s: %s", link, err)
// 	}

// 	path := link + "/foo"
// 	err = MkdirAll(path, 0755)
// 	if err != nil {
// 		t.Errorf("MkdirAll %q: %s", path, err)
// 	}
// }

// func TestMkdirAllAtSlash(t *testing.T) {
// 	switch runtime.GOOS {
// 	case "android", "plan9", "windows":
// 		t.Skipf("skipping on %s", runtime.GOOS)
// 	case "darwin":
// 		switch runtime.GOARCH {
// 		case "arm", "arm64":
// 			t.Skipf("skipping on darwin/%s, mkdir returns EPERM", runtime.GOARCH)
// 		}
// 	}
// 	RemoveAll("/_go_os_test")
// 	const dir = "/_go_os_test/dir"
// 	err := MkdirAll(dir, 0777)
// 	if err != nil {
// 		pathErr, ok := err.(*PathError)
// 		// common for users not to be able to write to /
// 		if ok && (pathErr.Err == syscall.EACCES || isReadonlyError(pathErr.Err)) {
// 			t.Skipf("could not create %v: %v", dir, err)
// 		}
// 		t.Fatalf(`MkdirAll "/_go_os_test/dir": %v, %s`, err, pathErr.Err)
// 	}
// 	RemoveAll("/_go_os_test")
// }
