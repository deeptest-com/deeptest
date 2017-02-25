package cn.linkr.testspace.dao;

import java.io.Serializable;
import java.util.Collection;
import java.util.List;
import java.util.Map;

import org.hibernate.criterion.DetachedCriteria;

import cn.linkr.testspace.vo.Page;

/**
 * <简述>baseDao 接口
 * <详细描述>
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
public interface IBaseDao {

    /**
     * 保存实体
     * 完整保存实体
     *
     * @param t 对象
     * @author xuxiang
     */
    void save(Object t);

    /**
     * <load>
     * <加载实体的load方法>
     *
     * @param clazz clazz
     * @param id    实体的id
     * @return 查询出来的实体
     */
    @SuppressWarnings("rawtypes")
    Object load(Class clazz, Serializable id);

    /**
     * <get>
     * <查找的get方法>
     *
     * @param clazz clazz
     * @param id    实体的id
     * @return 查询出来的实体
     */
    @SuppressWarnings("rawtypes")
    Object get(Class clazz, Serializable id);

    /**
     * <contains>
     *
     * @param t 实体
     * @return 是否包含
     */
    boolean contains(Object t);

    /**
     * <delete>
     * <删除表中的t数据>
     *
     * @param t 实体
     */
    void delete(Object t);

    /**
     * <根据ID删除数据>
     *
     * @param clazz clazz
     * @param id    实体id
     * @return 是否删除成功
     */
    @SuppressWarnings("rawtypes")
    boolean deleteById(Class clazz, Serializable id);

    /**
     * <删除所有>
     *
     * @param entities 实体的Collection集合
     */
    @SuppressWarnings("rawtypes")
    void deleteAll(Collection entities);

    /**
     * <执行Hql语句>
     *
     * @param hqlString hql
     * @param values    不定参数数组
     */
    void queryHql(String hqlString, Object... values);

    /**
     * <执行Sql语句>
     *
     * @param sqlString sql
     * @param values    不定参数数组
     */
    void querySql(String sqlString, Object... values);

    /**
     * <根据HQL语句查找唯一实体>
     *
     * @param hqlString HQL语句
     * @param values    不定参数的Object数组
     * @return 查询实体
     */
    Object getByHQL(String hqlString, Object... values);

    /**
     * <根据SQL语句查找唯一实体>
     *
     * @param sqlString SQL语句
     * @param values    不定参数的Object数组
     * @return 查询实体
     */
    Object getBySQL(String sqlString, Object... values);

    /**
     * <根据HQL语句，得到对应的list>
     *
     * @param hqlString HQL语句
     * @param values    不定参数的Object数组
     * @return 查询多个实体的List集合
     */
    @SuppressWarnings("rawtypes")
    List getListByHQL(String hqlString, Object... values);

    Object findFirstByHQL(String hqlString, Object... args);

    /**
     * <根据SQL语句，得到对应的list>
     *
     * @param sqlString HQL语句
     * @param values    不定参数的Object数组
     * @return 查询多个实体的List集合
     */
    @SuppressWarnings("rawtypes")
    List getListBySQL(String sqlString, Object... values);

    /**
     * 由sql语句得到List
     *
     * @param sql    sql语句
     * @param map    map
     * @param values values
     * @return List
     */
    @SuppressWarnings("rawtypes")
    List findListBySql(final String sql, final RowMapper map, final Object... values);

    @SuppressWarnings("rawtypes")
    List<Map> findListBySql(String queryString);

    /**
     * <refresh>
     *
     * @param t 实体
     */
    void refresh(Object t);

    /**
     * <update>
     *
     * @param t 实体
     */
    void update(Object t);

    /**
     * <根据HQL得到记录数>
     *
     * @param hql    HQL语句
     * @param values 不定参数的Object数组
     * @return 记录总数
     */
    Long countByHql(String hql, Object... values);

    /**
     * 执行hql语句
     *
     * @param hql    hql语句
     * @param values 值
     */
    void executeByHql(String hql, Object... values);

    /**
     * <HQL分页查询>
     *
     * @param hql      HQL语句
     * @param countHql 查询记录条数的HQL语句
     * @param pageNo   下一页
     * @param pageSize 一页总条数
     * @param values   不定Object数组参数
     * @return PageResults的封装类，里面包含了页码的信息以及查询的数据List集合
     */
    @SuppressWarnings("rawtypes")
    Page findPageByFetchedHql(String hql, String countHql, int pageNo, int pageSize, Object... values);

    /**
     * 将sql查询的结果集自动转成java的对象
     *
     * @param queryString 查询语句
     * @param pojoClass   要匹配的java类型
     * @return 返回List
     */
    @SuppressWarnings("rawtypes")
    List findObjectBySql(String queryString, Class pojoClass);

    /**
     * 使用离线Criteria查询返回page
     *
     * @param dc
     * @param start
     * @param limit
     * @return
     */
    @SuppressWarnings("rawtypes")
    Page findPage(DetachedCriteria dc, int start, int limit);

    /**
     * 使用离线Criteria查询返回所有值
     *
     * @param dc
     * @return
     */
    <T> List<T> findAllByCriteria(DetachedCriteria dc);

    /**
     * 批量save或者保存数据
     */
    void saveOrUpdateAll(List<?> list);

    <T> T getObjectByCriteria(DetachedCriteria dc, Class<T> classz);

    @SuppressWarnings("rawtypes")
    public List<Map> findMapByHQL(String hqlString, Object... values);

    void flush();
}
