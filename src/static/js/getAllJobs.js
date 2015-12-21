//$(function() {
$(document).ready(function () {
	console.log("mabide");
	jQuery.get("http://127.0.0.1:3000/api/job/alljobs", {}, success, "");

	function success(data, textStatus, jqXHR)
	{
		//console.log(data[0]);
		for (id in data) {
			//console.log(data[id].name);
			var name = data[id].name;
			var url = data[id].url;
		}
		addrows(data);
	};

	function addrows(data)
	{
		var dataLength = data.length;
		var tbody=document.getElementById("tbody");

		for (i=1; i<=dataLength; i++) {
			var tr = tbody.insertRow();
			tr.id = i;
      var jobname = data[i-1].name;
      // No.
			var td0 = tr.insertCell(0);
			td0.innerHTML = tr.id;
      // job name
			var td1 = tr.insertCell(1);
			td1.innerHTML = jobname;
      // url
			var td2 = tr.insertCell(2);
			td2.innerHTML = jobname;
      // actions
			var td3 = tr.insertCell(3);
      var text = "<div class=\"dropdown\"><button class=\"btn dropdown-toggle\" type=\"button\" id=\"dropdownMenu1\" data-toggle=\"dropdown\"> Action <span class=\"caret\"></span></button><ul class=\"dropdown-menu text-right\" role=\"menu\" aria-labelledby=\"dropdownMenu1\"><li role=\"presentation\"><a role=\"menuitem\" tabindex=\"-1\" id=\"build\" onclick=\"mybuild(this)\">Build</a></li><li role=\"presentation\"><a role=\"menuitem\" tabindex=\"-1\" onclick=\"mydelete(this)\" >Delete</a></li><li role=\"presentation\" class=\"divider\"></li><li role=\"presentation\"><a role=\"menuitem\" tabindex=\"-1\" onclick=\"jobconfig(this)\">Job Configuration</a></li></ul></div>";
			td3.innerHTML = text;

      // process
      var td4 = tr.insertCell(4);
      //td4.innerHTML = "<div class=\"progress progress-striped active\"><div class=\"bar\" style=\"width: 90%;\"></div></div>";
      //td4.innerHTML = "<div class=\"progress\"><div class=\"progress-bar progress-bar-striped active\" role=\"progressbar\" aria-valuenow=\"45\" aria-valuemin=\"0\" aria-valuemax=\"100\" style=\"width: 45%\"><span class=\"sr-only\">45% Complete</span></div></div>";
      td4.innerHTML = "<div class=\"progress\"><div class=\"progress-bar progress-bar-striped active\" role=\"progressbar\" aria-valuenow=\"10\" aria-valuemin=\"0\" aria-valuemax=\"100\" style=\"min-width: 1em; width: 100%\">100%</div></div>";
		};
	};

  /*
  function addDropdown()
  {
      var eDiv = document.createElement("div");
  eDiv.className = "dropdown";

  var eBt = document.createElement("button");
  eBt.className = "btn dropdown-toggle";
  eBt.setAttribute("type", "button");
  eBt.setAttribute("id", "dropdownMenu1");
  eBt.setAttribute("data-toggle", "dropdown");
  eBt.innerHTML = "Action";
  var eSpan = document.createElement("span");
  eSpan.className = "caret";
  eBt.appendChild(eSpan);
  eDiv.appendChild(eBt);

  var eUl = document.createElement("ul");
  eUl.className = "dropdown-menu text-right";
  eUl.setAttribute("role","menu");
  eUl.setAttribute("aria-labelledby","dropdownMenu1");
  
  var eLiA = document.createElement("li");
  eLiA.setAttribute("role", "presentation");

  var eA = document.createElement("a");
  eA.setAttribute("role", "menuitem");
  eA.setAttribute("tabindex", "-1");
  eA.setAttribute("href", "#");
  eA.innerHTML = "Build";

  eLiA.appendChild(eA);

  var eLiB = eLiA.cloneNode(true);

  var eLiC = document.createElement("li");
  eLiB.setAttribute("role", "presentation");
  eB.className = "divider";

  var eLiD = eLiA.cloneNode(true);
  eLiD.innerHTML = "Job Configuration";

  eUl.appendChild(eLiA);
  eUl.appendChild(eLiB);
  eUl.appendChild(eLiC);
  eUl.appendChild(eLiD);

  eDiv.appendChild(eUl);
  return eDiv
  };
  */
});
