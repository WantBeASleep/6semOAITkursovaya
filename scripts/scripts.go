package scripts

import (
	"fmt"

	"slices"

	"database/sql"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"
	"golang.org/x/exp/maps"

	"kra/data"

	"kra/scripts/model"
	"kra/scripts/nonrelation"
	"kra/scripts/relation"
	"kra/scripts/truncate"
	"kra/terminal"
)

func Manager(db *sql.DB) error {
	data, err := data.GetOrUpdateData(db)
	if err != nil {
		return fmt.Errorf("cant read csv: %w", err)
	}

	nonRelationFuncs := nonrelation.GetFuncs()
	relationFuncs := relation.GetFuncs()
	truncateFuncs := truncate.GetFuncs()

	handler := model.InitFuncHolder(
		db,
		data,
		slices.Concat(nonRelationFuncs, relationFuncs, truncateFuncs),
	)

	term := tea.NewProgram(terminal.InitialModel(handler.GetFuncsNames()))
	m, err := term.Run();
	if err != nil {
		return fmt.Errorf("error while run terminal: %w", err)
	}

	cmdIds := maps.Keys(m.(terminal.Model).Selected)
	slices.Sort(cmdIds)

	for _, id := range cmdIds {
		if err := handler.DoFunc(id); err != nil {
			return err
		}
	}

	return nil
}