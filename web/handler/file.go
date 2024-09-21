package handler

import (
	"fmt"
	"hms_api/web/mapper"
	"io/fs"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FileWebHandler struct {
	WebMapper mapper.WebMapper
	Logging   *logrus.Logger
}

func (f *FileWebHandler) Assesmen(c *gin.Context) {
	c.HTML(http.StatusOK, "asesmen.html", "")
}

func (f *FileWebHandler) OdontogramImage(c *gin.Context) {
	dir := "./images/odontogram/" // Directory path where the files are located
	// files, err := ioutil.ReadDir(dir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
		return
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
			return
		}
		infos = append(infos, info)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
		return
	}
	data := f.WebMapper.ToWebMapper(infos)
	c.HTML(http.StatusOK, "odontogram_index.html", gin.H{"users": data})
}

func (f *FileWebHandler) AnatomiImages(c *gin.Context) {

	dir := "./images/anatomi/" // Directory path where the files are located

	entries, err := os.ReadDir(dir)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
		return
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
			return
		}
		infos = append(infos, info)
	}

	data := f.WebMapper.ToWebMapper(infos)

	c.HTML(http.StatusOK, "anatomi_index.html", gin.H{"users": data})
}
