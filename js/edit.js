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
  let formData = new FormData(form);
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
