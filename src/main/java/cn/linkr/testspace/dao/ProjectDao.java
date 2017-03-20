package cn.linkr.testspace.dao;

import java.io.Serializable;
import java.util.Collection;
import java.util.List;
import java.util.Map;

import org.hibernate.criterion.DetachedCriteria;

import cn.linkr.testspace.vo.Page;

public interface ProjectDao {
	 public List findProjectByProcedure(Long companyId, Boolean isActive, String keywords);
}
