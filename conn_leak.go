rows, _ := db.Query("SELECT id, name FROM users WHERE id in (?)", ids)
defer rows.Close()

for rows.Next() {
    var id int
    if err := rows.Scan(&id); err != nil {
       return err
    }
    rows2, _ := db.Query("SELECT order_id, total  FROM order WHERE user_id = ?", id)
    defer rows2.Close()
    doSomeProcessing()
}
