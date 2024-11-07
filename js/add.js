const form = document.querySelector("form"),
  errTxt = document.querySelector(".error_text"),
  btn1 = document.querySelector("#submit");

form.onsubmit = (e) => {
  e.preventDefault();
};
btn1.onclick = () => {
  addData();
};

function addData() {
  let formdata = new FormData(form);
  let email = formdata.get("email"),
    phone = formdata.get("tel"),
    name = formdata.get("name");
    // kiểm tra điều kiện trước khi update data <không hiểu sao mình code được quả điều kiện đỉnh vl>
 if (!name){
    errTxt.innerHTML="vui lòng nhập tên liên lạc";
    return;
 }
 else if(!validateEmail(email)||!validatePhone(phone)){
    errTxt.innerHTML="email hoặc phone không hợp lệ";
 }
 else sendData(formdata);
}

function sendData(formdata) {
  fetch("http://localhost:8000/dbInsert", { method: "POST", body: formdata })
    .then((response) => response.text())
    .then((data) => {
      errTxt.innerHTML = data;
    })
    .catch((err) => {
      console.error("fetch data bị lỗi" + err);
    });
}

//hai hàm này sẽ kiểm tra định dạng của email và số điện thoại
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
