package handler

import (
	"github.com/gofiber/fiber/v2"
)

func (f *FileWebHandler) OdontogramImageFiber(c *fiber.Ctx) error {
	// dir := "./images/odontogram/" // Directory path where the files are located
	// // files, err := ioutil.ReadDir(dir)
	// entries, err := os.ReadDir(dir)
	// if err != nil {
	// 	c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
	// 	return
	// }
	// infos := make([]fs.FileInfo, 0, len(entries))
	// for _, entry := range entries {
	// 	info, err := entry.Info()
	// 	if err != nil {
	// 		c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
	// 		return
	// 	}
	// 	infos = append(infos, info)
	// }

	// if err != nil {
	// 	c.String(http.StatusInternalServerError, fmt.Sprintf("Error reading directory: %v", err))
	// 	return
	// }
	// data := f.WebMapper.ToWebMapper(infos)
	// c.HTML(http.StatusOK, "odontogram_index.html", gin.H{"users": data})
	// engine := html.New("./views", ".html")

	return c.Render("views/index", fiber.Map{
		"Title": "Hello, World!",
	})
}
