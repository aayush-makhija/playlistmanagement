
---

# Music Management System 🎵

This project is a comprehensive Music Management System built with Go (Golang) and PostgreSQL. It provides functionalities to manage albums, artists, songs, and playlists efficiently. The application leverages the `database/sql` package in Go to interact seamlessly with a PostgreSQL database.

## Features

- **Albums Management**: Create, read, update, and delete albums with attributes like name, tracks, and artist.
- **Artists Management**: Add and manage artists seamlessly.
- **Songs Management**: Handle songs with details like song ID, name, artist, and album.
- **Playlists Management**: Create playlists, add songs to playlists, and manage them effectively.
- **User Interaction**: User-friendly command-line interface for adding and managing songs in playlists.

## Project Structure

- `main.go`: The main entry point of the application, handling database connections and user interactions.
- `models.go`: Contains the definitions of the Album, Artist, Songs, and Playlist structs.
- `db.go`: Includes functions to interact with the PostgreSQL database for various CRUD operations.

## Getting Started

### Prerequisites

- Go (1.15+)
- PostgreSQL

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/your-username/music-management-system.git
    cd music-management-system
    ```

2. Set up the PostgreSQL database and update the connection details in `main.go`:
    ```go
    const (
        host     = "localhost"
        port     = 5432
        user     = "your-username"
        password = "your-password"
        dbname   = "stopify"
    )
    ```

3. Install Go dependencies:
    ```sh
    go get -u github.com/lib/pq
    ```

4. Run the application:
    ```sh
    go run main.go
    ```

## Usage

- Follow the prompts to add, view, and manage albums, artists, songs, and playlists.
- Example interactions include creating albums, adding songs to playlists, and deleting playlists.

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to customize the repository name, installation instructions, and other details based on your setup and preferences.
