$(function () {
  $('#build').on("click", function (e) {
    console.log("1111111");
   e.preventDefault();
   $.post("http://127.0.0.1/login", function() {
    console.log("post works!!!");
   });
   console.log("toggle build"); 
  });
});

