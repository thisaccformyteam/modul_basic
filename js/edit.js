
const form = document.querySelector("form"),
  btn = document.querySelector(".sumbit"),
  err = document.querySelector(".error_text");
form.onsubmit = (e) => {
  e.preventDefault();
};
fetch_edit();
btn.onclick=()=>{
    updatedata()
}

// hàm này lấy form được xử lý từ backend sau đó hiển thị ra màn hình khi load trang
function fetch_edit() {
  const urlParams = new URLSearchParams(window.location.search);
  const userId = urlParams.get("id");
  fetch("http://localhost:8000/edit?id=" + userId)
    .then((response) => response.text())
    .then((data) => {
      document.querySelector("form").innerHTML = data;
      console.log("load form sussecs ");
    });
}
function updatedata() {
  let formdata = new FormData(form);
  let email = formdata.get("email"),
  phone = formdata.get("tel"),
  name = formdata.get("name");
  if(!name){
    err.innerHTML="vui lòng nhập tên";
  }
  else if(!validateEmail(email)||!validatePhone(phone)){
    err.innerHTML="email hoặc phone không hợp lệ";
 }
 else fetch_Update(formdata);
}
function fetch_Update(formData){
   fetch("http://localhost:8000/updateUser", { method: "POST", body: formData })
    .then((re) => re.text())
    .then((data) => {
      err.innerHTML = data;
      console.log("yeah");
    })
    .catch((err) => {
      console.log("lỗi" + err);
    });
}
/* hai hàm này để kiểm tra định dạng của email và phone 
  -- vì cái chính sách cros không cho phép lấy nguồn trực tiếp trong thư mục nên là phải copy lại file 
  -- và để tránh mấy cái bugs suất hiện một cách vớ vẩn thì tạm thời để nó ở đây
  -- nếu dùng server để chạy thì cros sẽ không xuất hiện.
*/
function validateEmail(email) {
  if (email) {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(email);
  }
  return true;
}
 function validatePhone(phone) {
  if (phone) {
    const regex = /^(0[3|5|7|8|9])+([0-9]{8})$/;
    return regex.test(phone);
  }
  return true;
}