// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package filedatas

import (
	"datanode/ifilestorage"
	"sort"
	"strings"
)

// FileListSorter 文件排序
type FileListSorter struct {
	fsInfos   []ifilestorage.FNode
	SortField string
	Asc       bool
}

// Sort 排序
func (fsort FileListSorter) Sort(fis []ifilestorage.FNode) []ifilestorage.FNode {
	folders := make([]ifilestorage.FNode, 0)
	files := make([]ifilestorage.FNode, 0)
	for i := 0; i < len(fis); i++ {
		if fis[i].IsFile {
			files = append(files, fis[i])
		} else {
			folders = append(folders, fis[i])
		}
	}
	if len(folders) > 0 {
		sort.Sort(&FileListSorter{
			Asc:       fsort.Asc,
			SortField: fsort.SortField,
			fsInfos:   folders,
		})
	}
	if len(files) > 0 {
		sort.Sort(&FileListSorter{
			Asc:       fsort.Asc,
			SortField: fsort.SortField,
			fsInfos:   files,
		})
	}
	return append(folders, files...)
}

// 实现sort.Interface接口取元素数量方法
func (fsort *FileListSorter) Len() int {
	return len(fsort.fsInfos)
}

// 实现sort.Interface接口比较元素方法
func (fsort *FileListSorter) Less(i, j int) bool {
	less := false
	if fsort.SortField == "FileSize" {
		less = fsort.fsInfos[i].Size < fsort.fsInfos[j].Size
	} else if fsort.SortField == "Mtime" {
		less = fsort.fsInfos[i].Mtime < fsort.fsInfos[j].Mtime
	} else {
		// 默认Path
		lasti := strings.LastIndex(fsort.fsInfos[i].Path, "/")
		lastj := strings.LastIndex(fsort.fsInfos[j].Path, "/")
		if lasti < 0 {
			lasti = 0
		} else {
			lasti++
		}
		if lastj < 0 {
			lastj = 0
		} else {
			lastj++
		}
		less = fsort.fsInfos[i].Path[lasti:] < fsort.fsInfos[j].Path[lastj:]
	}
	if !fsort.Asc {
		less = !less
	}
	return less
}

// 实现sort.Interface接口交换元素方法
func (fsort *FileListSorter) Swap(i, j int) {
	fsort.fsInfos[i], fsort.fsInfos[j] = fsort.fsInfos[j], fsort.fsInfos[i]
}
