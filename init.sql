CREATE TABLE Users (
	UserId SERIAL PRIMARY KEY,
	Username TEXT
);

CREATE TABLE Posts (
    PostId SERIAL PRIMARY KEY,
    UserId INT NOT NULL,
    Title TEXT,
    CONSTRAINT fk_User FOREIGN KEY(UserId) REFERENCES Users(UserId) ON DELETE CASCADE
);

CREATE TABLE PostLikes (
    UserId INT NOT NULL,
    PostId INT NOT NULL,
    PRIMARY KEY (UserId, PostId),
    CONSTRAINT fk_User FOREIGN KEY(UserId) REFERENCES Users(UserId) ON DELETE CASCADE,
    CONSTRAINT fk_Post FOREIGN KEY(PostId) REFERENCES Posts(PostId) ON DELETE CASCADE
);

INSERT INTO 
    Users (Username) 
VALUES 
    ('Max'),
    ('Vitaly'),
    ('Andrei'),
    ('Danil');

INSERT INTO
    Posts (UserId, Title)
VALUES
    (1, 'first'),
    (1, 'Second'),
    (2, 'Third'),
    (3, 'new Post'),
    (3, 'Second Post'),
    (3, 'My world'),
    (4, 'sample text');

INSERT INTO
    PostLikes (UserId, PostId)
VALUES
    (1, 3),
    (2, 4),
    (3, 7),
    (4, 5),
    (3, 3),
    (4, 3),
    (2, 7);