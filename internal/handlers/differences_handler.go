package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/microsoft/go-mssqldb"

	"github.com/gonzalo-wi/DiferenciasNautom/internal/db"
	"github.com/gonzalo-wi/DiferenciasNautom/internal/models"
)

func GetDifferences(c *gin.Context) {
	desde := c.Query("desde")
	hasta := c.Query("hasta")
	if desde == "" || hasta == "" {
		c.JSON(400, gin.H{"error": "Missing 'desde' or 'hasta' query parameters"})
		return
	}

	database, err := db.NewSqlServerDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed: " + err.Error()})
		return
	}
	defer database.Close()

	query := `
		SELECT
			CONVERT(VARCHAR(10), fecha, 23) AS fecha,
			user_name,
			diferencia_vs_aguas,
			esperado_nuestro,
			esperado_aguas
		FROM dbo.vw_comparacion_reparto_dia_aguas
		WHERE fecha BETWEEN @desde AND @hasta
		  AND diferencia_vs_aguas <> 0
		ORDER BY fecha, user_name
	`

	rows, err := database.Query(query,
		sql.Named("desde", desde),
		sql.Named("hasta", hasta),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.Difference
	var stats models.Statistics

	for rows.Next() {
		var fecha, userName string
		var diferencia, esperadoNuestro, esperadoAguas sql.NullFloat64

		err := rows.Scan(&fecha, &userName, &diferencia, &esperadoNuestro, &esperadoAguas)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		diferenciaVal := 0.0
		if diferencia.Valid {
			diferenciaVal = diferencia.Float64
		}

		tipo := "sobrante"
		if diferenciaVal < 0 {
			tipo = "faltante"
			stats.TotalFaltantes++
		} else {
			stats.TotalSobrantes++
		}

		stats.TotalDiferencias++
		stats.Consolidados++

		esperadoNuestroVal := 0.0
		if esperadoNuestro.Valid {
			esperadoNuestroVal = esperadoNuestro.Float64
		}

		esperadoAguasVal := 0.0
		if esperadoAguas.Valid {
			esperadoAguasVal = esperadoAguas.Float64
		}

		d := models.Difference{
			Date:            fecha,
			Reparto:         userName,
			Diferencia:      diferenciaVal,
			Tipo:            tipo,
			UserName:        userName,
			DepositEsperado: esperadoAguasVal,
			TotalAmount:     esperadoNuestroVal,
		}
		items = append(items, d)
	}

	response := models.DifferenceResponse{
		Status:       "ok",
		Desde:        desde,
		Hasta:        hasta,
		Estadisticas: stats,
		Items:        items,
	}

	c.JSON(http.StatusOK, response)
}
