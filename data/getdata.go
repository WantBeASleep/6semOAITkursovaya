package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"kra/constants"
	"math/rand"
	"os"
	"path"
	"strconv"

	"database/sql"

	_ "github.com/lib/pq"
)

func GetOrUpdateData(db *sql.DB) ([]TrackInfo, error) {
	curDir, _ := os.Getwd()
	data, err := readCSVFile(path.Join(curDir, constants.DATASET_PATH))
	if err != nil {
		return nil, err
	}

	isAudioNull := false
	err = db.QueryRow(
		"SELECT NOT EXISTS (SELECT 1 FROM audio)",
	).Scan(&isAudioNull)
	if err != nil {
		return nil, fmt.Errorf("cant check audio table exists: %w", err)
	}

	if !isAudioNull {
		for _, track := range data {
			csvFile, _ := os.Open(path.Join(curDir, constants.CostylAudio))
			defer csvFile.Close()

			r := csv.NewReader(csvFile)
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				track.Audio.Id, _ = strconv.Atoi(record[0])
			}
		}
	}

	isGenreNull := false
	err = db.QueryRow(
		"SELECT NOT EXISTS (SELECT 1 FROM genre)",
	).Scan(&isGenreNull)
	if err != nil {
		return nil, fmt.Errorf("cant check genre table exists: %w", err)
	}

	if !isGenreNull {
		for _, track := range data {
			csvFile, _ := os.Open(path.Join(curDir, constants.CostylGenre))
			defer csvFile.Close()

			r := csv.NewReader(csvFile)
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				track.Genre.Id, _ = strconv.Atoi(record[0])
			}
		}
	}

	isAuthorNull := false
	err = db.QueryRow(
		"SELECT NOT EXISTS (SELECT 1 FROM author)",
	).Scan(&isAuthorNull)
	if err != nil {
		return nil, fmt.Errorf("cant check author table exists: %w", err)
	}

	if !isAuthorNull {
		for _, track := range data {
			csvFile, _ := os.Open(path.Join(curDir, constants.CostylAuthor))
			defer csvFile.Close()

			r := csv.NewReader(csvFile)
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				track.Author.Id, _ = strconv.Atoi(record[0])
			}
		}
	}

	return data, nil
}

func readCSVFile(path string) ([]TrackInfo, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cant open csv file: %w", err)
	}
	defer csvFile.Close()

	parsedTracks := make([]TrackInfo, 0, 30000)

	r := csv.NewReader(csvFile)
	for {
		var line TrackInfo
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("cant read csv file: %w", err)
		}

		line.Author.Appellation = record[1]

		line.Audio.Appellation = record[2]
		line.Audio.Release = record[3]
		line.Audio.Lyric = record[5]
		
		trackLen, err := strconv.Atoi(record[6])
		if err != nil {
			return nil, fmt.Errorf("cant convert track len: %w", err)
		}
		
		start, end := rand.Intn(trackLen), rand.Intn(trackLen)
		if start > end {
			start, end = end, start
		}
		line.Snippet.Start = start
		line.Snippet.End = end

		line.Genre.Appellation = record[4]
		
		parsedTracks = append(parsedTracks, line)
	}

	return parsedTracks, nil
}