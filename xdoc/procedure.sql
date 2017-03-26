DROP PROCEDURE IF EXISTS `testspace`.`move_node`; 
CREATE DEFINER = `root`@`localhost` PROCEDURE 
    `testspace`.`move_node`(IN node_table varchar(100), IN node_id BIGINT, IN parent_id BIGINT)

BEGIN
DECLARE sql_str varchar(5000);

/*获取老的原路径*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT node.path into @old_path FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

set @old_path = concat(@old_path, node_id, '/');

/*获取新的父路径*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT node.path into @node_path FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', parent_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

set @node_path = concat(@node_path, parent_id, '/');
set @child_path = concat(@node_path, node_id, '/');

/*更新自己*/
set sql_str = '';
set sql_str = concat(sql_str, '  UPDATE ', node_table , ' SET parent_id = ' , parent_id, ',');
set sql_str = concat(sql_str, '             path = ', '''' , @node_path, '''');
set sql_str = concat(sql_str, '  WHERE id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;

/*更新后代*/
set sql_str = '';
set sql_str = concat(sql_str, '  UPDATE ', node_table);
set sql_str = concat(sql_str, '   SET path = REPLACE(path, ', '''', @old_path , '''', ',', '''' , @child_path, '''', ')');
set sql_str = concat(sql_str, '  WHERE path LIKE ', '''', @old_path, '%''');

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;

/*查询*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT * FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

END
======
DROP PROCEDURE IF EXISTS `testspace`.`update_node`; 
CREATE DEFINER = `root`@`localhost` PROCEDURE 
    `testspace`.`update_node`(IN node_table varchar(100), IN node_id BIGINT, IN status_name varchar(100), IN status_value varchar(100))

BEGIN
DECLARE sql_str varchar(5000);

/*获取路径*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT node.path into @node_path FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

set @node_path = concat(@node_path, node_id, '/');

/* 更新 */
set sql_str = '';
set sql_str = concat(sql_str, '  update ', node_table);
set sql_str = concat(sql_str, '    SET ', status_name, ' = ' , status_value);
set sql_str = concat(sql_str, '  WHERE id =', node_id, ' OR path LIKE ', '''', @node_path, '%''');

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;  
EXECUTE stmt;

/*查询*/
set sql_str = '';
set sql_str = concat(sql_str, '  SELECT * FROM  ', node_table, ' node');
set sql_str = concat(sql_str, '     WHERE node.id = ', node_id);

set @sql_str = sql_str;
PREPARE stmt FROM @sql_str;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

END;