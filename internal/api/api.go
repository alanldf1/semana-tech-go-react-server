package api

import "github.com/alanldf1/semana-tech-go-react-server/internal/store/pgstore"

type apiHandler struct {
	q *pgstore.Queries
}
