CREATE TABLE Community (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4() ,
    name VARCHAR(255) NOT NULL,
    imageUrl VARCHAR(255),
    imageId VARCHAR(255),
    description TEXT,
    owner VARCHAR(255),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);