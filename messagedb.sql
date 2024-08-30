CREATE DATABASE IF NOT EXISTS messagedb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE messagedb;

-- Таблица Roles
CREATE TABLE Roles (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Вставка стандартных ролей
INSERT INTO Roles (Name) VALUES ('Администратор'), ('Пользователь'),('Техническая поддержка'),('Модератор'),('Менеджер');

-- Таблица Users
CREATE TABLE Users (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Username VARCHAR(255) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL UNIQUE,
    RoleID INT NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    FOREIGN KEY (RoleID) REFERENCES Roles(ID),
    INDEX idx_email (Email),
    INDEX idx_username (Username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица Chats (Чаты используются как комнаты поддержки)
CREATE TABLE Chats (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    AssignedTo INT,  -- Кто ответственный за чат (например, сотрудник техподдержки)
    EntryPoint VARCHAR(255), 
    Status VARCHAR(100) DEFAULT 'active',  -- Статус чата (active, pending, closed)
    FOREIGN KEY (AssignedTo) REFERENCES Users(ID),
    INDEX idx_assigned_to (AssignedTo),
    INDEX idx_status (Status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица Attachments
CREATE TABLE Attachments (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    FileName VARCHAR(255) NOT NULL,
    FileType VARCHAR(100) NOT NULL,
    FileData LONGBLOB NOT NULL,
    FileSize BIGINT NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    INDEX idx_file_name (FileName)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица Messages
CREATE TABLE Messages (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    ChatID INT NOT NULL,
    UserID INT NOT NULL,
    Content TEXT,
    CreatedAt DATETIME NOT NULL,
    IsRead BOOLEAN DEFAULT FALSE,
    IsChecked BOOLEAN DEFAULT FALSE,
    AttachedID INT DEFAULT NULL,  -- Поле для прикрепленных файлов
    FOREIGN KEY (ChatID) REFERENCES Chats(ID),
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    FOREIGN KEY (AttachedID) REFERENCES Attachments(ID),
    INDEX idx_chat_id (ChatID),
    INDEX idx_user_id (UserID),
    INDEX idx_attached_id (AttachedID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица ChatUsers
CREATE TABLE ChatUsers (
    ChatID INT,
    UserID INT,
    JoinedAt DATETIME NOT NULL,
    PRIMARY KEY (ChatID, UserID),
    FOREIGN KEY (ChatID) REFERENCES Chats(ID),
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    INDEX idx_chat_user (ChatID, UserID)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица Shifts
CREATE TABLE Shifts (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    StartTime DATETIME NOT NULL,
    EndTime DATETIME DEFAULT NULL,  -- Поле для окончания смены
    Active BOOLEAN DEFAULT TRUE,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    INDEX idx_user_id (UserID),
    INDEX idx_active (Active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица VirtualEmails
CREATE TABLE VirtualEmails (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    RealEmail VARCHAR(255) NOT NULL UNIQUE,
    VirtualEmail VARCHAR(255) NOT NULL UNIQUE,
    ChatID INT NOT NULL,
    CreatedAt DATETIME NOT NULL,
    FOREIGN KEY (ChatID) REFERENCES Chats(ID),
    INDEX idx_chat_id (ChatID),
    INDEX idx_real_email (RealEmail),
    INDEX idx_virtual_email (VirtualEmail)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица DelayedMessages
CREATE TABLE DelayedMessages (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    ChatID INT NOT NULL,
    Message TEXT,
    SendTime DATETIME NOT NULL,
    IsSent BOOLEAN DEFAULT FALSE,
    CreatedAt DATETIME NOT NULL,
    FOREIGN KEY (ChatID) REFERENCES Chats(ID),
    INDEX idx_chat_id (ChatID),
    INDEX idx_send_time (SendTime)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица SupportRatings
CREATE TABLE SupportRatings (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    ChatID INT NOT NULL,
    UserID INT NOT NULL,
    Rating INT NOT NULL CHECK (Rating BETWEEN 1 AND 5),  -- Ограничение рейтинга от 1 до 5
    Comment TEXT,
    CreatedAt DATETIME NOT NULL,
    FOREIGN KEY (ChatID) REFERENCES Chats(ID),
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    INDEX idx_chat_id (ChatID),
    INDEX idx_user_id (UserID),
    INDEX idx_rating (Rating)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица AccessTokens
CREATE TABLE AccessTokens (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    Token VARCHAR(512) NOT NULL,
    ExpiresAt DATETIME NOT NULL,
    CreatedAt DATETIME NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    INDEX idx_user_id (UserID),
    INDEX idx_token (Token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица Sessions
CREATE TABLE Sessions (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    IPAddress VARCHAR(45) NOT NULL,
    UserAgent VARCHAR(255),
    LastActive DATETIME NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL,
    DeletedAt DATETIME DEFAULT NULL,  -- Поле для мягкого удаления
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    INDEX idx_user_id (UserID),
    INDEX idx_ip_address (IPAddress)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Таблица RefreshTokens
CREATE TABLE RefreshTokens (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    Token VARCHAR(512) NOT NULL,
    ExpiresAt DATETIME NOT NULL,
    CreatedAt DATETIME NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(ID),
    INDEX idx_user_id (UserID),
    INDEX idx_token (Token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `LanguageSkills` (
    `ID` int NOT NULL AUTO_INCREMENT,
    `UserID` int NOT NULL,
    `Language` varchar(255) NOT NULL,
    `Level` varchar(255) NOT NULL,
    `CreatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `UpdatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`ID`),
    FOREIGN KEY (`UserID`) REFERENCES `Users`(`ID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;