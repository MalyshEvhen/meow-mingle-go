MATCH
    (u:User)
WHERE
    ID(u)=$author_id
CREATE
    (u)-[:WRITE {role: 'Author'}]->(p:Post {content: $content})
RETURN
    ID(p) AS id,
    p.content AS content,
    ID(u) AS author_id;