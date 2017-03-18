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
import cn.linkr.testspace.dao.ProjectDao;
import cn.linkr.testspace.dao.RowMapper;
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.util.EscColumnToBean;
import cn.linkr.testspace.util.ReflectionUtils;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.TestProjectVo;

@Repository("projectDao")
@SuppressWarnings("all")
public class ProjectDaoImpl implements ProjectDao {

    @Autowired
    private SessionFactory sessionFactory;
    
    @Override
    public List findProjectByProcedure(Long companyId, Long parentId) {
    	String functionStr = "queryProjectChildren(" + companyId + "," + parentId + ")";
    	String queryString = "select " + functionStr;
        Query query = this.getSession().createSQLQuery(queryString);
        query.setResultTransformer(CriteriaSpecification.ALIAS_TO_ENTITY_MAP);
        List<Map<String, String>> ls = query.list();
        String ids = ls.get(0).get(functionStr).replace("$,", "");
        
        String selectStr = "select * from tst_project where id in (" + ids + ")";
    	SQLQuery select = this.getSession().createSQLQuery(selectStr);
        
    	select.setResultTransformer(CriteriaSpecification.ALIAS_TO_ENTITY_MAP);
        List<Map> ls2 = select.list();
        return ls2;
    }
    
    public Session getSession() {
        return sessionFactory.getCurrentSession();
    }
}
