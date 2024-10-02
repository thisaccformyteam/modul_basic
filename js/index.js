const content=document.querySelector(".content"),
name=document.querySelector(".find_by_name"),
state=document.querySelector(".find_by_state");

fetchData();

function fetchData(){
    fetch("http://localhost:8000/fetchData",{method:"POST"})
    .then(response=>response.text())
    .then(data=>{content.innerHTML=data})
    .catch(err=>{console.error("fetch data bị lỗi"+err)});
}