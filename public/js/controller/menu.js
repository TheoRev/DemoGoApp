$(document).ready(function() {
  if (document.getElementById('index')) {
    $("#header").load("views/templates/menu.html");
  }else {
    $("#header").load("../views/templates/menu2.html");
  }
});
