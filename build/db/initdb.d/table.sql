-- テーブルが既にある場合は一旦削除
DROP TABLE IF EXISTS users;
-- テーブルの作成
CREATE TABLE users (
  id int primary key,
  username varchar(50),
  emailaddress varchar(50),
  password varchar(50),
  isadmin boolean,
  canlogin boolean,
  created_at datetime,
  updated_at datetime
);
-- データの挿入
-- INSERT INTO
--   users (
--     id,
--     username,
--     emailaddress,
--     password,
--     isadmin,
--     canlogin,
--     created_at,
--     updated_at
--   )
-- VALUES
--   (
--     1,
--     "makoto",
--     "hoge@email.com",
--     "password",
--     true,
--     true,
--     sysdate(),
--     sysdate()
--   );
-- -- テーブルの表示
-- SELECT
--   *
-- FROM
--   users;