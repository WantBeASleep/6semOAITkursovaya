package relation

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"

	"kra/data"
)

func Manager(data []data.TrackInfo, db *sql.DB) error {
	fmt.Println("M🤙:N🤙")
	
	// M:N
	if err := fillAudio_Genre(data, db); err != nil {
		return fmt.Errorf("cant fill audio genre: %w", err)
	}

	if err := fillAudio_User(db); err != nil {
		return fmt.Errorf("cant fill audio user: %w", err)
	}

	if err := fillAuthor_Audio(data, db); err != nil {
		return fmt.Errorf("cant fill author audio: %w", err)
	}

	if err := fillAlbum_Audio(db); err != nil {
		return fmt.Errorf("cant fill album audio: %w", err)
	}

	if err := fillAlbum_Genre(db); err != nil {
		return fmt.Errorf("cant fill album genre: %w", err)
	}

	if err := fillAlbum_User(db); err != nil {
		return fmt.Errorf("cant fill album user: %w", err)
	}

	if err := fillAlbum_Author(db); err != nil {
		return fmt.Errorf("cant fill album author: %w", err)
	}

	if err := relationUser_Author(data, db); err != nil {
		return fmt.Errorf("cant fill user author roles: %w", err)
	}

	return nil
}