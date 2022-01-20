package api

import "inventory-app/ent"

type API struct {
	Db *ent.Client
}
