package funcs

import (
	"database/sql"
	"math"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// самые залайканные аудио по годам
func DoHistogrammAudio_Years(db *sql.DB) error {
	yearsLikes, err := db.Query(
		"SELECT audio.\"release data\" as \"year\", SUM(metric.likes) as likes FROM audio " + 
		"JOIN metric " +
		"ON audio.\"metric id\" = audio.id " +
		"GROUP BY audio.\"release data\" " + 
		"ORDER BY audio.\"release data\"",
	)
	if err != nil {
		return fmt.Errorf("cant get year-likes")
	}
	defer yearsLikes.Close()

	years := []int{}
	likes := []int{}

	for yearsLikes.Next() {
		newYear, newLike := 0, 0
		err := yearsLikes.Scan(&newYear, &newLike)
		if err != nil {
			return fmt.Errorf("cant parse years + likes")
		}

		years = append(years, newYear)
		likes = append(likes, newLike)
	}

	p := plot.New()
	p.X.Tick.Label.Rotation = math.Pi / 4
	p.X.Tick.Label.Font.Size = 5
	p.Title.Text = "Гистограмма распределения лайков по годам релизам музыки"
	p.Y.Label.Text = "Количество лайков у музыки в этом году"

	v := make(plotter.Values, len(likes))
	for i := range v {
		v[i] = float64(likes[i])
	}
	
	w := vg.Points(5)

	bar, err := plotter.NewBarChart(v, w)
	if err != nil {
		panic(err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(0)

	p.Add(bar)
	yearsSrt := []string{}
	for i := range years {
		yearsSrt = append(yearsSrt, fmt.Sprint(years[i]))
	}
	p.NominalX(yearsSrt...)

	if err := p.Save(10*vg.Inch, 6*vg.Inch, "audio_years.png"); err != nil {
		return fmt.Errorf("cant do bar: %w", err)
	}

	return nil
}