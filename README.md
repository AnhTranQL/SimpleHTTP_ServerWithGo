# SimpleHTTP_ServerWithGo


SimPleHTTP_ServerWithGo là chương trình đơn giản về web server dùng 4 phương thức HTTP: Get, Post, Put và Delete sử dụng ngôn ngữ Golang.

Để chạy được chương trình, trước hết bạn phải cài đặt Golang
Chạy chương trình: go run sample1.go


Chương trình có 4 đường dẫn minh họa cho 4 method:


1. GET

Đường dẫn: http://127.0.0.1:3001/helloGet 

2. POST

Xem danh sách của n người đầu tiên trong danh sách, n là số người dùng nhập vào

Đường dẫn: http://127.0.0.1:3001/helloPost

Kiểu dữ liệu nhập trong phương thức Post là JSON, dữ liệu nhập vào là biến Count kiểu int

VD:
{
  "Count" : 2
}

3. PUT

Thêm một người vào trong danh sách

Đường dẫn: http://127.0.0.1:3001/helloPut

Kiểu dữ liệu nhập trong phương thức Put là JSON, dữ liệu nhập vào là struct person

type person struct {
	Id   int
	Name string
	Age  int
}

VD: 
{
	"Id" : 127,
	"Name": "Trong",
	"Age": 25
}


4.DELETE


Xóa một người có id(người dùng nhập) trong list danh sách
Đường dẫn: http://127.0.0.1:3001/helloDelete
Kiểu dữ liệu nhập trong phương thức Post là JSON, dữ liệu nhập vào là biến Count kiểu int

VD:
{
  "Count" : 2
}
Jus for test
