package controller_test

import (
	"encoding/json"
	"fmt"
	controller "go-starter/app/conroller"
	"go-starter/app/repository"
	"go-starter/app/route"
	"go-starter/app/service"
	"go-starter/core/response"
	"go-starter/core/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestRouter 设置测试路由器
func setupTestRouter() *gin.Engine {
	// 1. 手动创建依赖链
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)
	routeHandler := route.NewRouteHandler(bookController)

	// 2. 初始化路由器
	return server.Init(routeHandler)
}

func TestGetBook_Success(t *testing.T) {
	// 1. 设置测试路由器
	router := setupTestRouter()

	// 2. 创建测试请求
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/book?id=1", nil)
	router.ServeHTTP(w, req)

	// 3. 验证响应
	assert.Equal(t, 200, w.Code)
	response := response.Response{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	fmt.Println(response)
	assert.Equal(t, response.Code, 200)
}
