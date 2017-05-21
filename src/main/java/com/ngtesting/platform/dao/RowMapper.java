package com.ngtesting.platform.dao;

import java.sql.ResultSet;
import java.sql.SQLException;

public interface RowMapper {

    Object mapRow(ResultSet rs, int index) throws SQLException;
}
