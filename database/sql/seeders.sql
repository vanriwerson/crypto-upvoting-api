TRUNCATE cryptoUpvoting.users;
INSERT INTO cryptoUpvoting.users (name, email, password)
VALUES ("Adalberto", "adalberto@email.com", "$2a$10$mutnbTgnDCjCxPuU.q1z3.VlEoCtBERyhvw9p.8ML2d49LoaQkWmO"), # password: senha1
("Giovane", "giovane@email.com", "$2a$10$lYiAF6rla0e6ZeK2Arw34.acjjd28g0A6/2/IDPg5VenpUeoW4Tn2"), # password: senha2
("Jessy", "jessy@email.com", "$2a$10$9cI08MwoPmcAFkUbCJ54Tusd7uUlHj2Kl6.5MKwV1vhe2NYMeBh.e"), # password: senha3
("Rafael", "rafael@email.com", "$2a$10$AGRhss/Nv9Ze0KAwoqUHou8RMVqtoX1qDz/JLbuuTX3WxtWL4oMyG"), # password: senha4
("Maristela", "maristela@email.com", "$2a$10$wpywTh8fB59gwVNqM5IU0OrRNm4IJmJfHIP9kcLrhBoNhepxYFd.q"), # password: senha5
("Ana", "ana@email.com", "$2a$10$WWfAj/0gO5hx2BrHJmvEMOTnuXxS3Ecjcm/XxJOsdagT3dFW82IuC"), # password: senha6
("Clair", "clair@email.com", "$2a$10$mZzpwjxuYSUNIRiuDb81.uvsuxMRr.VKdVN5lzm8xnVYCwwgR1tQq"), # password: senha7
("Admin", "admin@email.com", "$2a$10$kSLHBFOOH4Fiib0E1J1NJO7ABtNrpwT52kutIJy1ojSm1dd69cYyC"); # password: admin

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