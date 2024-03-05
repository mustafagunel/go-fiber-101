package handler

import (
	"fiber-tutorial/api/rest/middleware"
	"fiber-tutorial/internal/datastore"
	"fiber-tutorial/internal/model"
	"fiber-tutorial/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RESTHandler struct {
	Services service.Services
}

func NewRESTHandler(db *datastore.Mysqldb) *RESTHandler {
	return &RESTHandler{
		Services: service.InitServices(db),
	}
}

func (h *RESTHandler) CreateUser(c *fiber.Ctx) error {
	var resp middleware.Response
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		resp.Status = fiber.StatusBadRequest
		resp.Error = err.Error()
		return resp.Send(c)
	}

	res, err := h.Services.UserService.CreateUser(&user)
	if err != nil {
		resp.Status = fiber.StatusInternalServerError
		resp.Error = err.Error()
		return resp.Send(c)
	}
	resp.Status = fiber.StatusOK
	resp.Data = res
	return resp.Send(c)
}

func (h *RESTHandler) GetUser(c *fiber.Ctx) error {
	var resp middleware.Response
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		resp.Status = fiber.StatusBadRequest
		resp.Error = err.Error()
		return resp.Send(c)
	}

	res, err := h.Services.UserService.GetUser(&user)
	if err != nil {
		resp.Status = fiber.StatusInternalServerError
		resp.Error = err.Error()
		return resp.Send(c)
	}

	resp.Status = fiber.StatusOK
	resp.Data = res
	return resp.Send(c)
}
