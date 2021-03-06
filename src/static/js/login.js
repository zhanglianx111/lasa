$( document ).ready(function() {
    // DOM ready

    // Test data
    /*
     * To test the script you should discomment the function
     * testLocalStorageData and refresh the page. The function
     * will load some test data and the loadProfile
     * will do the changes in the UI
     */
    // testLocalStorageData();
    // Load profile if it exits
    //loadProfile();
    // login button
    $('#btLogin').click(function() {
    console.log("login");
    var user= $('#inputUsername').val();
    var passwd = $('#inputPassword').val();
    console.log(user);
    console.log(passwd);
    var usrData = {
    	"user": user,
    	"passwd": passwd
    };
    $.ajax({
    	type: "POST",
    	url: "/login",
    	data: usrData,
    	success: success,
    	error: function() {
    		alert("login failed, please check you username and password");
    		return false;
    	}
    });
    function success() {
    	console.log("login success");
    	location.href="/dashboard";
    };
	});
    //register button
    $('#btRegister').click(function() {
    console.log("register");
      $.ajax({
    	type: "get",
    	url: "/register",
    	success: success,
    	error: function() {
    		alert("register failed");
    		return false;
    	}
    });
    function success() {
    	console.log("register success");
    	location.href="/register";
    };
	});
});











/**
 * Function that gets the data of the profile in case
 * thar it has already saved in localstorage. Only the
 * UI will be update in case that all data is available
 *
 * A not existing key in localstorage return null
 *
 
function getLocalProfile(callback){
    var profileImgSrc      = localStorage.getItem("PROFILE_IMG_SRC");
    var profileName        = localStorage.getItem("PROFILE_NAME");
    var profileReAuthEmail = localStorage.getItem("PROFILE_REAUTH_EMAIL");
    console.log(profileName);
    if(profileName !== null
            && profileReAuthEmail !== null
            && profileImgSrc !== null) {
        callback(profileImgSrc, profileName, profileReAuthEmail);  
    }
}
*/
/**
 * Main function that load the profile if exists
 * in localstorage

function loadProfile() {
    if(!supportsHTML5Storage()) { return false; }
    // we have to provide to the callback the basic
    // information to set the profile
    console.log("loadProfile");
    getLocalProfile(function(profileImgSrc, profileName, profileReAuthEmail) {
        //changes in the UI
        $("#profile-img").attr("src",profileImgSrc);
        $("#profile-name").html(profileName);
        $("#reauth-email").html(profileReAuthEmail);
        $("#inputEmail").hide();
        $("#remember-me").hide();
    });
}
 */
/**
 * function that checks if the browser supports HTML5
 * local storage
 *
 * @returns {boolean}

function supportsHTML5Storage() {
    try {
        return 'localStorage' in window && window['localStorage'] !== null;
    } catch (e) {
        return false;
    }
}
 */
/**
 * Test data. This data will be safe by the web app
 * in the first successful login of a auth user.
 * To Test the scripts, delete the localstorage data
 * and comment this call.
 *
 * @returns {boolean}
 
function testLocalStorageData() {
    if(!supportsHTML5Storage()) { return false; }
    localStorage.setItem("PROFILE_IMG_SRC", "//lh3.googleusercontent.com/-6V8xOA6M7BA/AAAAAAAAAAI/AAAAAAAAAAA/rzlHcD0KYwo/photo.jpg?sz=120" );
    localStorage.setItem("PROFILE_NAME", "César Izquierdo Tello");
    localStorage.setItem("PROFILE_REAUTH_EMAIL", "oneaccount@gmail.com");
}
*/
