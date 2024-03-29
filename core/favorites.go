package core

import (
	"encoding/json"
	"os"
)

var Favorites []int

func AddFavorite(artistID int) {
	Favorites = append(Favorites, artistID)
}

func RemoveFavorite(artist int) {
	index := -1
	for i, fav := range Favorites {
		if fav == artist {
			index = i
			break
		}
	}
	if index != -1 {
		Favorites = append(Favorites[:index], Favorites[index+1:]...)
	}
}

func SaveFavorites() error {
	jsonBytes, err := json.Marshal(Favorites)
	if err != nil {
		return nil
	}
	file, err := os.Create(FavFile)
	if err != nil {
		return nil
	}
	defer file.Close()
	_, err = file.Write(jsonBytes)
	if err != nil {
		return nil
	}
	return nil
}
