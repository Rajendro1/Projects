CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(100),
  email VARCHAR(100),
  age INT
);

INSERT INTO users (name, email, age)
SELECT CONCAT('User', n), CONCAT('user', n, '@mail.com'), 20 + (n % 30)
FROM (SELECT @row := @row + 1 AS n FROM (SELECT 0 UNION ALL SELECT 1) t1,
      (SELECT 0 UNION ALL SELECT 1) t2,
      (SELECT 0 UNION ALL SELECT 1) t3,
      (SELECT 0 UNION ALL SELECT 1) t4,
      (SELECT @row := 0) r LIMIT 1000) nums;
