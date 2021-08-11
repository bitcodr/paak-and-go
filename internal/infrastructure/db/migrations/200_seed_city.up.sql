-- why we ship the cities from txt file to the db:
-- reasons:

-- 1 - we can have more control on report we wanna get in the future and the relation between data

-- 2 - we can add other parameters for each city, maybe in the future per city the calculation of price is different
-- and each city will have its own ratio

-- 2 - we can cache the txt file in the memory that's right it is fast - but we have also do it from db with many tool exist
-- beside we have more control on data

INSERT INTO cities
    (name)
VALUES ('Barcelona'),
       ('Seville'),
       ('Madrid'),
       ('Valencia'),
       ('Andorra la Vella'),
       ('Malaga');

-- another way is import the txt file directly to db
-- COPY cities (name) FROM '../seed/cities.txt';
