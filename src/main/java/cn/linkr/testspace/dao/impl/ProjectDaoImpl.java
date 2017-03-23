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

import javax.persistence.ParameterMode;

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
import org.hibernate.procedure.ProcedureCall;
import org.hibernate.procedure.ProcedureOutputs;
import org.hibernate.result.Output;
import org.hibernate.result.ResultSetOutput;
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
	public TestProject moveProject(Long projectId, Long newParentId) {
    	Query query = this.getSession().createSQLQuery("CALL move_node(:node_table, :project_id, :parent_id)")
  			  .addEntity(TestProject.class)
  			  .setParameter("node_table", "tst_project")
  			  .setParameter("project_id", projectId)
  			  .setParameter("parent_id", newParentId);
  	
	  	List<TestProject> ls = query.list();
	  	return ls.get(0);
	}
    
    @Override
    public List<TestProject> findProjectByProcedure(Long companyId, Boolean isActive) {
    	Query query = this.getSession().createSQLQuery("CALL query_project(:company_id, :is_active)")
    			  .addEntity(TestProject.class)
    			  .setParameter("company_id", companyId)
    			  .setParameter("is_active", isActive);
    	
    	List ls = query.list();
    	return ls;
    }
    
    public Session getSession() {
        return sessionFactory.getCurrentSession();
    }

}
