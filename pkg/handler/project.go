package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	todo "github.com/delapaska/university-storage"
	"github.com/gin-gonic/gin"
)

type CreateProjectRequest struct {
	ProjectName string `json:"projectName"`
}

func (h *Handler) createProject(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "create.html", nil)
		return
	}
	var reqBody CreateProjectRequest
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode request body"})
		return
	}

	// Выводим введенное имя проекта в консоль
	projectName := reqBody.ProjectName
	println("Received project name:", projectName)
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		return
	}

	//var input todo.User

	input := todo.ProjectList{
		Title: projectName,
	}
	//var input todo.ProjectList
	fmt.Println(input)
	/*
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	*/
	id, err := h.services.ProjectList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.ProjectList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	fmt.Println(userId)
	if err != nil {
		return
	}

	projects, err := h.services.ProjectList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var projectsHTML strings.Builder
	projectsHTML.WriteString("<h2>Список проектов:</h2>")
	projectsHTML.WriteString("<ul>")
	for _, project := range projects {
		// Генерируем ссылку для каждого проекта
		projectsHTML.WriteString("<li><a href=\"/api/projects/list/" + strconv.Itoa(project.Id) + "\">" + project.Title + "</a></li>")
	}
	projectsHTML.WriteString("</ul>")

	// Отправляем HTML-страницу пользователю
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(projectsHTML.String()))
	/*
		c.JSON(http.StatusOK, getAllListsResponse{
			Data: projects,
		})
	*/
}

type TemplateData struct {
	Directory string
	Files     []string
}

type ProjectData struct {
	Project todo.ProjectList
	Folders []string
}

// Метод обработки запроса для отображения информации о проекте и списка папок
func (h *Handler) getProjectById(c *gin.Context) {
	// Получаем идентификатор пользователя из контекста запроса
	userId, err := getUserId(c)
	if err != nil {
		// Обработка ошибки, если не удалось получить идентификатор пользователя
		return
	}
	files, _ := listFiles("./uploads")
	fmt.Println(files)
	// Получаем идентификатор проекта из URL-параметра запроса
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Обработка ошибки, если не удалось преобразовать идентификатор проекта в число
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	// Получаем информацию о проекте по его идентификатору
	project, err := h.services.ProjectList.GetById(userId, id)
	if err != nil {
		// Обработка ошибки, если не удалось получить информацию о проекте из базы данных
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Получаем список папок для проекта по его идентификатору
	folders, err := h.services.GetAllFolders(id)
	if err != nil {
		// Обработка ошибки, если не удалось получить список папок проекта из базы данных
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении папок проекта из базы данных"})
		return
	}

	// Отображение HTML страницы с информацией о проекте и списком папок
	c.HTML(http.StatusOK, "project.html", gin.H{
		"files":   files,
		"Id":      project.Id,
		"Title":   project.Title,
		"Folders": folders,
	})
}
func listFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, filepath.Base(path))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
func (h *Handler) createFolder(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	projectName := c.PostForm("folder_name")

	h.services.CreateFolder(id, projectName)
	link := "/api/projects/list/" + strconv.Itoa(id)
	c.Redirect(http.StatusSeeOther, link)

}

func (h *Handler) uploadFiles(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))
		return
	}
	filename := header.Filename
	filename = strings.ReplaceAll(filename, " ", "_") // Заменяем пробелы в имени файла на нижнее подчеркивание
	out, err := os.Create(filepath.Join("./uploads", filename))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err: %s", err.Error()))
		return
	}

	_, err = db.Exec("INSERT INTO files (folder_id, filename, filepath) VALUES ($1, $2, $3)", 1, filename, filename)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error inserting file into database: %s", err.Error()))
		return
	}
	link := "/api/projects/list/" + strconv.Itoa(id)
	c.Redirect(http.StatusFound, link)
}
func (h *Handler) connectProject(c *gin.Context) {
	projectId := c.Param("id")

	intId, _ := strconv.Atoi(projectId)
	token, err := h.services.GenerateProjectToken(intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена доступа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) connectToProject(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "connect.html", nil)
		return
	}
	id, _ := strconv.Atoi(c.PostForm("id"))
	token := c.PostForm("token")
	userId, err := getUserId(c)
	if err != nil {
		// Обработка ошибки, если не удалось получить идентификатор пользователя
		return
	}

	h.services.ConnectUserToProject(userId, id, token)

}
