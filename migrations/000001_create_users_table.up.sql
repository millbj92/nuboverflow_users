CREATE TABLE IF NOT EXISTS Users (
    ID serial NOT NULL PRIMARY KEY,
    UserName varchar(50),
    Email varchar(50),
    Github varchar(50),
    Linkedin varchar(50),
    UserScore INT,
    Bio varchar(50),
    Profession varchar(50),
    WorkPlace varchar(50),
    Interests varchar(500)
);

CREATE TABLE IF NOT EXISTS Awards (
    ID serial NOT NULL PRIMARY KEY,
    AwardName varchar(50) NOT NULL,
    AwardDescription varchar(150) NOT NULL,
    AwardLevel int NOT NULL,
    AwardScore int NOT NULL
);

CREATE TABLE IF NOT EXISTS UserAwards (
    ID serial NOT NULL PRIMARY KEY,
    AwardID serial NOT NULL,
    UserID serial NOT NULL,
    CONSTRAINT fk_userawards_award_id
        FOREIGN KEY(AwardID)
            REFERENCES Awards(ID),
    CONSTRAINT fk_userawards_user_id
        FOREIGN KEY(UserID)
            REFERENCES Users(ID)
);

