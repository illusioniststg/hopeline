package routeplanning

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func AssignBoat(w http.ResponseWriter, r *http.Request) {
	// Parse request to get survivorID and boatID
	survivor_id := r.URL.Query().Get("survivor_id")
	boat_id := r.URL.Query().Get("boat_id")
	fmt.Println("Request data", survivor_id, boat_id)
	/*
	// Call the function to assign the boat
	if err := AssignBoatToSurvivor(r.Context(), db, req.SurvivorID, req.BoatID); err != nil {
		http.Error(w, fmt.Sprintf("Failed to assign boat: %v", err), http.StatusInternalServerError)
		return
	}
	*/
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hurray!!! Boat assigned successfully")
}

func AssignBoatToSurvivor(ctx context.Context, db *sql.DB, survivorID int64, boatID string) error {
	query := "UPDATE survivors SET boat_id = ? WHERE id = ?"
	result, err := db.ExecContext(ctx, query, boatID, survivorID)
	if err != nil {
		return fmt.Errorf("failed to assign boat to survivor: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no survivor found with ID %d", survivorID)
	}

	log.Printf("Successfully assigned boat %s to survivor %d", boatID, survivorID)
	return nil
}	