tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
if err != nil {
        log.Fatal(err)
}
id := 37

// transactional update
_, execErr := tx.Exec(`UPDATE users SET status = ? WHERE id = ?`, "paid", id)
if execErr != nil {
        _ = tx.Rollback()
        log.Fatal(execErr)
}

// non transactional update
_, execErr := db.Exec(`UPDATE payment SET paid = ? WHERE id = ?`, "yes", 47)

if err := tx.Commit(); err != nil {
        log.Fatal(err)
}
