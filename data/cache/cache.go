package cache

import (
	"gitee.com/quant1x/pandas/data/category"
	"os"
	"path/filepath"
)

const (
	INFO_PATH = "info"
	DAY_PATH  = "day"
)

type CacheType int

const (
	CACHE_TARS CacheType = iota
	CACHE_CSV
)

var (
	// CACHE_ROOT_PATH cache路径
	CACHE_ROOT_PATH = category.DATA_ROOT_PATH
	CACHE_TYPE      CacheType
)

func init() {
	CACHE_TYPE = CACHE_CSV
}

// CheckFilepath
// 检查filename 文件路径, 如果不存在就创建
func CheckFilepath(filename string) error {
	path := filepath.Dir(filename)
	dir, err := os.Stat(path)
	if err != nil {
		return err
	}
	if dir.IsDir() {
		return nil
	}
	return os.MkdirAll(path, category.CACHE_DIR_MODE)
}

// Init
// 初始化缓存路径
func Init(path string) error {
	err := os.MkdirAll(path, category.CACHE_DIR_MODE)
	if err != nil {
		return err
	}
	CACHE_ROOT_PATH = path
	return nil
}

// GetInfoPath 证券信息路径
func GetInfoPath() string {
	return CACHE_ROOT_PATH + "/" + INFO_PATH
}

// GetDayPath 历史数据-日线缓存路径
func GetDayPath() string {
	return CACHE_ROOT_PATH + "/" + DAY_PATH
}