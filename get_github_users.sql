SELECT u.*, p.id AS "github id" FROM users as u, users_github as p
JOIN users_github ON u.id = p.user_id