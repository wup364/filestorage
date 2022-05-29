// Copyright (C) 2019 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 配置工具-JSON文件实现
// 依赖包: utypes.Object fileutil

package jsonconfig

import (
	"encoding/json"
	"errors"
	"os"
	"pakku/ipakku"
	"pakku/utils/fileutil"
	"pakku/utils/strutil"
	"pakku/utils/utypes"
	"path/filepath"
	"strings"
	"sync"
)

func init() {
	// 注册实例实现
	ipakku.Override.RegisterInterfaceImpl(new(Config), "IConfig", "json")
}

// Config json配置器
type Config struct {
	jsonObject map[string]interface{}
	configPath string
	l          *sync.RWMutex
}

// Init 初始化解析器
func (config *Config) Init(appName string) error {
	path, err := filepath.Abs("./.conf/" + appName + ".json")
	if nil != err {
		return err
	}
	return config.InitConfig(path)
}

// InitConfig 初始化解析器
func (config *Config) InitConfig(configPath string) error {
	if len(configPath) == 0 {
		return errors.New("config file path is empty")
	}
	// 创建父级目录
	parent := strutil.GetPathParent(configPath)
	if !fileutil.IsExist(parent) {
		err := fileutil.MkdirAll(parent)
		if nil != err {
			return err
		}
	}
	config.configPath = configPath
	// 文件不存在则创建
	if !fileutil.IsFile(config.configPath) {
		err := config.writeFileAsJSON(config.configPath, make(map[string]interface{}))
		if nil != err {
			return err
		}
	}

	config.l = new(sync.RWMutex)
	config.l.Lock()
	defer config.l.Unlock()
	// Json to map
	config.jsonObject = make(map[string]interface{})
	return config.readFileAsJSON(config.configPath, &config.jsonObject)
}

// GetConfig 读取key的value信息
// 返回ConfigBody对象, 里面的值可能是string或者map
func (config *Config) GetConfig(key string) (res utypes.Object) {
	config.l.RLock()
	defer config.l.RUnlock()
	if len(key) == 0 || config.jsonObject == nil || len(config.jsonObject) == 0 {
		return
	}
	keys := strings.Split(key, ".")
	if keys == nil {
		return
	}
	var temp interface{}
	keyLength := len(keys)
	for i := 0; i < keyLength; i++ {
		// last key
		if i == keyLength-1 {
			if i == 0 {
				if tp, ok := config.jsonObject[keys[i]]; ok {
					res = utypes.NewObject(tp)
				}
			} else if temp != nil {
				if tp, ok := temp.(map[string]interface{})[keys[i]]; ok {
					res = utypes.NewObject(tp)
				}
			}
			return
		}

		//
		var _temp interface{}
		if temp == nil { // first
			if tp, ok := config.jsonObject[keys[i]]; ok {
				_temp = tp
			}
		} else { //
			if tp, ok := temp.(map[string]interface{})[keys[i]]; ok {
				_temp = tp
			}
		}

		// find
		if _temp != nil {
			temp = _temp
		} else {
			return
		}
	}
	return
}

// SetConfig 保存配置, key value 都为stirng
func (config *Config) SetConfig(key string, value string) error {
	if len(key) == 0 || len(value) == 0 {
		return errors.New("key or value is empty")
	}
	config.l.Lock()
	defer config.l.Unlock()
	keys := strings.Split(key, ".")
	keyLength := len(keys)
	var temp interface{}
	for i := 0; i < keyLength; i++ {
		// last key
		if i == keyLength-1 {
			if i == 0 {
				config.jsonObject[keys[i]] = value
			} else if temp != nil {
				temp.(map[string]interface{})[keys[i]] = value
			}
			err := config.writeFileAsJSON(config.configPath, config.jsonObject)
			return err
		}

		//
		var _temp interface{}
		if temp == nil { // first
			if tp, ok := config.jsonObject[keys[i]]; ok {
				_temp = tp
			} else {
				_temp = make(map[string]interface{})
				config.jsonObject[keys[i]] = _temp
			}
		} else { //
			if tp, ok := temp.(map[string]interface{})[keys[i]]; ok {
				_temp = tp
			} else {
				_temp = make(map[string]interface{})
				temp.(map[string]interface{})[keys[i]] = _temp
			}
		}

		// find
		if _temp != nil {
			temp = _temp
		}
	}
	return nil
}

// readFileAsJSON 读取Json文件
func (config *Config) readFileAsJSON(path string, v interface{}) error {
	if len(path) == 0 {
		return fileutil.PathNotExist("ReadFileAsJSON", path)
	}
	fp, err := os.OpenFile(path, os.O_RDONLY, 0)
	defer func() {
		if nil != fp {
			fp.Close()
		}
	}()

	if err == nil {
		st, stErr := fp.Stat()
		if stErr == nil {
			data := make([]byte, st.Size())
			_, err = fp.Read(data)
			if err == nil {
				return json.Unmarshal(data, v)
			}
		} else {
			err = stErr
		}
	}
	return err
}

// writeFileAsJSON 写入Json文件
func (config *Config) writeFileAsJSON(path string, v interface{}) error {
	if len(path) == 0 {
		return fileutil.PathNotExist("WriteFileAsJSON", path)
	}
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer func() {
		if nil != fp {
			fp.Close()
		}
	}()

	if err == nil {
		data, err := json.Marshal(v)
		if err == nil {
			_, err := fp.Write(data)
			return err
		}
		return err
	}
	return err
}
