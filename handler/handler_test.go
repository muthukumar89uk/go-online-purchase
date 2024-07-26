package handler

import (
	//Inbuild package(s)
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	//User defined package(s)
	"online/driver"
	"online/middleware"

	//Third party package(s)
	"github.com/labstack/echo"
)

var (
	AdminToken string
	UserToken  string
)

func TestSignup(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	e := echo.New()
	e.POST("/signup", database.Signup)
	t.Run("missing username", func(t *testing.T) {
		body := `{
			"username":"",
			"email":"hareesh@gmail.com",
			"password":"12345678",
			"role":"admin"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing email", func(t *testing.T) {
		body := `{
			"username":"hari",
			"email":"",
			"password":"12345678",
			"role":"admin"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing password", func(t *testing.T) {
		body := `{
			"username":"hari",
			"email":"hareesh@gmail.com",
			"password":"",
			"role":"admin"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing role", func(t *testing.T) {
		body := `{
			"username":"hari",
			"email":"hareesh@gmail.com",
			"password":"12345678",
			"role":""
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid email", func(t *testing.T) {
		body := `{
			"username":"hari",
			"email":"Hareeshgmailcom",
			"password":"12345678",
			"role":"admin"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid role", func(t *testing.T) {
		body := `{
			"username":"hari",
			"email":"hareesh@gmailcom",
			"password":"12345678",
			"role":"customer"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Checking the length of password", func(t *testing.T) {
		body := `{
			"username":"hari",
			"email":"hareesh@gmailcom",
			"password":"1234678",
			"role":"customer"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("signup successful(By Admin)", func(t *testing.T) {
		//While running this case, need to change email field for every time
		body := `{
			"username":"Ajith",
			"email":"ajith@gmail.com", 
			"password":"12345678",
			"role":"admin"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("signup successful(By User)", func(t *testing.T) {
		//While running this case, need to change email field for every time
		body := `{
			"username":"Vijay",
			"email":"vijay@gmail.com", 
			"password":"12345678",
			"role":"user"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("User already exist", func(t *testing.T) {
		body := `{
			"username":"Hari",
			"email":"ajith@gmailcom",
			"password":"12345678",
			"role":"admin"
		}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestLogin(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	e := echo.New()
	e.POST("/login", database.Login)
	t.Run("missing password", func(t *testing.T) {
		body := `{
			"email":"hareesh@gmail.com",
			"password":""
		}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing email", func(t *testing.T) {
		body := `{
			"email":"",
			"password":"12345678"
		}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid email", func(t *testing.T) {
		body := `{
			"email":"Hareeshgmailcom",
			"password":"12345678"
		}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("User not found", func(t *testing.T) {
		body := `{
			"email":"bharathi@gmail.com",
			"password":"65432178"
		}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		fmt.Println(resp.Result().StatusCode)
		if want, got := http.StatusNotFound, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}

	})

	t.Run("Incorrect password", func(t *testing.T) {
		body := `{
			"email":"ajith@gmail.com",
			"password":"65432178"
		}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Login successful(By Admin)", func(t *testing.T) {
		body := `{
			"email":"ajith@gmail.com",
			"password":"12345678"
		}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
		str := strings.Split(resp.Body.String(), `"`)
		AdminToken = str[9]
	})

	t.Run("Login successful(By user)", func(t *testing.T) {
		body := `{
		"email":"vijay@gmail.com",
		"password":"12345678"
	}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
		str := strings.Split(resp.Body.String(), `"`)
		UserToken = str[9]
	})
}
func TestPostProduct(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.POST("/admin/postProduct", database.PostProduct, middleware.AuthMiddleware)

	t.Run("Missing token", func(t *testing.T) {
		body := `{
			"brand_name": "dell",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		body := `{
			"brand_name": "dell",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		body := `{
			"brand_name": "dell",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing brand_name", func(t *testing.T) {
		body := `{
			"brand_name": "",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing product_price", func(t *testing.T) {
		body := `{
			"brand_name": "dell",
			"product_price": "",
			"ram_capacity": "2GB",
			"ram_price": "2000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing ram_capacity", func(t *testing.T) {
		body := `{
			"brand_name": "dell",
			"product_price": "20000",
			"ram_capacity": "",
			"ram_price": "2000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing ram_price", func(t *testing.T) {
		body := `{
			"brand_name": "dell",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": ""
		}`
		req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Product added successfully", func(t *testing.T) {
		body := `{
			"brand_name": "dell",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000"
		}`
		for i := 0; i < 2; i++ {
			req := httptest.NewRequest(http.MethodPost, "/admin/postProduct", strings.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
			resp := httptest.NewRecorder()
			e.ServeHTTP(resp, req)
			if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
				t.Fatalf("expected: %d, got: %d", want, got)
			}
		}
	})
}

func TestGetAllProducts(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.GET("/common/getAllProducts", database.GetAllProducts, middleware.AuthMiddleware)
	t.Run("Missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getAllProducts", nil)
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodGet, "/common/getAllProducts", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("All Posts are retrieved successfully(By admin)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getAllProducts", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("All Posts are retrieved successfully(By user)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getAllProducts", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestUpdateProductById(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.PUT("/admin/updateProduct/:product_id", database.UpdateProductById, middleware.AuthMiddleware)

	t.Run("Missing token", func(t *testing.T) {
		body := `{
			"brand_name": "hp"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateProduct/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		body := `{
			"brand_name": "hp"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateProduct/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		body := `{
			"brand_name": "hp"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateProduct/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Post not found", func(t *testing.T) {
		body := `{
			"brand_name": "hp"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateProduct/5", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusNotFound, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Product updated successfully", func(t *testing.T) {
		body := `{
			"brand_name": "hp"
		}`
		//Make sure that the Id(in URL) should present in the database
		req := httptest.NewRequest(http.MethodPut, "/admin/updateProduct/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestDeleteProductById(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.DELETE("/admin/deleteProduct/:product_id", database.DeleteProductById, middleware.AuthMiddleware)

	t.Run("Missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/deleteProduct/1", nil)
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodDelete, "/admin/deleteProduct/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/deleteProduct/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Post not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/admin/deleteProduct/5", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusNotFound, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("post deleted successfully", func(t *testing.T) {
		//While running this case, need to change the Id in URL for every time
		req := httptest.NewRequest(http.MethodDelete, "/admin/deleteProduct/2", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestAddOrder(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.POST("/user/postOrder", database.AddOrder, middleware.AuthMiddleware)

	t.Run("Missing token", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing brand_name", func(t *testing.T) {
		body := `{
			"brand_name": "",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing product_price", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing ram_capacity", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing ram_price", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing name", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing address", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing phone_number", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": ""
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid phone_number", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "923647823"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order added successfully", func(t *testing.T) {
		body := `{
			"brand_name": "hp",
			"product_price": "20000",
			"ram_capacity": "2GB",
			"ram_price": "2000",
			"dvd_rw_drive": true,
			"name": "Hari",
			"address": "5th street",
			"phone_number": "9876543210"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/postOrder", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestGetOrder(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.GET("/common/getOrders", database.GetOrders, middleware.AuthMiddleware)
	t.Run("Missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getOrders", nil)
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodGet, "/common/getOrders", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("All orders are retrieved successfully(By admin)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getOrders", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("All orders are retrieved successfully(By user)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getOrders", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestPayment(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.POST("/user/payment/:order_id", database.Payment, middleware.AuthMiddleware)
	t.Run("Missing token", func(t *testing.T) {
		body := `{
			"payment":"27000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/payment/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		body := `{
			"payment":"27000"
		}`
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodPost, "/user/payment/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		body := `{
			"payment":"25000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/payment/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("missing payment", func(t *testing.T) {
		body := `{
			"payment":""
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/payment/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid amount", func(t *testing.T) {
		body := `{
			"payment":"10000"
		}`
		req := httptest.NewRequest(http.MethodPost, "/user/payment/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Payment Successful", func(t *testing.T) {
		body := `{
			"payment":"25000"
		}`
		//Make sure that the order-Id(in URL) should present in the database
		req := httptest.NewRequest(http.MethodPost, "/user/payment/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestCancelOrderById(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.DELETE("/user/cancelOrder/:order_id", database.CancelOrderById, middleware.AuthMiddleware)
	t.Run("Missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/user/cancelOrder/1", nil)
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodDelete, "/user/cancelOrder/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/user/cancelOrder/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/user/cancelOrder/5", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusNotFound, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order deleted successfully", func(t *testing.T) {
		//Make sure that the order-Id(in URL) should present in the database
		req := httptest.NewRequest(http.MethodDelete, "/user/cancelOrder/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestGetOrderStatusById(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.GET("/common/getOrderStatus/:order_id", database.GetOrderStatusById, middleware.AuthMiddleware)
	t.Run("Missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getOrderStatus/1", nil)
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodGet, "/common/getOrderStatus/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getOrderStatus/5", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusNotFound, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order-status is retrieved successfully(By admin)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/common/getOrderStatus/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order-status is retrieved successfully(By user)", func(t *testing.T) {
		//Make sure that the order-Id(in URL) should present in the database
		req := httptest.NewRequest(http.MethodGet, "/common/getOrderStatus/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestGetAllOrderStatus(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.GET("/admin/getOrderStatuses", database.GetAllOrderStatus, middleware.AuthMiddleware)
	t.Run("Missing token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/getOrderStatuses", nil)
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		req := httptest.NewRequest(http.MethodGet, "/admin/getOrderStatuses", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/getOrderStatuses", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order-status is empty", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/getOrderStatuses", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("All order-status is retrieved successfully", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/getOrderStatuses", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}

func TestUpdateOrderStatusById(t *testing.T) {
	db := driver.TestDbConnection()
	database := Database{Connection: db}
	middleware := middleware.Database{Connection: db}
	e := echo.New()
	e.PUT("/admin/updateStatus/:order_id", database.UpdateOrderStatusById, middleware.AuthMiddleware)

	t.Run("Missing token", func(t *testing.T) {
		body := `{
			"order_status":"shipped"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateStatus/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Invalid token", func(t *testing.T) {
		InvalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFeHBpcmVzQXQiOjE2OTE3MzQ4MjcsIklzc3VlZEF0IjoxNjkxNjQ4ND" +
			"I3LCJSb2xlLWlkIjoiMSIsIlVzZXItaWQiOiIxIn0.lVTEa9Ddpu-EyeXNQZYyGw8JpNeBhgvFt8INc-n-8C"
		body := `{
			"order_status":"shipped"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateStatus/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", InvalidToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusBadRequest, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Unauthorized entry", func(t *testing.T) {
		body := `{
			"order_status":"shipped"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateStatus/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", UserToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusUnauthorized, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order-status not found", func(t *testing.T) {
		body := `{
		    "order_status":"shipped"
		}`
		req := httptest.NewRequest(http.MethodPut, "/admin/updateStatus/5", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusNotFound, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})

	t.Run("Order-status updated successfully", func(t *testing.T) {
		body := `{
			"order_status":"shipped"
		}`
		//Make sure that the order-Id(in URL) should present in the database
		req := httptest.NewRequest(http.MethodPut, "/admin/updateStatus/1", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AdminToken))
		resp := httptest.NewRecorder()
		e.ServeHTTP(resp, req)
		if want, got := http.StatusOK, resp.Result().StatusCode; want != got {
			t.Fatalf("expected: %d, got: %d", want, got)
		}
	})
}
