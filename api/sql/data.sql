-- Seed data for local development.
--
-- Every user below shares the same bcrypt hash, which is the password
-- "123456" - handy for testing since you can log in as any of them with
-- the same credentials (see collection/login/login.posting.yaml).
INSERT INTO users (name, nickname, email, password)
VALUES
("Ana Silva", "ana.silva", "ana.silva@example.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("Bruno Costa", "bruno.costa", "bruno.costa@example.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("Carla Souza", "carla.souza", "carla.souza@example.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("Diego Martins", "diego.martins", "diego.martins@example.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("Elisa Ferreira", "elisa.ferreira", "elisa.ferreira@example.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("Fernando Lima", "fernando.lima", "fernando.lima@example.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC");

-- ids, in insertion order: 1 Ana, 2 Bruno, 3 Carla, 4 Diego, 5 Elisa, 6 Fernando
INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2), -- Bruno follows Ana
(1, 3), -- Carla follows Ana
(2, 1), -- Ana follows Bruno
(2, 4), -- Diego follows Bruno
(3, 1), -- Ana follows Carla
(3, 5), -- Elisa follows Carla
(4, 6), -- Fernando follows Diego
(5, 2), -- Bruno follows Elisa
(6, 1); -- Ana follows Fernando

INSERT INTO posts (title, content, author_id)
VALUES
("Getting started with Go", "Today I set up my first Go module and it finally clicked - structs and interfaces are great.", 1),
("Weekend hike recap", "Spent the weekend hiking near the coast, the view from the top was totally worth it.", 1),
("Learning MySQL indexes", "Digging into fulltext indexes this week, they make search so much faster.", 2),
("Coffee brewing tips", "Switched to a pour-over setup and my mornings improved a lot.", 2),
("Refactoring the API layer", "Split the controllers from the repositories, code is much easier to test now.", 3),
("Book recommendation", "Just finished a great book on distributed systems, highly recommend it to backend folks.", 3),
("JWT authentication notes", "Wrote up some notes on how JWT claims work, useful for the next teammate.", 4),
("New keyboard day", "Got a new mechanical keyboard, typing feels so much better now.", 5),
("Docker compose for local dev", "Set up docker compose so the whole team can spin up MySQL in one command.", 6),
("First contribution merged", "Opened my first PR on the project and got it merged today!", 6);
