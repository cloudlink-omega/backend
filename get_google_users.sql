SELECT u.*, p.id AS "google id" FROM users as u, users_google as p
JOIN users_google ON u.id = p.user_id