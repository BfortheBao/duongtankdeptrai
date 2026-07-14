hello moi nguoi
## 5. YÊU CẦU VỀ TESTING

### 5.1 Unit Test
- Validate dữ liệu đầu vào (email, điểm số, sĩ số)
- Hàm tính điểm trung bình
- Hàm hash/verify password
- Hàm tạo/verify JWT
- Logic phân quyền theo role
- Coverage tối thiểu 60–70% cho tầng nghiệp vụ (service/business logic)

### 5.2 Repository/Data Access Test
- Test CRUD từng bảng
- Test ràng buộc khóa ngoại, unique constraint
- Không test trên dữ liệu production

### 5.3 Service/Business Logic Test
- Test độc lập với tầng dữ liệu thật (mock/giả lập)
- Test các case nghiệp vụ: thêm học sinh vào lớp đã đầy, điểm danh trùng ngày, nhập điểm ngoài khoảng hợp lệ, học sinh không thuộc lớp

### 5.4 API/Integration Test
- Test status code, response body cho từng endpoint
- Test middleware: không có token → 401, sai role → 403
- Test luồng nhiều bước (đăng nhập → thao tác → kiểm tra kết quả)

### 5.5 Regression Test
- Có bộ regression test chạy được độc lập, không phụ thuộc vào việc chạy tay
- Bộ regression test bao phủ toàn bộ luồng nghiệp vụ chính qua API thật (không mock):
  - Đăng ký → đăng nhập → lấy token
  - Tạo lớp học → thêm học sinh → kiểm tra sĩ số tối đa
  - Điểm danh → kiểm tra trùng ngày
  - Nhập điểm → kiểm tra ràng buộc điểm hợp lệ
  - Kiểm tra phân quyền: student gọi API admin → bị từ chối
  - Kiểm tra không có token → bị từ chối
- Nếu hệ thống có thêm giao diện web (frontend) sau này: bổ sung test luồng người dùng thật trên trình duyệt (đăng nhập, điểm danh, nhập điểm, v.v.)
- Bộ regression test phải chạy tự động được trong pipeline CI, độc lập với mã nguồn backend

### 5.6 Checklist test theo module
- [ ] Auth: đăng ký trùng username, sai password, token hết hạn
- [ ] Class: tạo lớp thiếu tên, thêm học sinh khi lớp đầy, xóa lớp có học sinh
- [ ] Attendance: điểm danh trùng ngày, học sinh không thuộc lớp
- [ ] Score: điểm âm/ngoài khoảng hợp lệ, học sinh không tồn tại
- [ ] Middleware: không có token, token sai, sai role
- [ ] Regression: toàn bộ luồng nghiệp vụ end-to-end qua API

### 5.7 Yêu cầu về quy trình
- Test phải chạy tự động khi có thay đổi mã nguồn (CI)
- Báo cáo kết quả test và coverage phải xem được sau mỗi lần chạy
- Test phải chạy được trên môi trường riêng biệt với dữ liệu thật

---

## 6. LỘ TRÌNH PHÁT TRIỂN

1. **Giai đoạn 1**: Nền tảng — CRUD cơ bản (lớp học, học sinh), chưa auth
2. **Giai đoạn 2**: Xác thực & phân quyền
3. **Giai đoạn 3**: Nghiệp vụ chính — điểm danh, nhập điểm, quan hệ nhiều-nhiều
4. **Giai đoạn 4**: Hoàn thiện validate, xử lý lỗi, phân trang + unit test, repository test
5. **Giai đoạn 5**: Integration test + regression test
6. **Giai đoạn 6**: Container hóa, CI/CD, mở rộng (thông báo, thời khóa biểu, export báo cáo)

---

## 7. TIÊU CHÍ HOÀN THÀNH

- Chạy được toàn bộ API CRUD cơ bản
- Phân quyền hoạt động đúng theo role
- Dữ liệu có ràng buộc khóa ngoại và unique hợp lý
- Kiến trúc mã nguồn tách lớp rõ ràng
- Có đầy đủ unit test, repository test, integration test
- Có bộ regression test chạy độc lập, tích hợp CI
- Coverage tối thiểu 60–70% cho tầng service/business logic
- Toàn bộ hệ thống chạy được bằng container hóa



--------------------------------

buihin bat dau viet:
package main

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(Sum(3, 5))
}

