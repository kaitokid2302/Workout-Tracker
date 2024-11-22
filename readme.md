Hãy hình dung tình huống sau: bạn đã hoàn thành việc phát triển một API hoàn toàn mới, và giờ cần viết tài liệu để hướng dẫn khi xây dựng các ứng dụng client-side sử dụng API đó. Bạn bắt đầu nghĩ về các cách khác nhau để thực hiện điều này và liệt kê ra nhiều lựa chọn như Swagger, Docusaurus, Postman và nhiều hơn nữa.

![](https://blog.logrocket.com/wp-content/uploads/2021/10/Documenting-Go-Web-APIs-Swag.jpg)

Bạn nhớ lại áp lực liên quan đến giai đoạn viết tài liệu API và tự hỏi liệu có cách tắt nào để tăng tốc hay không – bạn không thể bỏ qua giai đoạn này, bởi phần mềm có ích gì nếu không ai có thể sử dụng nó?

Tool yêu thích của tôi để tạo tài liệu API là [Swagger](https://swagger.io/) bởi sự dễ dàng trong việc tạo, duy trì và xuất bản tài liệu API. Swagger là một bộ công cụ mã nguồn mở chuyên nghiệp giúp người dùng, nhóm và doanh nghiệp dễ dàng tạo và tài liệu hóa APIs ở quy mô lớn. Đây là một [demo](https://petstore.swagger.io/) để bạn cảm nhận cách Swagger hoạt động.

Lợi ích của việc sử dụng Swagger trong dự án tiếp theo của bạn bao gồm:

- Cho phép bạn tạo, duy trì và xuất bản tài liệu API một cách nhanh chóng và dễ dàng
- Tạo ra tài liệu tương tác đẹp mắt cho phép bạn xác thực và kiểm thử các endpoint API từ trình duyệt mà không cần phần mềm bên thứ ba
- Dễ dàng hiểu được bởi cả nhà phát triển và người không chuyên
- Chức năng để [generate API client libraries](https://github.com/swagger-api/swagger-codegen) (SDKs) cho nhiều ngôn ngữ và frameworks khác nhau trực tiếp từ một OpenAPI specification

Hướng dẫn này sẽ chỉ cho bạn cách tạo tài liệu Swagger cho Go web APIs trực tiếp từ source code bằng cách sử dụng annotations và [Swag](https://github.com/swaggo/swag). Trong bài viết này, chúng ta sẽ xây dựng một web API demo với [Go và Gin](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/), sau đó tạo tài liệu cho nó bằng Swag.

## Prerequisites

## Yêu cầu trước khi bắt đầu

Để theo dõi và hiểu hướng dẫn này, bạn sẽ cần:

- Kiến thức cơ bản về cách APIs hoạt động
- [Kiến thức cơ bản về Go](https://blog.logrocket.com/getting-started-with-go-for-frontend-developers/)
- Đã cài đặt [Postman](https://www.postman.com/) trên máy của bạn
- Đã cài đặt Go 1.x trên máy của bạn
- Môi trường phát triển Go

## Build a demo Go web API

## Xây dựng một Go web API demo

Gin là web framework đầy đủ tính năng nhanh nhất cho Go, với API giống như [Martini](https://github.com/go-martini/martini) nhấn mạnh vào hiệu suất và năng suất. Gin nhanh, không gặp crash, có thể mở rộng với built-in rendering, và có các tính năng hỗ trợ middleware, grouping routes và quản lý lỗi tiện lợi.

Bây giờ hãy xây dựng web API cho một ứng dụng "to do" cơ bản.

### Step 1: Set up your development environment

### Bước 1: Thiết lập môi trường phát triển

Tạo một dự án Go mới trong trình soạn thảo văn bản hoặc IDE của bạn và khởi tạo file `go.mod`. Bạn có thể sử dụng bất kỳ tên nào cho package của mình:

```go
go mod init swag-gin-demo
```

### Step 2: Install Gin

### Bước 2: Cài đặt Gin

Cài đặt Gin web framework trong dự án của bạn. Trong terminal, nhập:

```go
go get -u github.com/gin-gonic/gin
```

### Step 3: Set up a Gin server

### Bước 3: Thiết lập Gin server

Tạo một file tên `main.go` và lưu đoạn code sau vào:

```go
package main

import (
        "github.com/gin-gonic/gin"
        "net/http"
)

func main() {
        // configure the Gin server
        router := gin.Default()

        // run the Gin server
        router.Run()
}
```

Hãy tạo một kiểu `todo` và khởi tạo danh sách với một số dữ liệu. Thêm đoạn code sau vào file `main.go`:

```go
// todo represents data about a task in the todo list
type todo struct {
        ID   string `json:"id"`
        Task string `json:"task"`
}

// message represents request response with a message
type message struct {
        Message string `json:"message"`
}

// todo slice to seed todo list data
var todoList = []todo{
        {"1", "Learn Go"},
        {"2", "Build an API with Go"},
        {"3", "Document the API with swag"},
}
```

Thêm đoạn code sau vào file `main.go`:

```go
func getAllTodos(c *gin.Context) {
        c.JSON(http.StatusOK, todoList)
}
```

```go
func main() {
        // configure the Gin server
        router := gin.Default()
        router.GET("/todo", getAllTodos)

        // run the Gin server
        router.Run()
}
```

Chạy ứng dụng:

```go
go run main.go
```

### Step 5: Create the `getTodoByID` route

### Bước 5: Tạo route `getTodoByID`

Tạo một route handler sẽ nhận một request `GET` từ client và một `todo ID`, sau đó trả về chi tiết của mục tương ứng từ danh sách todo.

Thêm đoạn code sau vào file `main.go`:

```go
func getTodoByID(c *gin.Context) {
        ID := c.Param("id")

        // loop through todoList and return item with matching ID
        for _, todo := range todoList {
                if todo.ID == ID {
                        c.JSON(http.StatusOK, todo)
                        return
                }
        }

        // return error message if todo is not found
        r := message{"todo not found"}
        c.JSON(http.StatusNotFound, r)
}
```

Thêm route vào router:

```go
router.GET("/todo/:id", getTodoByID)
```

Kiểm thử route bằng Postman:

![Phản hồi Postman cho route getTodoByID](https://blog.logrocket.com/wp-content/uploads/2021/10/Postman-response-getTodoByID.png)

### Step 6: Create the `createTodo` route

### Bước 6: Tạo route `createTodo`

Tạo một route handler sẽ nhận một request `POST` từ client với một `todo ID` và `task`, sau đó thêm một mục mới vào danh sách todo.

Thêm đoạn code sau vào file `main.go`:

```go
func createTodo(c *gin.Context) {
        var newTodo todo

        // bind the received JSON data to newTodo
        if err := c.BindJSON(&newTodo); err != nil {
                r := message{"an error occurred while creating todo"}
                c.JSON(http.StatusBadRequest, r)
                return
        }

        // add the new todo item to todoList
        todoList = append(todoList, newTodo)
        c.JSON(http.StatusCreated, newTodo)
}
```

Thêm route vào router:

```go
router.POST("/todo", createTodo)
```

Kiểm thử route bằng Postman:

![Phản hồi Postman cho route createTodo](https://blog.logrocket.com/wp-content/uploads/2021/10/Postman-response-createTodo.png)

### Step 7: Create the `deleteTodo` route

### Bước 7: Tạo route `deleteTodo`

Tạo một route handler sẽ nhận một request `DELETE` từ client cùng với một `todo ID`, sau đó xóa mục tương ứng khỏi danh sách todo. Thêm đoạn code sau vào file `main.go`:

```go
func deleteTodo(c *gin.Context) {
        ID := c.Param("id")

        // loop through todoList and delete item with matching ID
        for index, todo := range todoList {
                if todo.ID == ID {
                        todoList = append(todoList[:index], todoList[index+1:]...)
                        r := message{"successfully deleted todo"}
                        c.JSON(http.StatusOK, r)
                        return
                }
        }

        // return error message if todo is not found
        r := message{"todo not found"}
        c.JSON(http.StatusNotFound, r)
}
```

Thêm route vào router:

```go
router.DELETE("/todo/:id", deleteTodo)
```

Kiểm thử route bằng Postman:

![Phản hồi Postman cho route deleteTodo](https://blog.logrocket.com/wp-content/uploads/2021/10/Postman-response-deleteTodo.png)

## Document the web API with Swag

## Tài liệu hóa web API với Swag

Swag là middleware giúp tự động tạo tài liệu RESTful API với Swagger 2.0 cho Go trực tiếp từ source code bằng cách sử dụng annotations. Nó yêu cầu bạn chỉ định cách các route hoạt động và tự động hóa toàn bộ quá trình tạo tài liệu Swagger.

Swag tương thích với nhiều [Go web frameworks](https://github.com/swaggo/swag#supported-web-frameworks) và có nhiều tích hợp cho chúng. Hướng dẫn này sẽ sử dụng tích hợp với Gin.

### Step 1: Install Swag

### Bước 1: Cài đặt Swag

Cài đặt package Swag trong dự án của bạn. Trong terminal, nhập:

```go
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

Khởi tạo Swag trong dự án của bạn. Trong terminal, nhập:

```go
swag init
```

Nếu terminal của bạn không nhận ra lệnh `swag init`, bạn cần thêm thư mục `Go bin` vào PATH.

![Cấu trúc thư mục sau khi chạy swag init](https://blog.logrocket.com/wp-content/uploads/2021/10/folder-structure-after-running-swag-init.png)

### Step 3: Import the Swag package into your project

### Bước 3: Import package Swag vào dự án của bạn

Cập nhật phần import trong file `main.go` với đoạn code dưới đây:

```go
import (
        "github.com/gin-gonic/gin"
        swaggerFiles "github.com/swaggo/files"
        ginSwagger "github.com/swaggo/gin-swagger"
        "net/http"
        _ "swag-gin-demo/docs"
)
```

### Step 4: Add general API annotations to the code

### Bước 4: Thêm các annotation API chung vào code

Các [annotation API chung](https://github.com/swaggo/swag#general-api-info) chứa thông tin cơ bản về tài liệu API (tiêu đề, mô tả, phiên bản, thông tin liên hệ, host và giấy phép).

Thêm bộ annotation sau vào file `main.go` (tốt nhất là trước hàm `main`):

```go
// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
```

### Step 5: Add API operation annotations to `controller` code

### Bước 5: Thêm các annotation hoạt động API vào code `controller`

Các annotation hoạt động API chứa cách `controller` hoạt động (mô tả, route, loại request, tham số và mã phản hồi). Hãy xem cách thêm annotation cho route `getAllTodos`.

Thêm các annotation sau ngay trước hàm `getAllTodos` trong file `main.go`:

```go
// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} todo
// @Router /todo [get]
```

Trong đoạn code trên, chúng ta đã định nghĩa:

- `@Summary`, tóm tắt về chức năng của route
- `@ID`, một định danh duy nhất cho route (bắt buộc cho mỗi route)
- `@Produce`, kiểu dữ liệu phản hồi của route
- `@Success 200`, mô hình phản hồi cho mã trạng thái mong đợi
- `@Router /todo [get]`, URI của route và phương thức request được chấp nhận

Tiếp theo, chúng ta sẽ thêm annotation cho route `getTodoByID`. Thêm đoạn code sau ngay trước hàm `getTodoByID` trong file `main.go`:

```go
// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [get]
```

Ở đây, chúng ta chỉ định cho Swag rằng route chấp nhận một tham số bắt buộc kiểu `string` tên `id` được đính kèm vào `path` của request. Nó có tên là `todo ID` với `@Param id path string true "todo ID"`.

Tiếp theo, chúng ta sẽ thêm annotation cho route `createTodo`. Thêm đoạn code sau ngay trước hàm `createTodo` trong file `main.go`:

```go
// @Summary add a new item to the todo list
// @ID create-todo
// @Produce json
// @Param data body todo true "todo data"
// @Success 200 {object} todo
// @Failure 400 {object} message
// @Router /todo [post]
```

Ở đây, chúng ta chỉ định cho Swag rằng route chấp nhận một tham số bắt buộc kiểu `todo` tên `data` được đính kèm vào `body` của request. Nó có tên là `todo data` với `@Param data body todo true "todo data"`.

Chúng ta sẽ thêm annotation cho route `deleteTodo`. Thêm đoạn code sau ngay trước hàm `deleteTodo` trong file `main.go`:

```go
// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /todo/{id} [delete]
```

## View and test the documentation

## Xem và kiểm thử tài liệu

Bây giờ bạn đã định nghĩa tất cả các annotation cho server và các route, hãy xem và kiểm thử tài liệu.

Để tạo tài liệu từ code của bạn, chạy lại lệnh `swag init` trong terminal như sau:

```shell
swag init
```

Chúng ta cũng cần đăng ký một route handler vào router của Gin chịu trách nhiệm render tài liệu Swagger được tạo bởi Swag. Thêm đoạn code sau vào cấu hình router trong `main.go`:

```go
// docs route
router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

Khởi chạy ứng dụng:

```go
go run main.go
```

Truy cập `http://localhost:8080/docs/index.html` trong trình duyệt của bạn để xem tài liệu Swagger:

![Tài liệu cho route getAllTodos](https://blog.logrocket.com/wp-content/uploads/2021/10/Documentation-getAllTodos-route.png)

![Các mô hình dữ liệu Swagger được Swag tài liệu hóa](https://blog.logrocket.com/wp-content/uploads/2021/10/Swagger-data-models-documented-swag.png)

## Conclusion

## Kết luận

Bài viết này đã chỉ cho bạn cách tạo tài liệu Swagger cho web APIs được xây dựng bằng Go một cách liền mạch bằng cách sử dụng Swag. Bạn có thể tìm hiểu thêm về Swag từ [tài liệu chính thức](https://github.com/swaggo/swag).

Chúng tôi đã chọn sử dụng Swagger vì nhiều tính năng và chức năng giúp việc tạo và duy trì tài liệu cho web APIs trở nên dễ dàng.

Source code của web API được xây dựng và tài liệu hóa trong hướng dẫn này có sẵn trên [GitHub](https://github.com/LordGhostX/swag-gin-demo) để bạn khám phá.
