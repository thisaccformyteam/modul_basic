const form=document.querySelector("form"),
errTxt=document.querySelector(".error_text"),
btn=document.querySelector("#sumbit");

form.onsubmit=(e)=>{
    e.preventDefault();
};
btn.onclick=()=>{
    addData();
}

function addData(){
let formdata= new FormData(form);
fetch("http://localhost:8000/dbInsert",{method:"POST",body:formdata})
.then(response=>response.text())
.then(data=>{errTxt.innerHTML=data})
.catch(err=>{console.error("fetch data bị lỗi"+err)});
}