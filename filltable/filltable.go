package filltable

import (
	"fmt"

	"slices"

	"database/sql"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/lib/pq"

	"kra/data"

	"kra/filltable/model"

	"kra/cli"
	"kra/filltable/nonrelation"
	"kra/filltable/relation"
	"kra/filltable/reports"
)

func Manager(db *sql.DB) error {
	data, err := data.GetOrUpdateData(db)
	if err != nil {
		return fmt.Errorf("cant read csv: %w", err)
	}

	nonRelationFuncs := nonrelation.GetFuncs(db)
	relationFuncs := relation.GetFuncs(db)
	reports := reports.GetFuncs(db)

	handler := model.InitFuncHolder(
		db,
		data,
		slices.Concat(nonRelationFuncs, relationFuncs, reports),
	)

	term := tea.NewProgram(cli.InitialModel(handler.GetOptionDescription()))
	m, err := term.Run()
	if err != nil {
		return fmt.Errorf("error while run terminal: %w", err)
	}

	cmdIds := m.(cli.Model).Selected

	for id := range cmdIds {
		if !cmdIds[id] {
			continue
		}
		if err := handler.DoFunc(id); err != nil {
			return err
		}
	}

	return nil
}
