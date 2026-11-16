package api

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
)

const (
	DefaultTagName = "dataframe"
)

// NewError 创建带状态码的错误信息
func NewError(code int, message string, a ...any) error {
	return fmt.Errorf("#%d, message=%s", code, fmt.Sprintf(message, a...))
}

// Reverse 反转切片
func Reverse[S ~[]E, E any](s S) S {
	d := slices.Clone(s)
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	return d
}

// CloseQuietly 关闭 io, 忽略错误
func CloseQuietly(closer io.Closer) {
	if closer != nil {
		_ = closer.Close()
	}
}

// CheckFilepath 检查文件路径, 不存在时按需创建
func CheckFilepath(filename string, notExistToCreate ...bool) error {
	filename, _ = Expand(filename)
	path := filepath.Dir(filename)
	if _, err := os.Lstat(path); err == nil || os.IsExist(err) {
		return nil
	}
	create := len(notExistToCreate) > 0 && notExistToCreate[0]
	if !create {
		return os.ErrNotExist
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	return os.ErrNotExist
}

// Expand 展开路径中前缀 ~ 为用户主目录
func Expand(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}
	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}
	dir, err := homeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, path[1:]), nil
}

func homeDir() (string, error) {
	if home := os.Getenv("GOX_HOME"); home != "" {
		return home, nil
	}
	return os.UserHomeDir()
}
