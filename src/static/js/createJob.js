$(document).ready(function () {
  $('#btCreateJob').click(function () {
    /*
    var jobName = ,
        description = ,
        repositryUrl = ,
        userName = ,
        password = ,
        branches = ,
        registryUrl = ,
        repositryname = ,
        tag = ,
        command = ,
        dockerHostUrl = 
    if ($('#skippush').attr('checked') == undefined) {
      console.log("undefined");
      $('#skippush').attr("value", 'false');
    } else {
      console.log("definded");
      $('#skippush').attr("value", 'true');
    }
*/
    var job = {
      "jobname": $('#jobName').val(),
      "description": $('#description').val(),
      "repositryurl": $('#repositryUrl').val(),
      "username": $('#userName').val(),
      "password": $('#password').val(),
      "branches": $('#branches').val(),
      "dockerregistryurl": $('#registryUrl').val(),
      "repositryname": $('#repositryname').val(),
      "tag": $('#tag').val(),
      "command": $('#command').val(),
      "dockerhosturi": $('#dockerHostUrl').val(),
      "skippush": $("#skippush").attr("value")
    };
//    console.log(job)
    $.ajax({
      type: "POST",
      url: "/createjob",
      data: JSON.stringify(job),
      success: function(data, textStatus, jqXHR) {
        console.log(data);
        console.log(textStatus);
        console.log(jqXHR);
 //       if(data.msg == "true")
//        {
          alert("success");
          window.location.reload();
 //       } else {
      //    alert(data);
 //       }
      },
      error: function() {
        console.log("error");
      }
    });
  });
});
