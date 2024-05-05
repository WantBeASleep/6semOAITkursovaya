package truncate

import (
	"kra/scripts/model"
)

// не работает
func GetFuncs() []model.FuncInfo {
	return []model.FuncInfo{
		{
			Name: "fill audio",
			Efunc: truncateAllTables,
		},
	}
}