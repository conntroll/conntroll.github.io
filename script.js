// function toggleElementVisibilityById(id){
function toggle(id){
  elem = document.getElementById(id);
  disp = elem.style.display;
  if (disp != "block") {
    elem.style.display = "block";
    setCookie("prevtab", id, 7)
  } else {
    elem.style.display = "";
  }
  ["story", "download", "examples", "meaning-of-colors", "stats", "about"].forEach((e)=>{ // "how-it-works", 
    if (e != id) {
      document.getElementById(e).style.display="";
    }
  });
}
// https://stackoverflow.com/questions/14573223/set-cookie-and-get-cookie-with-javascript
function setCookie(name,value,days) {
    var expires = "";
    if (days) {
        var date = new Date();
        date.setTime(date.getTime() + (days*24*60*60*1000));
        expires = "; expires=" + date.toUTCString();
    }
    document.cookie = name + "=" + (value || "")  + expires + "; path=/";
}
function getCookie(name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i=0;i < ca.length;i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
    }
    return null;
}
function eraseCookie(name) {   
    document.cookie = name+'=; Max-Age=-99999999;';  
}
var tab = getCookie("prevtab");
if (tab) {
  toggle(tab)
}
["story", "download", "examples", "meaning-of-colors", "stats", "about"].forEach((e)=>{ // "how-it-works",
  document.getElementById("a-"+e).onclick=(()=>toggle(e));
});

document.getElementById("a-story-examples").onclick=(()=>toggle("examples"));
