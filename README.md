# Music Player -G

Music Player - G is a lightweight and fast website-based music player, designed to be used on both local networks and the internet. The player is written using Golang, JS, and HTML and has several features.

## Features

- Spotify-based search functionality
- 160kbps audio quality
- Cross-platform (Windows, Linux, Mac, Android, iOS)
- Lightweight and fast (less than 10mb)

## Installation

To install Music Player -G, follow the steps below:

1. Clone this repository using the command below:

```bash
git clone https://github.com/amarnathcjd/musicplayerg.git
```

2. Build the executable using the command below:

```bash
go build
```


3. Run the executable using the command below:

```bash 
./music-player-g
```

4. Open your preferred web browser and go to localhost:8080.

> The default port is 8080. To change the port, set the $ENV variable port.

## TODO

- Add more audio quality options
- Introduce playlist support
- Implement cookies to save user liked songs
- Add support for more music sources
- Implement UI improvements

## License

This project is licensed under the MIT License.
