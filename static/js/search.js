search = function () {
  searchBarNeon();
  var query = document.getElementById("search").value;
  if (query.length % 2 == 0 && query.length > 2) {
    $.ajax({
      url: "/search",
      type: "GET",
      data: {
        query: query,
        html: true,
      },
      success: function (data) {
        $("#results").html(data);
        $("title").html("Music Player - G:" + query);
      },
      error: function (data) {
        console.log("error");
        console.log(data);
      },
    });
  }
};

var colorChange = false;

searchBarNeon = function () {
  var query = document.getElementById("search");
  colorChange = true;
  function getRandomColor() {
    var letters = "0123456789ABCDEF";
    var color = "#";
    for (var i = 0; i < 6; i++) {
      color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
  }
  function changeColor() {
    if (colorChange) {
      query.style.borderColor = getRandomColor();
    }
  }
  setInterval(changeColor, 2000);
  after = setTimeout(endSearchBarNeon, 10000);
};

endSearchBarNeon = function () {
  var query = document.getElementById("search");
  colorChange = false;
  query.style.color = "white";
};

document.getElementById("search").addEventListener("keyup", search);
