
INSERT INTO users (name, nickname, email, password)
VALUES 
("User 1", "user_1", "user1@gmail.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$MPr2V9.4oX4fBpgrjzFC5.KBDuGTi6H680Zc74Z6IXBfzKE5ylQwC");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(2, 3);

INSERT INTO posts(title, content, author_id)
values
("User post number 1", "User post content number 1 for user with id 1", "1"),
("User post number 2", "User post content number 2 for user with id 1", "1"),
("User post number 3", "User post content number 3 for user with id 2", "2");
