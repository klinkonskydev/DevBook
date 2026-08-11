INSERT INTO users (name, nickname, email, password)
VALUES 
("User 1", "user_1", "user1@gmail.com", "$2a$10$jTu1.UMtW6cAQeBjl9WXk.Yys.PxYS.8/CPuLY8GrbdJVbd8Vv6bO"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$jTu1.UMtW6cAQeBjl9WXk.Yys.PxYS.8/CPuLY8GrbdJVbd8Vv6bO"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$jTu1.UMtW6cAQeBjl9WXk.Yys.PxYS.8/CPuLY8GrbdJVbd8Vv6bO"),
("User 4", "user_4", "user4@gmail.com", "$2a$10$jTu1.UMtW6cAQeBjl9WXk.Yys.PxYS.8/CPuLY8GrbdJVbd8Vv6bO"),
("User 5", "user_5", "user5@gmail.com", "$2a$10$jTu1.UMtW6cAQeBjl9WXk.Yys.PxYS.8/CPuLY8GrbdJVbd8Vv6bO");

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(2, 3);
