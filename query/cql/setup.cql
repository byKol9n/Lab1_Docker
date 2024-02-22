LOAD CSV WITH HEADERS FROM 'file:///data/people.csv' AS line 
FIELDTERMINATOR ',' 
MERGE (p:Person {person_id:line.personId, name:line.name});

USING PERIODIC COMMIT 
LOAD CSV WITH HEADERS FROM "file:///data/friendships.csv" AS line 
FIELDTERMINATOR ',' 
MATCH (p1:Person),(p2:Person) 
WHERE p1.person_id = line.personId1 AND p2.person_id= line.personId2 
CREATE UNIQUE (p1)-[r:KNOWS]->(p2);