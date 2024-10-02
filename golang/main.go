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
	rows, err := conn.Query("SELECT id,name,email,state,inserttime FROM lienlac")
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
		html += fmt.Sprintf(`<div style="display: flex;" > <input type="checkbox" name="del" value="%d"> <a style="display: flex;" href="edit.html?id=%d"><p>%s</p><p>%s</p><p>%s</p><p>%s</p></a></div>`, id, id, name, email, state, updatetime)
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

// tìm kiếm dữ liệu
func hind_data(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	w.Header().Set("Content-Type", "text/paint")

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
	http.HandleFunc("/hind_data", hind_data)
	//cho sever chạy trên localhost:8000
	log.Println("Server đang chạy tại http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
