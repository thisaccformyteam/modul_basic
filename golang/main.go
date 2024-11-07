package main

import (
	"database/sql"
	"fmt"
	_ "html/template"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Hàm kết nối đến cơ sở dữ liệu Mysql trong xampp
func dbConn() (*sql.DB, error) {
	// dataname:nhacungcap
	// port:127.0.0.1:3379 - port mặc định là 3306
	// password:""
	// user:root
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3379)/nhacungcap") // Không có password
	if err != nil {
		return nil, fmt.Errorf("không kết nối được tới cơ sở dữ liệu: %v", err)
	}
	return db, nil
}

// biến để kết nối đến sever
var conn, err = dbConn()

// hàm để chấp nhận chính sách cros
func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

}

// hàm insert dữ liệu
func dbInsert(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "text/paint")
	//kiểm tra xem phương thức truyền vào có hợp lệ hay không
	if r.Method != http.MethodPost {
		fmt.Fprint(w, "phương thức không hợp lệ")
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	gender := r.FormValue("gender")
	company := r.FormValue("company")
	state := r.FormValue("state")
	email := r.FormValue("email")
	tel := r.FormValue("tel")
	department := r.FormValue("department")
	position := r.FormValue("position")
	now := time.Now()
	// hiệu chỉnh biến now để có thể cho vào database

	formattedTime := now.Format("2006-01-02 15:04:05")
	if name == "" {
		fmt.Fprint(w, "vui lòng nhập tên liên lạc")
		return
	}
	query := "INSERT INTO `lienlac`( `name`, `address`, `gender`, `company`, `state`, `email`, `phone`, `department`, `position`, `inserttime`,`updatetime`) VALUES ('" + name + "','" + address + "','" + gender + "','" + company + "','" + state + "','" + email + "','" + tel + "','" + department + "','" + position + "','" + formattedTime + "','" + formattedTime + "')"
	fmt.Println(query)
	_, err := conn.Exec(query)
	if err != nil {
		fmt.Fprint(w, "lỗi khi thêm đối tượng")
		return
	}
	// khi thành công thì dòng này sẽ được trả về
	fmt.Fprint(w, "thêm dữ liệu thành công")
}

// hàm này dùng để hiển thị dữ liệu đã được định dạng ở index
func fetchData(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "text/paint")
	rows, err := conn.Query("SELECT id,name,email,state,updatetime FROM lienlac")
	if err != nil {
		log.Println("Lỗi truy vấn dữ liệu")
		http.Error(w, "Lỗi truy vấn dữ liệu", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	//tạo chuỗi html
	html := ""
	for rows.Next() {
		var id int
		var name, state, email, updatetime string
		if err := rows.Scan(&id, &name, &email, &state, &updatetime); err != nil {
			log.Printf("Lỗi khi đọc dữ liệu từ hàng: %v", err)
			http.Error(w, "Lỗi khi đọc dữ liệu", http.StatusInternalServerError)
			return
		}
		html += fmt.Sprintf(`<div class="item" > <input type="checkbox" name="del" value="%d"> <a style="display: flex;" href="edit.html?id=%d"><p>%s</p><p>%s</p><p>%s</p><p>%s</p></a></div>`, id, id, name, email, state, updatetime)
	}

	if rows.Err() != nil {
		log.Printf("Lỗi khi duyệt kết quả: %v", rows.Err())
		http.Error(w, "Lỗi duyệt dữ liệu", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, html)
}

// hàm xóa dữ liệu getgo!!
func deletedb(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "đầu vào func delete bị lỗi")
	}
	id := string(body)
	query := "DELETE FROM lienlac WHERE id IN (" + id + ")"
	fmt.Println(query)
	_, err = conn.Exec(query)
	if err != nil {
		fmt.Fprint(w, "truy vẫn lỗi")
		return
	}
	fmt.Fprint(w, "xóa thành công")
}

// đếm số bản trả về
// func checkRowCount(db *sql.DB, query string) (int, error) {
// 	var count int

// 	// Điều chỉnh query để sử dụng COUNT(*)
// 	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM (%s) AS subquery", query)

// 	// Thực thi truy vấn để đếm số hàng
// 	err := db.QueryRow(countQuery).Scan(&count)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return count, nil
// }

// tìm kiếm dữ liệu
func hinddata(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "text/paint")
	namehind := r.FormValue("name")
	statess := r.FormValue("state")
	// Chuẩn bị câu truy vấn
	var query string
	var rows *sql.Rows
	var err error

	if namehind == "" {
		query = "SELECT id, name, email, state, updatetime FROM lienlac WHERE state LIKE ?"
		rows, err = conn.Query(query, "%"+statess+"%")
	} else if statess == "" {
		query = "SELECT id, name, email, state, updatetime FROM lienlac WHERE name LIKE ?"
		rows, err = conn.Query(query, "%"+namehind+"%")
	} else {
		query = "SELECT id, name, email, state, updatetime FROM lienlac WHERE state LIKE ? AND name LIKE ?"
		rows, err = conn.Query(query, "%"+statess+"%", "%"+namehind+"%")
	}

	if err != nil {
		fmt.Print(w, `nah bro some thing stupit is go wrong`)
	}
	defer rows.Close()

	// Tạo chuỗi HTML
	var requestText string = ""

	for rows.Next() {
		var id int
		var name, state, email, updatetime string
		// Quét các cột trong hàng hiện tại vào biến
		if err := rows.Scan(&id, &name, &email, &state, &updatetime); err != nil {
			log.Printf("Lỗi khi đọc dữ liệu từ hàng: %v", err)
			http.Error(w, "Lỗi khi đọc dữ liệu", http.StatusInternalServerError)
			return
		}
		requestText += fmt.Sprintf(
			`<div style="display: flex;">
				<input type="checkbox" name="del" value="%d">
				<a style="display: flex;" href="edit.html?id=%d">
					<p>%s</p><p>%s</p><p>%s</p><p>%s</p>
				</a>
			</div>`, id, id, name, email, state, updatetime)
	}

	fmt.Println(requestText)
	// Trả về kết quả HTML
	fmt.Fprint(w, requestText)
}

// edit showing dât
func editUser(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	// Lấy ID của người dùng từ query string
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Thiếu ID người dùng", http.StatusBadRequest)
		return
	}
	// Truy vấn lấy thông tin người dùng theo ID
	var name, address, gender, company, state, email, phone, department, position string
	err = conn.QueryRow("SELECT `name`, `address`, `gender`, `company`, `state`, `email`, `phone`, `department`, `position`  FROM lienlac WHERE id = ?", id).Scan(&name, &address, &gender, &company, &state, &email, &phone, &department, &position)
	if err != nil {
		http.Error(w, "Không tìm thấy người dùng", http.StatusNotFound)
		return
	}

	// Hiển thị form HTML với các giá trị đã có trong database
	html := `		
	<input type="hidden" name="id" value="` + id + `">
	 <label for="name">tên liên lạc</label>
        <input type="text" name="name" id="name" value="` + name + ` "required>
        <label for="address">địa chỉ</label>
        <input type="text" name="address" id="address" value="` + address + `">
        <label for="gender">giới tính</label>
        <input type="radio" name="gender" value="nam" id="gender">nam
        <input type="radio" name="gender" value="nu" id="gender">nữ
        <input type="radio" name="gender" value="other" id="gender"> other
        <label for="company">tên công ty</label>
        <input type="text" name="company" id="company" value="` + company + `">
        <label for="state">trạng thái</label>
        <select name="state" id="state">
            <option value="thụ động">Thụ động</option>
            <option value="đang mở">đang mở</option>
            <option value="đã phản hồi">đã phản hồi</option>
        </select>
        <label for="email">Email</label>
        <input type="email" name="email" id="email" value="` + email + `">
        <label for="phone">điện thoại</label>
        <input type="tel" name="tel" id="tel"value="` + phone + `">
        <label for="department">Phòng ban</label>
        <input type="text" name="department" id="department"value="` + department + `">
        <label for="position">Chức vụ</label>
        <input type="text" name="position"id="position"value="` + position + `">	
		`
	// Trả về trang HTML
	w.Header().Set("Content-Type", "text/paint")
	fmt.Fprint(w, html)
}

// update du lieu
func updateUser(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	id := r.FormValue("id")
	name := r.FormValue("name")
	address := r.FormValue("address")
	gender := r.FormValue("gender")
	company := r.FormValue("company")
	state := r.FormValue("state")
	email := r.FormValue("email")
	tel := r.FormValue("tel")
	department := r.FormValue("department")
	position := r.FormValue("position")
	now := time.Now()
	// hiệu chỉnh biến now để có thể cho vào database
	formattedTime := now.Format("2006-01-02 15:04:05")
	_, err = conn.Exec("UPDATE `lienlac` SET `name`=?,`address`=?,`gender`=?,`company`=?,`state`=?,`email`=?,`phone`=?,`department`=?,`position`=?,`updatetime`=? WHERE id=?", name, address, gender, company, state, email, tel, department, position, formattedTime, id)
	if err != nil {
		fmt.Fprint(w, "cập nhật thông tin thất bại")
		return
	}
	//thông báo cập nhật
	fmt.Fprintf(w, "Cập nhật thành công người dùng có ID: %s", id)
}
func main() {
	defer conn.Close()
	//api insert data
	http.HandleFunc("/dbInsert", dbInsert)
	//không cần hiểu hàm này
	http.HandleFunc("/options", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(w)
	})
	//api xóa dl
	http.HandleFunc("/deletedb", deletedb)
	//api fetchdata
	http.HandleFunc("/fetchData", fetchData)
	//api truy vấn dữ liệu
	http.HandleFunc("/hinddata", hinddata)
	//api edit
	http.HandleFunc("/edit", editUser)
	//cho sever chạy trên localhost:8000
	http.HandleFunc("/updateUser", updateUser)
	log.Println("Server đang chạy tại http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
