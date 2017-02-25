package cn.linkr.testspace.service.impl;

import java.io.Serializable;
import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.ResultSetMetaData;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.PreparedStatementCreator;
import org.springframework.jdbc.core.RowMapper;

import cn.linkr.testspace.dao.IBaseDao;
import cn.linkr.testspace.entity.BaseEntity;
import cn.linkr.testspace.service.BaseService;
import cn.linkr.testspace.util.DaoHelper;
import cn.linkr.testspace.util.SpringContextHolder;
import cn.linkr.testspace.vo.Page;

/**
 * service实现类的基类
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
public class BaseServiceImpl implements BaseService {

    /**
     * 基本dao
     */
    @Autowired
    private IBaseDao dao;

    /**
     * jdbcTemplate
     */
    @Autowired
    private JdbcTemplate jdbcTemplate;

    /**
     * 获得dao
     *
     * @return 返回dao
     */
    protected IBaseDao getDao() {
        if (dao == null) {
            dao = SpringContextHolder.getBean("baseDao", IBaseDao.class);
        }
        return dao;
    }

    /**
     * 获得JdbcTemplate
     *
     * @return 返回JdbcTemplate
     */
    protected JdbcTemplate getJdbcTemplate() {
        if (jdbcTemplate == null) {
            jdbcTemplate = SpringContextHolder.getBean("jdbcTemplate", JdbcTemplate.class);
        }
        return jdbcTemplate;
    }

    /**
     * 判断编码是否存在，确保编码的唯一性
     *
     * @param clazz     <font color='red'>*必要参数</font>
     * @param codeName  编码的字段名,如果为null，则默认为code
     * @param codeValue 编码的字段值 <font color='red'>*必要参数</font>
     * @param id        对象的Id，<font color='red'>如果编辑的时候此Id必须得有</font>
     * @return 如果返回true，则说明该编码已经存在，否则返回false，编码不存在
     */
    public boolean checkCodeUnique(Class<?> clazz, String codeName, String codeValue, Long id) {
        if (StringUtils.isBlank(codeName)) {
            codeName = "code";
        }
        String hql = "select count(t.id) from " + clazz.getSimpleName() + " t where  t." + codeName.trim() + "=? and t.ifDelete=0";
        long c = 0;
        if (id != null) {
            hql += " and t.id!=?";
            c = this.getDao().countByHql(hql, new Object[]{codeValue, id});
        } else {
            c = this.getDao().countByHql(hql, new Object[]{codeValue});
        }
        return c == 0 ? true : false;
    }

    /**
     * 根据id获得对象
     *
     * @param clazz 类型
     * @param id    类型的id
     * @return 返回object
     */
    @SuppressWarnings("rawtypes")
    public Object getById(Class clazz, Serializable id) {
        return this.getDao().get(clazz, id);
    }

    /**
     * 根据hql获得list
     *
     * @param hql    hql语句
     * @param values 值
     * @return 返回list
     */
    @SuppressWarnings("rawtypes")
    public List getListByHql(String hql, Object... values) {
        return this.getDao().getListByHQL(hql, values);
    }

    /**
     * 分页查询
     *
     * @param sql   sql语句
     * @param start 分页开始索引
     * @param end   分页结束索引
     * @param args  值
     * @return 返回page
     */
    @SuppressWarnings({"deprecation", "unchecked", "rawtypes"})
    public Page<Map<String, Object>> findPage(final String sql,
                                              final int start, final int end, final Object... args) {
        List<Map<String, Object>> items = new ArrayList<Map<String, Object>>();

        final String countSql = "select count(*) "
                + DaoHelper.removeSelect(DaoHelper.removeOrders(sql));

        final long count = this.getJdbcTemplate().queryForLong(countSql, args);

        if (count < 1) {
            return new Page(start, end - start, Integer.parseInt(String.valueOf(count)), new ArrayList());
        }

        final String sqlTemp = getLimitString(sql);
        items = this.getJdbcTemplate().query(new PreparedStatementCreator() {
            public PreparedStatement createPreparedStatement(Connection arg0) throws SQLException {
                PreparedStatement pstmt = arg0.prepareStatement(sqlTemp);
                int i = 0;
                if (args != null) {
                    for (i = 0; i < args.length; ++i) {
                        pstmt.setObject(i + 1, args[i]);
                    }
                }
                pstmt.setInt(i + 1, start);
                pstmt.setInt(i + 2, end);
                return pstmt;
            }
        }, new RowMapper() {
            public Object mapRow(ResultSet rs, int row) throws SQLException {
                Map<String, Object> item = new HashMap<String, Object>();
                ResultSetMetaData meta = rs.getMetaData();
                for (int i = 1; i <= meta.getColumnCount(); ++i) {
                    String name = meta.getColumnName(i);
                    Object value = rs.getObject(i);
                    item.put(name, value);
                }
                return item;
            }
        });

        return new Page(start, end - start, Integer.parseInt(String.valueOf(count)),
                items);
    }


    private String getLimitString(String sql) {
        int orderByIndex = sql.toLowerCase().lastIndexOf("order by");
        if (orderByIndex <= 0) {
            throw new UnsupportedOperationException(
                    "must specify 'order by' statement to support limit operation with offset in sql server 2005");
        }
        String sqlOrderBy = sql.substring(orderByIndex + 8);
        String sqlRemoveOrderBy = sql.substring(0, orderByIndex);
        int insertPoint = getSqlAfterSelectInsertPoint(sql);
        return new StringBuffer(sql.length() + 100)
                .append("with tempPagination as(")
                .append(sqlRemoveOrderBy)
                .insert(insertPoint + 23,
                        " ROW_NUMBER() OVER(ORDER BY " + sqlOrderBy + ") as RowNumber,")
                .append(") select * from tempPagination where RowNumber>?  and RowNumber<=?")
                .toString();
    }

    protected static int getSqlAfterSelectInsertPoint(String sql) {
        int selectIndex = sql.toLowerCase().indexOf("select");
        int selectDistinctIndex = sql.toLowerCase().indexOf("select distinct");
        return selectIndex + ((selectDistinctIndex == selectIndex) ? 15 : 6);
    }


    @Override
    public Object getByHQL(String hqlString, Object... values) {
        return getDao().getByHQL(hqlString, values);
    }

    @Override
    public List getListByHQL(String hqlString, Object... values) {
        return getDao().getListByHQL(hqlString, values);
    }

    @Override
    public Page findPageByFetchedHql(String hql, String countHql, int pageNo,
                                     int pageSize, Object... values) {
        return getDao().findPageByFetchedHql(hql, countHql, pageNo, pageSize, values);
    }

    @Override
    public Object get(Class clazz, Serializable id) {
        return getDao().get(clazz, id);
    }

    @Override
    public List findObjectBySql(String queryString, Class pojoClass) {
        return getDao().findObjectBySql(queryString, pojoClass);
    }

    @Override
    public Object getBySQL(String sqlString, Object... values) {
        return getDao().getBySQL(sqlString, values);
    }

    @Override
    public Page findPage(DetachedCriteria dc, int start, int limit) {
        return getDao().findPage(dc, start, limit);
    }

    @Override
    public List findAllByCriteria(DetachedCriteria dc) {
        return getDao().findAllByCriteria(dc);
    }

    @Override
    public List<Map> findListBySql(String queryString) {
        return getDao().findListBySql(queryString);
    }

    @Override
    public Object getFirstByHql(String hql, Object... values) {
        return getDao().findFirstByHQL(hql, values);
    }

    @Override
    public <T> T getObjectByCriteria(DetachedCriteria dc, Class<T> classz) {
        return getDao().getObjectByCriteria(dc, classz);
    }

    @Override
    public List<Map> findMapByHQL(String hqlString, Object... values) {
        return getDao().findMapByHQL(hqlString, values);
    }

    @Override
    public void saveOrUpdate(BaseEntity e) {
        if (e.getId() != null) {
        	e.setUpdateTime(new Date());
            getDao().update(e);
        } else {
            getDao().save(e);
        }
    }
}
