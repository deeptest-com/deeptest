package cn.linkr.testspace.dao.impl;

import java.io.Serializable;
import java.math.BigDecimal;
import java.sql.Connection;
import java.sql.Date;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Timestamp;
import java.sql.Types;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.Map;

import org.hibernate.Criteria;
import org.hibernate.Query;
import org.hibernate.SQLQuery;
import org.hibernate.ScrollableResults;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.criterion.CriteriaSpecification;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Projection;
import org.hibernate.criterion.Projections;
import org.hibernate.internal.CriteriaImpl;
import org.hibernate.jdbc.Work;
import org.hibernate.transform.ResultTransformer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import cn.linkr.testspace.dao.BaseDao;
import cn.linkr.testspace.dao.RowMapper;
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.util.EscColumnToBean;
import cn.linkr.testspace.util.ReflectionUtils;
import cn.linkr.testspace.vo.Page;

/**
 * <简述>base dao的实现
 * <详细描述>
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
@Repository("baseDao")
@SuppressWarnings("all")
public class BaseDaoImpl implements BaseDao {

    /**
     * sessionFactory
     */
    @Autowired
    private SessionFactory sessionFactory;

    /**
     * <保存实体>
     * <完整保存实体>
     *
     * @param t 实体参数
     * @see com.BaseDao.launcher.util.IBaseDao#save(java.lang.Object)
     */
    @Override
    public void save(Object t) {
        this.getSession().save(t);
    }

    /**
     * 立即保存实体
     */
    @Override
    public void flush() {
        this.getSession().flush();
    }

    /**
     * <load>
     * <加载实体的load方法>
     *
     * @param clazz clazz
     * @param id    实体的id
     * @return 查询出来的实体
     * @see cn.linkr.testspace.dao.BaseDao#load(java.io.Serializable)
     */
    @Override
    public Object load(Class clazz, Serializable id) {
        Object load = (Object) this.getSession().load(clazz, id);
        return load;
    }

    /**
     * <get>
     * <查找的get方法>
     *
     * @param clazz clazz
     * @param id    实体的id
     * @return 查询出来的实体
     * @see com.telehot.sjjh.hzhg.dao.IBaseDao#get(Class clazz, java.io.Serializable)
     */
    @Override
    public Object get(Class clazz, Serializable id) {
        Object load = (Object) this.getSession().get(clazz, id);
        return load;
    }

    /**
     * <contains>
     *
     * @param t 实体
     * @return 是否包含
     * @see cn.linkr.testspace.dao.BaseDao#contains(java.lang.Object)
     */
    @Override
    public boolean contains(Object t) {
        return this.getSession().contains(t);
    }

    /**
     * <delete>
     * <删除表中的t数据>
     *
     * @param t 实体
     * @see cn.linkr.testspace.dao.BaseDao#delete(java.lang.Object)
     */
    @Override
    public void delete(Object t) {
        this.getSession().delete(t);
    }

    /**
     * <根据ID删除数据>
     *
     * @param clazz clazz
     * @param id    实体id
     * @return 是否删除成功
     * @see com.telehot.sjjh.hzhg.dao.IBaseDao#deleteById(Class clazz, java.io.Serializable)
     */
    @Override
    public boolean deleteById(Class clazz, Serializable id) {
        Object t = get(clazz, id);
        if (t == null) {
            return false;
        }
        delete(t);
        return true;
    }

    /**
     * <删除所有>
     *
     * @param entities 实体的Collection集合
     * @see cn.linkr.testspace.dao.BaseDao#deleteAll(java.util.Collection)
     */
    @Override
    public void deleteAll(Collection entities) {
        for (Object entity : entities) {
            this.getSession().delete(entity);
        }
    }

    /**
     * <执行Hql语句>
     *
     * @param hqlString hql
     * @param values    不定参数数组
     * @see cn.linkr.testspace.dao.BaseDao#queryHql(java.lang.String, java.lang.Object[])
     */
    @Override
    public void queryHql(String hqlString, Object... values) {
        Query query = this.getSession().createQuery(hqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        query.executeUpdate();
    }

    /**
     * <执行Sql语句>
     *
     * @param sqlString sql
     * @param values    不定参数数组
     * @see cn.linkr.testspace.dao.BaseDao#querySql(java.lang.String, java.lang.Object[])
     */
    @Override
    public void querySql(String sqlString, Object... values) {
        Query query = this.getSession().createSQLQuery(sqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        query.executeUpdate();
    }

    /**
     * <根据HQL语句查找唯一实体>
     *
     * @param hqlString HQL语句
     * @param values    不定参数的Object数组
     * @return 查询实体
     * @see cn.linkr.testspace.dao.BaseDao#getByHQL(java.lang.String, java.lang.Object[])
     */
    @Override
    public Object getByHQL(String hqlString, Object... values) {
        Query query = this.getSession().createQuery(hqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        return (Object) query.uniqueResult();
    }

    /**
     * <根据SQL语句查找唯一实体>
     *
     * @param sqlString SQL语句
     * @param values    不定参数的Object数组
     * @return 查询实体
     * @see cn.linkr.testspace.dao.BaseDao#getBySQL(java.lang.String, java.lang.Object[])
     */
    @Override
    public Object getBySQL(String sqlString, Object... values) {
        Query query = this.getSession().createSQLQuery(sqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        return (Object) query.uniqueResult();
    }

    /**
     * <根据HQL语句，得到对应的list>
     *
     * @param hqlString HQL语句
     * @param values    不定参数的Object数组
     * @return 查询多个实体的List集合
     * @see cn.linkr.testspace.dao.BaseDao#getListByHQL(java.lang.String, java.lang.Object[])
     */
    @Override
    public List<Object> getListByHQL(String hqlString, Object... values) {
        Query query = this.getSession().createQuery(hqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        return query.list();
    }

    public List<Map> findMapByHQL(String hqlString, Object... values) {
        Query query = this.getSession().createQuery(hqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        query.setResultTransformer(CriteriaSpecification.ALIAS_TO_ENTITY_MAP);
        return query.list();
    }

    /**
     * <根据SQL语句，得到对应的list>
     *
     * @param sqlString HQL语句
     * @param values    不定参数的Object数组
     * @return 查询多个实体的List集合
     * @see cn.linkr.testspace.dao.BaseDao#getListBySQL(java.lang.String, java.lang.Object[])
     */
    @Override
    public List<Object> getListBySQL(String sqlString, Object... values) {
        Query query = this.getSession().createSQLQuery(sqlString);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        return query.list();
    }

    /**
     * 由sql语句得到List
     *
     * @param sql    sql语句
     * @param map    map
     * @param values 值
     * @return List
     * @see cn.linkr.testspace.dao.BaseDao#findListBySql(java.lang.String, com.itv.launcher.util.RowMapper, java.lang.Object[])
     */
    @Override
    public List findListBySql(final String sql, final RowMapper map, final Object... values) {
        final List list = new ArrayList();
        // 执行JDBC的数据批量保存
        Work jdbcWork = new Work() {
            public void execute(Connection connection) throws SQLException {
                PreparedStatement ps = null;
                ResultSet rs = null;
                try {
                    ps = connection.prepareStatement(sql);
                    for (int i = 0; i < values.length; i++) {
                        setParameter(ps, i, values[i]);

                    }
                    rs = ps.executeQuery();
                    int index = 0;
                    while (rs.next()) {
                        Object obj = map.mapRow(rs, index++);
                        list.add(obj);

                    }
                } finally {
                    if (rs != null) {
                        rs.close();

                    }
                    if (ps != null) {
                        ps.close();
                    }
                }
            }
        };
        this.getSession().doWork(jdbcWork);
        return list;
    }

    /**
     * <refresh>
     *
     * @param t 实体
     * @see cn.linkr.testspace.dao.BaseDao#refresh(java.lang.Object)
     */
    @Override
    public void refresh(Object t) {
        this.getSession().refresh(t);
    }

    /**
     * <update>
     *
     * @param t 实体
     * @see cn.linkr.testspace.dao.BaseDao#update(java.lang.Object)
     */
    @Override
    public void update(Object t) {
        this.getSession().merge(t);
    }

    /**
     * <根据HQL得到记录数>
     *
     * @param hql    HQL语句
     * @param values 不定参数的Object数组
     * @return 记录总数
     * @see cn.linkr.testspace.dao.BaseDao#countByHql(java.lang.String, java.lang.Object[])
     */
    @Override
    public Long countByHql(String hql, Object... values) {
        Query query = this.getSession().createQuery(hql);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        return (Long) query.uniqueResult();
    }

    /**
     * <HQL分页查询>
     *
     * @param hql      HQL语句
     * @param countHql 查询记录条数的HQL语句
     * @param pageNo   下一页
     * @param pageSize 一页总条数
     * @param values   不定Object数组参数
     * @return Page的封装类，里面包含了页码的信息以及查询的数据List集合
     * @see cn.linkr.testspace.dao.BaseDao#findPageByFetchedHql(java.lang.String, java.lang.String, int, int, java.lang.Object[])
     */
    @Override
    public Page<Object> findPageByFetchedHql(String hql, String countHql, int pageNo, int pageSize, Object... values) {
        Page<Object> retValue = new Page<Object>();
        Query query = this.getSession().createQuery(hql);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        int currentPage = pageNo > 1 ? pageNo : 1;
        retValue.setLimit(pageSize);
        if (countHql == null) {
            ScrollableResults results = query.scroll();
            results.last();
            //设置总记录数
            retValue.setTotal(results.getRowNumber() + 1);
        } else {
            Long count = countByHql(countHql, values);
            retValue.setTotal(count.intValue());
        }
        List<Object> itemList = query.setFirstResult((currentPage - 1) * pageSize).setMaxResults(pageSize).list();
        if (itemList == null) {
            itemList = new ArrayList<Object>();
        }
        retValue.setItems(itemList);
        return retValue;
    }

/*    *//**
     * @return the sessionFactory
     *//*
    public SessionFactory getSessionFactory() {
        return sessionFactory;
    }

    *//**
     * @param sessionFactory the sessionFactory to set
     *//*
    public void setSessionFactory(SessionFactory sessionFactory) {
        this.sessionFactory = sessionFactory;
    }*/

    /**
     * @return session
     */
    @Override
    public Session getSession() {
        //需要开启事物，才能得到CurrentSession
        return sessionFactory.getCurrentSession();
    }

    /**
     * 设置每行批处理参数
     *
     * @param ps   PreparedStatement
     * @param pos  ?占位符索引，从0开始
     * @param data data
     * @throws SQLException 异常
     * @see [类、类#方法、类#成员]
     */
    private void setParameter(PreparedStatement ps, int pos, Object data) throws SQLException {
        if (data == null) {
            ps.setNull(pos + 1, Types.VARCHAR);
            return;
        }
        Class dataCls = data.getClass();
        if (String.class.equals(dataCls)) {
            ps.setString(pos + 1, (String) data);
        } else if (boolean.class.equals(dataCls)) {
            ps.setBoolean(pos + 1, ((Boolean) data));
        } else if (int.class.equals(dataCls)) {
            ps.setInt(pos + 1, (Integer) data);
        } else if (double.class.equals(dataCls)) {
            ps.setDouble(pos + 1, (Double) data);
        } else if (Date.class.equals(dataCls)) {
            Date val = (Date) data;
            ps.setTimestamp(pos + 1, new Timestamp(val.getTime()));
        } else if (BigDecimal.class.equals(dataCls)) {
            ps.setBigDecimal(pos + 1, (BigDecimal) data);
        } else {
            // 未知类型
            ps.setObject(pos + 1, data);
        }

    }

    @Override
    public void executeByHql(String hql, Object... values) {
        Query query = this.getSession().createQuery(hql);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }
        int c = query.executeUpdate();
    }

    @Override
    public List findObjectBySql(String queryString, Class pojoClass) {
        Query query = this.getSession().createSQLQuery(queryString);
        //设置结果集转换器，这是本文重点所在
        query.setResultTransformer(new EscColumnToBean(pojoClass));
        //返回查询结果
        return query.list();
    }

    public List<Map> findListBySql(String queryString) {
        Query query = this.getSession().createSQLQuery(queryString);
        query.setResultTransformer(CriteriaSpecification.ALIAS_TO_ENTITY_MAP);
        return query.list();
    }

    @Override
    public Page findPage(DetachedCriteria dc, int start, int limit) {

        Criteria c = dc.getExecutableCriteria(getSession());
        CriteriaImpl impl = (CriteriaImpl) c;
        Projection projection = impl.getProjection();
        ResultTransformer transformer = impl.getResultTransformer();


        List<CriteriaImpl.OrderEntry> orderEntries = (List<CriteriaImpl.OrderEntry>) ReflectionUtils.getFieldValue(impl, "orderEntries");

        ReflectionUtils.setFieldValue(impl, "orderEntries", new ArrayList<CriteriaImpl.OrderEntry>());


        // 执行Count查询
        c.setResultTransformer(CriteriaImpl.DISTINCT_ROOT_ENTITY);
        long total = (Long) c.setProjection(
                Projections.countDistinct("id")).uniqueResult();


        // 将之前的Projection和OrderBy条件重新设回去
        c.setProjection(projection);
        c.setResultTransformer(transformer);
        ReflectionUtils.setFieldValue(impl, "orderEntries", orderEntries);


        c.setFirstResult(start);
        c.setMaxResults(limit);

        List list = c.list();
        return new Page(start, limit, Integer.parseInt(String.valueOf(total)), (list == null ? new ArrayList() : list));

    }

    @Override
    public <T> List<T> findAllByCriteria(DetachedCriteria dc) {
        Criteria ca = dc.getExecutableCriteria(getSession());
        return (List<T>) ca.list();
    }

    @Override
    public Object findFirstByHQL(String hql, Object... values) {
        Query query = this.getSession().createQuery(hql);
        if (values != null) {
            for (int i = 0; i < values.length; i++) {
                query.setParameter(i, values[i]);
            }
        }

        List<?> list = query.setFirstResult(0).setMaxResults(1).list();

        return (list != null && list.size() > 0) ? list.get(0) : null;
    }

    @Override
    public void saveOrUpdateAll(List<?> list) {
        Session session = getSession();
        for (int i = 0; i < list.size(); i++) {
            Object ob = list.get(i);
            session.saveOrUpdate(ob);
            if (i % 100 == 0) { //20, same as the JDBC batch size
                session.flush();
                session.clear();
            }
        }
    }

    @Override
    public <T> T getObjectByCriteria(DetachedCriteria dc, Class<T> classz) {
        Criteria ca = dc.getExecutableCriteria(getSession());

        return (T) ca.uniqueResult();
    }
    
	@Override
	public Integer moveNode(String nodeTable, Long nodeId, Long newParentId) {
    	Query query = this.getSession().createSQLQuery("CALL move_node(:node_table, :project_id, :parent_id)")
  			  .addEntity(TestProject.class)
  			  .setParameter("node_table", nodeTable)
  			  .setParameter("project_id", nodeId)
  			  .setParameter("parent_id", newParentId);
    	
    	TestProject projec = (TestProject) query.uniqueResult();
	  	return 1;
	}
	
	@Override
	public Integer updateNode(String nodeTable, Long nodeId, String statusName, String statusValue) {
    	Query query = this.getSession().createSQLQuery("CALL update_node(:node_table, :node_id, :status_name, :status_value)")
  			  .addEntity(TestProject.class)
  			  .setParameter("node_table", nodeTable)
  			  .setParameter("node_id", nodeId)
  			  .setParameter("status_name", statusName)
    		  .setParameter("status_value", statusValue);
    	
    	TestProject projec = (TestProject) query.uniqueResult();
	  	return 1;
	}
}
