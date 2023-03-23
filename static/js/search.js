search = function () {
  searchBarNeon();
  var query = document.getElementById("search").value;
  if (query.length % 2 == 0 && query.length > 2) {
    if (
      $("#search").html() ==
      `<p class="text-center">Enter a query to search for a song</p>`
    ) {
      $("#search").html(
        "<span class='text-center text-lg'>Searching...</span>"
      );
    }
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
        if (data.status == 404) {
          $("#results").html(
            "<div class='container'><div class='row'><div class='text-center text-lg'><h1>No results found</h1></div></div></div>"
          );
        } else {
          console.log("error");
        }
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

var getTrending = function () {
  $.ajax({
    url: "/get_top",
    type: "GET",
    data: {
      html: true,
    },
    success: function (data) {
      $("#results").html("<h1>This Week's Top Songs</h1>" + data);
    },
    error: function (data) {
      if (data.status == 404) {
        $("#results").html(
          "<div class='container'><div class='row'><div class='text-center text-lg'><h1>Failed to get top songs</h1></div></div></div>"
        );
      } else {
        console.log("error");
      }
    },
  });
};

getTrending();

