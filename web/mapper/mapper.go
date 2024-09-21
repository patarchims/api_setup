package mapper

import "io/fs"

type WebMapper struct{}

type FileInfo struct {
	FileName string
	FileSize int
}

func (w *WebMapper) ToWebMapper(files []fs.FileInfo) (values []FileInfo) {
	for _, V := range files {
		if !V.IsDir() {
			values = append(values, FileInfo{
				FileName: V.Name(),
				FileSize: int(V.Size()),
			})
		}
	}

	return values

}
