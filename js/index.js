const content=document.querySelector(".content"),
name_hind=document.querySelector(".find_by_name"),
state_hind=document.querySelector(".find_by_state"),
del= document.querySelector(".del_bnt");
//lấy dữ liệu
fetchData()

del.onclick=()=>{
    getSelectedCheckboxes_toDel();
}
//ta sẽ kiểm tra hai thành phần input có gì thay đổi hay không
name_hind.onkeyup = ()=>{hindingdata();}
state_hind.onclick = ()=>{hindingdata();}


//hàm tìm kiếm
function hindingdata(){
    let name= name_hind.value;
    let state=document.querySelector(".find_by_state").value;
    //nếu không có dữ liệu đầu vào thì không tìm kiếm nữa
    if (name===""&&state===" "){
        fetchData()
    }
    else{ 
         console.log("name:"+name+" stat:"+state)
          // Tạo FormData để gửi dữ liệu
        let formData = new FormData();
        formData.append("name", name);
        formData.append("state", state);

     fetch("http://localhost:8000/hinddata",{method:"POST",body:formData})
     .then(response=>response.text())
     .then(data=>{
        content.innerHTML=data}
    )
     .catch(err=>{console.error("wrong!!in hinding:"+err)});
    }
}
//lấy dữ liệu đầy từ database
function fetchData(){
    fetch("http://localhost:8000/fetchData",{method:"POST"})
    .then(response=>response.text())
    .then(data=>{content.innerHTML=data})
    .catch(err=>{console.error("fetch data bị lỗi"+err)});
}

//hàm này sẽ xóa các mục được chọn 
function getSelectedCheckboxes_toDel() {
    let checkboxes= document.querySelectorAll('input[name="del"]:checked');//hàm lấy các input
    let selectedValues = [];
    checkboxes.forEach(function (checkbox) {
        selectedValues.push(checkbox.value);
    });

    //đóng gói dữ liệu để gửi đi 
    //chú ý xóa dấu cách
    let id = selectedValues.join(",");
    console.log(id);
    fetch("http://localhost:8000/deletedb", { method: "POST", body: id })
        .then(repost => repost.text())
        .then(data => {
            console.log("Delete succses:" + data);
            fetchData();
        })
        .catch(err => { console.error("Data không fetch") + err })
}