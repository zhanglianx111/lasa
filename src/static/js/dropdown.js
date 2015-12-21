/*
<div class="dropdown">
<button class="btn dropdown-toggle" type="button" id="dropdownMenu1" data-toggle="dropdown"> Action <span class="caret"></span>
</button>
<ul class="dropdown-menu text-right" role="menu" aria-labelledby="dropdownMenu1">
  <li role="presentation"><a role="menuitem" tabindex="-1" href="#" id="build">Build</a></li>
  <li role="presentation"><a role="menuitem" tabindex="-1" href="#" id="delete">Delete</a></li>
  <li role="presentation" class="divider"></li>
  <li role="presentation"><a role="menuitem" tabindex="-1" href="#" id="config">Job Configuration</a></li>
</ul>
</div>

*/

<script>
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

</script>