-- When manual modifications must be made to in-use dbs, 
-- update the schema.sql file for testing and add them below. 
-- Then apply the new (via git) updates one at a time, as needed:

CREATE UNIQUE INDEX Users_AccountEmail_IDX ON Users (AccountEmail);
