package cn.mobiu.events.dao;

import java.sql.ResultSet;
import java.sql.SQLException;

/**
 * <简述> base Dao 所需要的RowMapper 接口
 * <详细描述>
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
public interface RowMapper {
    /**
     * 〈简述〉
     * 〈详细描述〉
     *
     * @param rs    结果集
     * @param index 索引
     * @return 返回
     * @throws SQLException 异常
     * @author xuxiang
     */
    Object mapRow(ResultSet rs, int index) throws SQLException;
}
