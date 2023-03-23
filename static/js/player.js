let theme = window.matchMedia("(prefers-color-scheme: dark)").matches
  ? "dark"
  : "light";

if (theme !== "dark") {
  document.documentElement.classList.add("dark");
}

document
  .getElementById("dark-mode-toggle")
  .addEventListener("click", function () {
    document.documentElement.classList.toggle("dark");
  });

document.getElementById("song-saved").addEventListener("click", function () {
  document.getElementById("song-saved").classList.toggle("saved");
});

var songs = [
  {
    name: "First Snow",
    artist: "Emancipator",
    album: "Soon It Will Be Cold Enough",
    url: "https://521dimensions.com/song/FirstSnow-Emancipator.mp3",
    cover_art_url:
      "https://521dimensions.com/img/open-source/amplitudejs/album-art/soon-it-will-be-cold-enough.jpg",
  },
];

// Function to get the songs from the server
function getSongs() {
  let searchParams = new URLSearchParams(window.location.search);
  let id = searchParams.get("id");
  $.ajax({
    url: "/get_song?id=" + id,
    type: "GET",
    dataType: "json",
    success: function (data) {
      document.getElementById("player").classList.remove("hidden");
      songs.push({
        name: data.metadata.title,
        artist: data.metadata.artist,
        album: data.metadata.album,
        url: data.link,
        cover_art_url: data.metadata.cover,
      });
      songs.shift();
      Amplitude.init({
        bindings: {
          37: "prev",
          39: "next",
          32: "play_pause",
        },
        callbacks: {
          timeupdate: function () {
            let percentage = Amplitude.getSongPlayedPercentage();

            if (isNaN(percentage)) {
              percentage = 0;
            }

            let slider = document.getElementById("song-percentage-played");
            slider.style.backgroundSize = percentage + "% 100%";
          },
        },
        songs: songs,
      });

      window.onkeydown = function (e) {
        return !(e.keyCode == 32);
      };

      console.log(songs);
    },
    error: function (data) {
      console.log("error");
      console.log(data);
    },
  });
}

getSongs();

search = function () {
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

document.getElementById("search").addEventListener("keyup", search);