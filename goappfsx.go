// Package goappfsx provides ergonomic wrappers for common file system tasks for
// go applications.
package goappfsx

import (
	"io/ioutil"
	"os"
	"path"
)

// DataCategory is an enumeration of the categories used in Windows to dileniate
// what files should be sync'd (Roaming) or available to restricted privelege
// (LocalLow) runs of the program.  'None' is also an option.  When converted to
// a `string`, 'None' will be the empty string.
type DataCategory int

// DataCategory enumerations.
const (
	None DataCategory = iota
	Local
	LocalLow
	Roaming
)

func (dc DataCategory) String() string {
	categories := [4]string{"", "Local", "LocalLow", "Roaming"}
	return categories[dc]
}

// ExeDir returns the directory of the executable.  Any errors encountered are
// passed directly to the caller.
func ExeDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return path.Dir(exePath), nil
}

// OpenFileExeDir returns a reference to an `os.File` given a path supplement to
// the executable directory.  Any errors encountered are passed directly to the
// caller.
func OpenFileExeDir(pathSupplement string) (*os.File, error) {
	ed, err := ExeDir()
	if err != nil {
		return nil, err
	}
	fp := path.Join(ed, pathSupplement)
	fyle, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return fyle, err
}

// ReadFileExeDir returns a `[]byte` that is the result of reading the entire
// contents of the requested file.  `pathSupplement` is the portion of the path
// relative to the executable directory.
func ReadFileExeDir(pathSupplement string) ([]byte, error) {
	ed, err := ExeDir()
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadFile(path.Join(ed, pathSupplement))
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteFileExeDir writes a `[]byte` to the path provided in pathSupplement.  If
// the file does not exist, it creates it.  If the file does exist, it truncates
// it first.  Returns the number of `byte`s written.  Wraps the `os.File`
// methods and passes any errors encounted directly to the caller.
func WriteFileExeDir(pathSupplement string, data []byte) (int, error) {
	ed, err := ExeDir()
	if err != nil {
		return 0, err
	}
	fyle, err := os.Create(path.Join(ed, pathSupplement))
	if err != nil {
		return 0, err
	}
	defer fyle.Close()
	return fyle.Write(data)
}

// AppDataDir returns the application directory within the current user's
// configuration directory.  It wraps `os.UserConfigDir` and passes any errors
// encountered directly to the caller.  If the directory doesn't yet exist, it
// creates it prior to returning.
func AppDataDir(category DataCategory) (string, error) {
	progName := path.Base(os.Args[0])
	progExt := path.Ext(progName)
	progName = progName[:len(progName)-len(progExt)]

	appDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appDir = path.Join(appDir, category.String(), progName)
	os.MkdirAll(appDir, os.ModeDir)
	return appDir, nil
}

// OpenFileAppDataDir returns a reference to an `os.File` given a path
// supplement to the User's configuration directory.  It wraps
// `os.UserConfigDir` and passes any errors encountered directly to the caller.
// If the application specific configuration directory doesn't yet exist, it
// creates it in the process.  It does *not* create the file if that doesn't
// exist.
func OpenFileAppDataDir(pathSupplement string, category DataCategory) (*os.File, error) {
	add, err := AppDataDir(category)
	if err != nil {
		return nil, err
	}
	fp := path.Join(add, pathSupplement)
	fyle, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return fyle, err
}

// ReadFileAppDataDir returns a `[]byte` that is the result of reading the
// entire contents of the requested file.  `pathSupplement` is the portion of
// the path relative to the application specific directory in the User's
// configuration directory.  If the application specific configuration directory
// doesn't yet exist, it creates it in the process.  It does *not* create the
// file if that doesn't exist.
func ReadFileAppDataDir(pathSupplement string, category DataCategory) ([]byte, error) {
	add, err := AppDataDir(category)
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadFile(path.Join(add, pathSupplement))
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WriteFileAppDataDir writes a `[]byte` to the path provided in pathSupplement.
// If the file does not exist, it creates it.  If the file does exist, it
// truncates it first.  Returns the number of `byte`s written.  Wraps the
// `os.File` methods and passes any errors encounted directly to the caller.
func WriteFileAppDataDir(pathSupplement string, category DataCategory, data []byte) (int, error) {
	add, err := AppDataDir(category)
	if err != nil {
		return 0, err
	}
	fyle, err := os.Create(path.Join(add, pathSupplement))
	if err != nil {
		return 0, err
	}
	defer fyle.Close()
	return fyle.Write(data)
}
