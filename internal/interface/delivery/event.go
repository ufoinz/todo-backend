package delivery

import (
	"net/http"
	"strconv"
	"time"
	"todo-app/internal/domain/event"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (app *Application) CreateEvent(c *gin.Context) {
	var payload struct {
		Name    string    `json:"name" binding:"required,min=3"`
		Content string    `json:"content" binding:"required,min=10"`
		Time    time.Time `json:"time" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	ownerID := userID.(int64)

	ev := event.Event{
		OwnerId: ownerID,
		Name:    payload.Name,
		Content: payload.Content,
		Time:    payload.Time,
	}

	if err := app.Events.Create(&ev); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create event"})
		return
	}
	c.JSON(http.StatusCreated, ev)
}

func (app *Application) GetEvents(c *gin.Context) {
	all, err := app.Events.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch events"})
		return
	}

	userID, _ := c.Get("user_id")
	ownerID := userID.(int64)

	mine := make([]event.Event, 0, len(all))
	for _, e := range all {
		if e.OwnerId == ownerID {
			mine = append(mine, e)
		}
	}
	c.JSON(http.StatusOK, mine)
}

func (app *Application) GetEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	ev, err := app.Events.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch event"})
		}
		return
	}

	userID, _ := c.Get("user_id")
	if ev.OwnerId != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	c.JSON(http.StatusOK, ev)
}

func (app *Application) UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	existing, err := app.Events.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch event"})
		}
		return
	}

	userID, _ := c.Get("user_id")
	if existing.OwnerId != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	var payload struct {
		Name    string    `json:"name" binding:"required,min=3"`
		Content string    `json:"content" binding:"required,min=10"`
		Time    time.Time `json:"time" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing.Name = payload.Name
	existing.Content = payload.Content
	existing.Time = payload.Time

	if err := app.Events.Update(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}
	c.JSON(http.StatusOK, existing)
}

func (app *Application) DeleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event id"})
		return
	}

	ev, err := app.Events.GetByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch event"})
		}
		return
	}
	userID, _ := c.Get("user_id")
	if ev.OwnerId != userID.(int64) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		return
	}

	if err := app.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.Status(http.StatusNoContent)
}
