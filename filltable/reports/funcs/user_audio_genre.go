package funcs
import (
	"database/sql"
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func DoHistogramUser_Audio(db *sql.DB) error {
	p := plot.New()
	p.Title.Text = "Гистограмма распределения жанров по музыке, добавленной пользователями"
	p.Y.Label.Text = "Количество аудио данного жанра"

	w := vg.Points(70)

	genreFrq, err := db.Query(
		"SELECT genre.appellation, count(*) FROM user_audio " + 
		"JOIN audio_genre " + 
		"ON user_audio.\"audio id\" = audio_genre.\"audio id\" " + 
		"JOIN genre " + 
		"ON genre.id = audio_genre.\"genre id\" " + 
		"GROUP BY genre.appellation",
	)
	if err != nil {
		return fmt.Errorf("cant get user_audio stats: %w", err)
	}
	defer genreFrq.Close()

	type desc struct {
		genre string
		count int
	}

	data := []desc{}
	for genreFrq.Next() {
		newdesc := desc{}
		err := genreFrq.Scan(&newdesc.genre, &newdesc.count)
		if err != nil {
			return fmt.Errorf("cant parse genre-count: %w", err)
		}
		data = append(data, newdesc)
	}
	
	v := make(plotter.Values, len(data))
	for i := range v {
		v[i] = float64(data[i].count)
	}

	bar, err := plotter.NewBarChart(v, w)
	if err != nil {
		panic(err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(0)

	p.Add(bar)
	names := []string{}
	for i := range data {
		names = append(names, data[i].genre)
	}
	p.NominalX(names...)

	if err := p.Save(10*vg.Inch, 6*vg.Inch, "user_genre_chart.png"); err != nil {
		return fmt.Errorf("cant do bar: %w", err)
	}
	return nil
}

