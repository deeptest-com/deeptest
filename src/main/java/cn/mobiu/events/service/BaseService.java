package cn.mobiu.events.service;

import java.io.Serializable;
import java.util.List;
import java.util.Map;

import org.hibernate.criterion.DetachedCriteria;

import cn.mobiu.events.entity.BaseEntity;
import cn.mobiu.events.vo.Page;

public interface BaseService {

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
    List getListByHQL(String hqlString, Object... values);

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

    Object get(Class clazz, Serializable id);

    /**
     * 将sql查询的结果集自动转成java的对象
     *
     * @param queryString 查询语句
     * @param pojoClass   要匹配的java类型
     * @return 返回List
     */
    @SuppressWarnings("rawtypes")
    List findObjectBySql(String queryString, Class pojoClass);

    @SuppressWarnings("rawtypes")
    List<Map> findListBySql(String queryString);

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
    @SuppressWarnings("rawtypes")
    List findAllByCriteria(DetachedCriteria dc);

    Object getFirstByHql(String hql, Object... values);

    <T> T getObjectByCriteria(DetachedCriteria dc, Class<T> classz);

    List<Map> findMapByHQL(String hqlString, Object... values);

    void saveOrUpdate(BaseEntity e);
}
