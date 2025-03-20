package cat

import (
	"net/http"

	catschemas "github.com/Nxwbtk/Mono-repo-template/Backend-template/services/Cat/cat-schemas"
	"github.com/gin-gonic/gin"
)

type CatHandler struct {
	CatService CatService
}

func NewCatHandler(catService CatService) *CatHandler {
	return &CatHandler{
		CatService: catService,
	}
}

// @Summary Get all cat
// @Schemes
// @Tags Cat
// @Description Retrieve all cat information
// @Accept json
// @Produce json
// @Success 200 {array} catschemas.TGetCat
// @Router /cat [get]
// @Security BearerAuth
func (h *CatHandler) GetCats(c *gin.Context) {
	cats, err := h.CatService.GetCatsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

// @Summary Get a cat
// @Schemes
// @Tags Cat
// @Description Retrieve a cat information
// @Accept json
// @Produce json
// @Param id path string true "Cat ID"
// @Success 200 {object} catschemas.TGetCat
// @Router /cat/{id} [get]
// @Security BearerAuth
func (h *CatHandler) GetCatsByID(c *gin.Context) {
	id := c.Param("id")
	cat, err := h.CatService.GetCatsByIDService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// @Summary Post a cat
// @Schemes
// @Tags Cat
// @Description Post a cat
// @Accept json
// @Produce json
// @Param cat body catschemas.TPostCat true "Cat information"
// @Success 201 {object} catschemas.TGetCat
// @Router /cat [post]
// @Security BearerAuth
func (h *CatHandler) CreateCat(c *gin.Context) {

	var cat catschemas.TPostCat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.CatService.CreateCatService(cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// @Summary Update a cat
// @Schemes
// @Tags Cat
// @Description Update a cat
// @Accept json
// @Produce json
// @Param id path string true "Cat ID"
// @Param cat body catschemas.TPostCat true "Cat information"
// @Success 200 {object} catschemas.TGetCat
// @Router /cat/{id} [put]
// @Security BearerAuth
func (h *CatHandler) UpdateCatByID(c *gin.Context) {
	id := c.Param("id")
	var cat catschemas.TPostCat
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.CatService.UpdateCatService(id, cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete a cat
// @Schemes
// @Tags Cat
// @Description Delete a cat
// @Accept json
// @Produce json
// @Param id path string true "Cat ID"
// @Success 200
// @Router /cat/{id} [delete]
// @Security BearerAuth
func (h *CatHandler) DeleteCatByID(c *gin.Context) {
	id := c.Param("id")
	err := h.CatService.DeleteCatService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *CatHandler) RegisterCatrouters(r gin.IRoutes) {
	r.GET("/", h.GetCats)
	r.GET("/:id", h.GetCatsByID)
	r.POST("/", h.CreateCat)
	r.PUT("/:id", h.UpdateCatByID)
	r.DELETE("/:id", h.DeleteCatByID)
}
