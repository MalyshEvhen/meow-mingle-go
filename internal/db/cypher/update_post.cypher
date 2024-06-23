MATCH
    (u:User {id: $author_id})-[:WRITE]->(p:Post {id: $id})
OPTIONAL MATCH
    (p)-[l:LIKE]-()
WITH
    u,p,count(l) AS likes
SET
    p.content=$content
RETURN
    p.id AS id,
    p.content AS content,
    u.id AS author_id,
    likes;
