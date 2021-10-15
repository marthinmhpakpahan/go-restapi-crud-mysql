package main

func create(videoGame VideoGame) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", videoGame.Name, videoGame.Genre, videoGame.Year)
	return err
}

func delete(id int64) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM video_games WHERE id = ?", id)
	return err
}

func update(videoGame VideoGame) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?", videoGame.Name, videoGame.Genre, videoGame.Year, videoGame.Id)
	return err
}

func all() ([]VideoGame, error) {
	videoGames := []VideoGame{}
	db, err := getDB()
	if err != nil {
		return videoGames, err
	}

	rows, err := db.Query("SELECT id, name, genre, year FROM video_games")
	if err != nil {
		return videoGames, err
	}

	for rows.Next() {
		var videoGame VideoGame
		err = rows.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genre, &videoGame.Year)
		if err != nil {
			return videoGames, err
		}
		videoGames = append(videoGames, videoGame)
	}

	return videoGames, nil
}

func get(id int64) (VideoGame, error) {
	var videoGame VideoGame
	db, err := getDB()
	if err != nil {
		return videoGame, err
	}

	row := db.QueryRow("SELECT id, name, genre, year FROM video_games WHERE id = ?", id)
	err = row.Scan(videoGame.Id, videoGame.Name, videoGame.Genre, videoGame.Year)
	if err != nil {
		return videoGame, err
	}

	return videoGame, nil
}