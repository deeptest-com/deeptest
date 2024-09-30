package _fileUtils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/deeptest-com/deeptest/pkg/consts"
	_commonUtils "github.com/deeptest-com/deeptest/pkg/lib/comm"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFile(filePath string) string {
	buf := ReadFileBuf(filePath)
	str := string(buf)
	str = _commonUtils.RemoveBlankLine(str)
	return str
}

func ReadFileBuf(filePath string) []byte {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []byte(err.Error())
	}

	return buf
}

func WriteFile(filePath string, content string) (err error) {
	dir := filepath.Dir(filePath)
	MkDirIfNeeded(dir)

	var d1 = []byte(content)
	err = ioutil.WriteFile(filePath, d1, 0666) //写入文件(字节数组)

	return
}
func WriteFileIfNotExist(filePath string, content string) (err error) {
	if FileExist(filePath) {
		return
	}

	WriteFile(filePath, content)

	return
}

func FileExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func MkDirIfNeeded(dir string) error {
	if !FileExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}

	return nil
}
func RmDir(dir string) error {
	if FileExist(dir) {
		err := os.RemoveAll(dir)
		return err
	}

	return nil
}

func IsDir(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

func AbsolutePath(pth string) string {
	if !IsAbsolutePath(pth) {
		pth, _ = filepath.Abs(pth)
	}

	pth = AddSepIfNeeded(pth)

	return pth
}

func IsAbsolutePath(pth string) bool {
	return path.IsAbs(pth) ||
		strings.Index(pth, ":") == 1 // windows
}

func GetFilesFromParams(arguments []string) []string {
	ret := make([]string, 0)

	for _, arg := range arguments {
		if strings.Index(arg, "-") != 0 {
			if arg == "." {
				arg = AbsolutePath(".")
			} else if strings.Index(arg, "."+_consts.FilePthSep) == 0 {
				arg = AbsolutePath(".") + arg[2:]
			} else if !IsAbsolutePath(arg) {
				arg = AbsolutePath(".") + arg
			}

			ret = append(ret, arg)
		} else {
			break
		}
	}

	return ret
}

func GetExeDir(workDir string) string { // where zd.exe file in
	var dir string
	arg1 := strings.ToLower(os.Args[0])

	if strings.Index(arg1, "go-build") < 0 { // release
		p, _ := exec.LookPath(os.Args[0])
		if strings.Index(p, _consts.FilePthSep) > -1 {
			dir = p[:strings.LastIndex(p, _consts.FilePthSep)]
		}
	} else { // debug
		dir = workDir
	}

	dir, _ = filepath.Abs(dir)
	dir = AddSepIfNeeded(dir)

	return dir
}

func GetWorkDir() string { // where ztf command in
	dir, _ := os.Getwd()
	dir, _ = filepath.Abs(dir)
	dir = AddSepIfNeeded(dir)

	return dir
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func GetFileName(pathOrUrl string) string {
	index := strings.LastIndex(pathOrUrl, _consts.FilePthSep)

	name := pathOrUrl[index+1:]
	return name
}

func GetFileNameWithoutExt(pathOrUrl string) string {
	name := GetFileName(pathOrUrl)
	index := strings.LastIndex(name, ".")
	return name[:index]
}

func GetExtName(pathOrUrl string) string {
	index := strings.LastIndex(pathOrUrl, ".")

	return pathOrUrl[index:]
}

func GetAbsolutePath(pth string) string {
	if !IsAbsolutePath(pth) {
		pth, _ = filepath.Abs(pth)
	}

	pth = AddSepIfNeeded(pth)

	return pth
}

func AddSepIfNeeded(pth string) string {
	if strings.LastIndex(pth, _consts.FilePthSep) < len(pth)-1 {
		pth += _consts.FilePthSep
	}
	return pth
}

func ListDir(pth string) (ret []string, err error) {
	dir, err := ioutil.ReadDir(pth)
	if err != nil {
		return
	}

	for _, fi := range dir {
		name := fi.Name()

		ret = append(ret, filepath.Join(pth, name))
	}

	return
}

func GetUserHome() (dir string, err error) {
	user, err := user.Current()
	if err == nil {
		dir = user.HomeDir
	} else { // cross compile support

		if "windows" == runtime.GOOS { // windows
			dir, err = homeWindows()
		} else { // Unix-like system, so just assume Unix
			dir, err = homeUnix()
		}
	}

	dir = AddSepIfNeeded(dir)

	return
}
func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
