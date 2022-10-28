TRUNCATE cryptoUpvoting.users;
INSERT INTO cryptoUpvoting.users (name, email, password)
VALUES ("Adalberto", "adalberto@email.com", "superSenhaSecreta1"),
("Giovane", "giovane@email.com", "superSenhaSecreta2"),
("Jessy", "damasceno@email.com", "superSenhaSecreta3"),
("Rafael", "rafael@email.com", "superSenhaSecreta4"),
("Maristela", "maristela@email.com", "superSenhaSecreta5"),
("Ana", "anahelena@email.com", "superSenhaSecreta6"),
("Clair", "clair@email.com", "superSenhaSecreta7");

TRUNCATE cryptoUpvoting.cryptos;

INSERT INTO cryptoUpvoting.cryptos (name)
VALUES ("Bitcoin"), ("Ethereum"), ("Litecoin"),
("Tether"), ("BNB"), ("USD Coin"),
("XRP"), ("Binance USD"), ("Cardano"),
("Solana"), ("Dogecoin"), ("Polygon"),
("Polkadot"), ("Dai"), ("Shiba Inu");

TRUNCATE cryptoUpvoting.votes;
INSERT INTO cryptoUpvoting.votes (userId, cryptoId)
VALUES ("1", "1"), ("1", "2"), ("2", "1"),
("2", "3"), ("2", "4"), ("3", "1"), ("3", "2"),
("4", "1"), ("4", "4"), ("4", "5"), ("5", "6"),
("5", "1"), ("5", "7"), ("6", "8"), ("6", "9"),
("6", "10"), ("7", "14"), ("7", "15");