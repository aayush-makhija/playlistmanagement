type Album struct {
	Name   string
	Tracks int
	artist string
}

type Songs struct {
	songid int
	Name   string
	Artist string
	Album  string
}

type Songs1 struct {
	songid int
	name   string
}

type Artist struct {
	Name string
}

type Playlist struct {
	playlistID int
	Name       string
}

func createAlbum(db *sql.DB, album Album) int {
	query := `INSERT INTO albums (name, tracks, artist) VALUES($1, $2, $3) RETURNING albumid`
	var pk int
	err := db.QueryRow(query, album.Name, album.Tracks, album.artist).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk

}
func createArtist(db *sql.DB, artist Artist) int {
	query := `INSERT INTO artist(name) VALUES($1) RETURNING artistid`
	var pk int
	err := db.QueryRow(query, artist.Name).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func createPlaylist(db *sql.DB, playlist Playlist) int {
	query := `INSERT INTO playlist(name) VALUES($1) RETURNING playlistid`
	var pk int
	err := db.QueryRow(query, playlist.Name).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func createSongs(db *sql.DB, song Songs) int {
	query := `INSERT INTO songs(name,artist,album) VALUES($1, $2, $3) RETURNING songid`
	var pk int
	err := db.QueryRow(query, song.Name, song.Artist, song.Album).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func addSongToPlaylist(db *sql.DB) {
	songList := []Songs{}
	rows, err := db.Query(`SELECT * FROM songs`)
	if err != nil {
		log.Fatal(err)
	}
	var songid int
	var name, artist, album string

	for rows.Next() {
		err := rows.Scan(&songid, &name, &artist, &album)
		if err != nil {
			log.Fatal(err)

		}
		songList = append(songList, Songs{songid, name, artist, album})
	}
	for i := 0; i < len(songList); i++ {
		fmt.Println("Song id:", songList[i].songid, "| Name: ", songList[i].Name, "| Album: ", songList[i].Album, "| Artist: ", songList[i].Artist)
	}
	defer rows.Close()
	fmt.Println("Enter the songid written next to the song you want in your playlist")
	var request int
	fmt.Scanln(&request)

	playlistList := []Playlist{}
	rows, err = db.Query(`SELECT * FROM playlist`)
	if err != nil {
		log.Fatal(err)
	}
	var playlistID int
	var playlistName string

	for rows.Next() {
		err = rows.Scan(&playlistID, &playlistName)
		if err != nil {
			log.Fatal(err)
		}
		playlistList = append(playlistList, Playlist{playlistID, playlistName})
	}
	for i := 0; i < len(playlistList); i++ {
		fmt.Println("Playlist ID: ", playlistList[i].playlistID, "| Name: ", playlistName)
	}
	defer rows.Close()
	fmt.Println("Enter the playlistID written next to the playlist you want the song to be in: ")
	var destination int
	fmt.Scanln(&destination)

	query1 := `INSERT INTO playlist_songs(playlistid, songid) VALUES ($1, $2)`
	db.QueryRow(query1, destination, request)
	fmt.Println("Task completed.")

}

func deleteFromPlaylist(db *sql.DB) {
	playlistList := []Playlist{}
	rows, err := db.Query(`SELECT playlist_songs.playlistid, name FROM playlist_songs JOIN playlist ON playlist_songs.playlistid = playlist.playlistid `)
	if err != nil {
		log.Fatal(err)
	}
	var playlistID int
	var playlistName string

	for rows.Next() {
		err = rows.Scan(&playlistID, &playlistName)
		if err != nil {
			log.Fatal(err)
		}
		playlistList = append(playlistList, Playlist{playlistID, playlistName})
	}
	for i := 0; i < len(playlistList); i++ {
		fmt.Println("Playlist ID: ", playlistList[i].playlistID, "| Name: ", playlistName)
	}
	defer rows.Close()
	fmt.Println("Which playlist(ID) would you like to delete the song from : ")
	var destination int
	fmt.Scanln(&destination)

	songList := []Songs1{}
	query := (`SELECT songs.songid, name FROM songs JOIN playlist_songs ON playlist_songs.songid = songs.songid WHERE playlistid = $1`)
	rows, err = db.Query(query, destination)
	if err != nil {
		log.Fatal(err)
	}
	var songID int
	var name string

	for rows.Next() {
		err = rows.Scan(&songID, &name)
		if err != nil {
			log.Fatal(err)
		}
		songList = append(songList, Songs1{songID, name})

	}
	for i := 0; i < len(songList); i++ {
		fmt.Println("Song ID: ", songList[i].songid, "| Name: ", songList[i].name)
	}
	defer rows.Close()
	fmt.Println("Which song(ID) would you like to delete: ")
	var request int
	fmt.Scanln(&request)

	query = (`DELETE FROM playlist_songs WHERE playlistid = $1 AND songid = $2`)
	db.Query(query, destination, request)
	fmt.Println("Task completed")

}

func deletePlaylist(db *sql.DB) {
	playlistList := []Playlist{}
	rows, err := db.Query(`SELECT * FROM playlist`)
	if err != nil {
		log.Fatal(err)
	}
	var playlistID int
	var name string

	for rows.Next() {
		err := rows.Scan(&playlistID, &name)
		if err != nil {
			log.Fatal(err)
		}
		playlistList = append(playlistList, Playlist{playlistID, name})
	}
	for i := 0; i < len(playlistList); i++ {
		fmt.Println("Playlist ID: ", playlistList[i].playlistID, "| Name: ", playlistList[i].Name)
	}
	defer rows.Close()
	fmt.Println("Which playlist(Id) would you like to delete")
	var destination int
	fmt.Scanln(&destination)

	query := (`DELETE FROM playlist WHERE playlistid = $1`)
	db.Query(query, destination)
	query = (`DELETE FROM playlist_songs WHERE playlistid = $1`)
	db.Query(query, destination)
	fmt.Println("Task completed")

}
