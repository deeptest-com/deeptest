package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.entity.TestProject.ProjectType;
import com.ngtesting.platform.entity.TestProjectAccessHistory;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.hibernate.FetchMode;
import org.hibernate.Filter;
import org.hibernate.criterion.CriteriaSpecification;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

@Service
public class ProjectServiceImpl extends BaseServiceImpl implements
		ProjectService {

	private static final Log log = LogFactory
			.getLog(ProjectService.class);

	@Autowired
	private ProjectDao projectDao;
    @Autowired
    private CaseService caseService;

	@Override
	// @Cacheable(value="orgProjects",key="#orgId.toString().concat('_').concat(#disabled)")
	public List<TestProjectVo> listVos(Long orgId, String keywords, String disabled) {
		// CacheManager manager = CacheManager.create();
		// net.sf.ehcache.Cache cache = manager.getCache("orgProjects");
		// String key = orgId + "_" + disabled;
		// Element el = null;
		// if(cache.isKeyInCache(key)){
		// el = cache.get(key);
		// return (Map<String, Object>)el.getObjectValue();
		// }
		
		List<TestProject> pos = list(orgId, keywords, disabled);

		List<TestProjectVo> vos = this.genVos(pos, keywords, disabled);

		// el = new Element(key, ret);
		// cache.put(el);
		return vos;
	}
	
	@Override
	public List<TestProject> list(Long orgId, String keywords, String disabled) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);

		dc.add(Restrictions.eq("orgId", orgId));
		if (StringUtil.isNotEmpty(disabled)) {
			dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
		}

		dc.add(Restrictions.eq("type", ProjectType.group));
		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.addOrder(Order.asc("id"));
		
		dc.setFetchMode("children", FetchMode.JOIN);
		dc.setResultTransformer(CriteriaSpecification.DISTINCT_ROOT_ENTITY); 
		
		Filter filter = getDao().getSession().enableFilter("filter_project_deleted");
		filter.setParameter("isDeleted", Boolean.valueOf(false));
		List<TestProject> pos = findAllByCriteria(dc);
		getDao().getSession().disableFilter("filter_project_deleted");

		return pos;
	}
	
	@Override
	public List<TestProjectVo> listProjectGroups(Long orgId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);

		dc.add(Restrictions.eq("orgId", orgId));
		dc.add(Restrictions.eq("type", ProjectType.group));
		dc.add(Restrictions.eq("disabled", false));
		
		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.addOrder(Order.asc("id"));
		
		List<TestProject> pos = findAllByCriteria(dc);

		List<TestProjectVo> vos = this.genGroupVos(pos);

		return vos;
	}
	
	@Override
	public List<TestProjectAccessHistory> listRecentProject(Long orgId, Long userId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectAccessHistory.class);

		dc.add(Restrictions.eq("orgId", orgId));
		dc.add(Restrictions.eq("userId", userId));
		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", false));
		dc.addOrder(Order.desc("lastAccessTime"));
		 
		List<TestProjectAccessHistory> pos = findPage(dc, 0, 4).getItems();

		return pos;
	}
	@Override
	public List<TestProjectAccessHistoryVo> listRecentProjectVo(Long orgId, Long userId) {
		List<TestProjectAccessHistory> pos = listRecentProject(orgId, userId);
		List<TestProjectAccessHistoryVo> vos = genHistoryVos(pos);

		return vos;
	}

	@Override
	public TestProject getDetail(Long id) {
		if (id == null) {
			return null;
		}
		TestProject po = (TestProject) get(TestProject.class, id);

		return po;
	}

	@Override
	public TestProject save(TestProjectVo vo, Long orgId, Long userId) {
		if (vo == null) {
			return null;
		}

		boolean isNew = vo.getId() == null;
		TestProject po = new TestProject();
		if (isNew) {
            po.setOrgId(orgId);
		} else {
            po = (TestProject) get(TestProject.class, vo.getId());
		}
		
		boolean disableChanged = vo.getDisabled() != po.getDisabled();
		
		po.setParentId(vo.getParentId());
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setType(ProjectType.valueOf(vo.getType()));
		po.setDisabled(vo.getDisabled());

		saveOrUpdate(po);

        if(isNew && ProjectType.project.equals(ProjectType.valueOf(vo.getType()))) {
            caseService.createRoot(po.getId(), userId);
        }
		
		if (!disableChanged) {
			return po;
		}
		
		// 项目被启用
		if (!po.getDisabled()) {
			if (po.getType().equals(ProjectType.project)) {
				// 启用父
				TestProject parent = po.getParent();
				if (parent.getDisabled()) {
					parent.setDisabled(false);
					saveOrUpdate(parent);
				}
			} else {
				// 启用子
				for (TestProject child : po.getChildren()) {
					if (child.getDisabled()) {
						child.setDisabled(false);
						saveOrUpdate(child);
					}
				}
			}
		}
		
		// 项目组被归档，归档子项目
		if (po.getDisabled() && po.getType().equals(ProjectType.group)) {
			for (TestProject child : po.getChildren()) {
				if (!child.getDisabled()) {
					child.setDisabled(true);
					saveOrUpdate(child);
				}
			}
		}

		// this.removeCache(user.getOrgId());

		return po;
	}

	@Override
	public Boolean delete(Long id) {
		if (id == null) {
			return null;
		}

		TestProject po = (TestProject) get(TestProject.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		// 项目组被删除，删除子项目
		if (po.getType().equals(ProjectType.group)) {
			for (TestProject child : po.getChildren()) {
				child.setDeleted(true);
				saveOrUpdate(child);
			}
			
		}

		return true;
	}
	
	@Override
	public List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, Long projectId, UserVo userVo) {
		TestUser user = (TestUser) get(TestUser.class, userVo.getId());
		
		List<TestProjectAccessHistoryVo> recentProjects = listRecentProjectVo(orgId, userVo.getId());
		if (recentProjects.size() > 0) {
			user.setDefaultProjectId(recentProjects.get(0).getId());
		}
		
		saveOrUpdate(user);
		
		userVo.setDefaultProjectId(user.getDefaultProjectId());
		
		return recentProjects;
	}

	// @Override
	// public void removeCache(Long orgId) {
	// CacheManager manager = CacheManager.create();
	// net.sf.ehcache.Cache cache = manager.getCache("orgProjects");
	// String prefix = orgId + "_";
	// if(cache.isKeyInCache(prefix + "true")){
	// cache.remove(prefix + "true");
	// }
	// if(cache.isKeyInCache(prefix)){
	// cache.remove(prefix);
	// }
	// }

	@Override
	public List<TestProjectVo> genVos(List<TestProject> pos) {
		return this.genVos(pos, null, null);
	}
	@Override
	public List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled) {
		List<TestProjectVo> voList = new LinkedList<TestProjectVo>();
		for (TestProject po : pos) {
			TestProjectVo vo = genVo(po);
			voList.add(vo);
			
			List<TestProjectVo> voList2 = new LinkedList<TestProjectVo>();
			List<TestProject> children = po.getChildren();
			for (TestProject child : children) {
				if ( (StringUtil.IsEmpty(keywords) || child.getName().toLowerCase().indexOf(keywords.toLowerCase()) > -1) 
						&& ( StringUtil.IsEmpty(disabled) || child.getDisabled() == Boolean.valueOf(disabled)) ) {
					TestProjectVo childVo = genVo(child);
					voList2.add(childVo);
				}
			}
			vo.setChildrenNumb(voList2.size());
			voList.addAll(voList2);
		}
		
		return voList;
	}
	@Override
	public List<TestProjectVo> genGroupVos(List<TestProject> pos) {
		List<TestProjectVo> voList = new LinkedList<TestProjectVo>();
		for (TestProject po : pos) {
				TestProjectVo vo = genVo(po);
				voList.add(vo);
		}
		
		return voList;
	}

	@Override
	public TestProjectVo genVo(TestProject po) {
		if (po == null) {
			return null;
		}
		TestProjectVo vo = new TestProjectVo();
		BeanUtilEx.copyProperties(vo, po);
		if (po.getParentId() == null) {
			vo.setParentId(null);
		}

		return vo;
	}
	
	@Override
	public List<TestProjectAccessHistoryVo> genHistoryVos(List<TestProjectAccessHistory> pos) {
		List<TestProjectAccessHistoryVo> voList = new LinkedList<TestProjectAccessHistoryVo>();
		for (TestProjectAccessHistory po : pos) {
			TestProjectAccessHistoryVo vo = genHistoryVo(po);
			voList.add(vo);
		}
		
		return voList;
	}
	
	@Override
	public TestProjectAccessHistoryVo genHistoryVo(TestProjectAccessHistory po) {
		if (po == null) {
			return null;
		}
		TestProjectAccessHistoryVo vo = new TestProjectAccessHistoryVo();
		BeanUtilEx.copyProperties(vo, po);
		vo.setProjectName(po.getProjectName());
		
		return vo;
	}

	@Override
	public TestProjectVo viewPers(UserVo userVo, Long projectId) {
		TestUser user = (TestUser) get(TestUser.class, userVo.getId());
		TestProject project = getDetail(projectId);
		
		TestProjectAccessHistory history = getHistory(project.getOrgId(), userVo.getId(), projectId, project.getName());
		history.setLastAccessTime(new Date());
		saveOrUpdate(history);
		
		if (user.getDefaultProjectId() != projectId) {
			user.setDefaultProjectId(projectId);
			
			saveOrUpdate(user);
			
			userVo.setDefaultProjectId(projectId);
		}
		
		TestProjectVo vo = genVo(project);
		return vo;
	}

	private TestProjectAccessHistory getHistory(Long orgId, Long userId, Long projectId, String projectName) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectAccessHistory.class);

		dc.add(Restrictions.eq("orgId", orgId));
		dc.add(Restrictions.eq("projectId", projectId));
		dc.add(Restrictions.eq("userId", userId));
		dc.add(Restrictions.eq("deleted", false));
		dc.add(Restrictions.eq("disabled", false));
		 
		TestProjectAccessHistory history;
		List<TestProjectAccessHistory> pos = findAllByCriteria(dc);
		if (pos.size() > 0) {
			history = pos.get(0);
		} else {
			history = new TestProjectAccessHistory(orgId, userId, projectId, projectName);
		}
		return history;
	}

	@Override
	public boolean isLastestProjectGroup(Long orgId, Long projectGroupId) {
		String hql = "select count(prj.id) from TestProject prj where prj.id != ?"
				+ " and orgId=? and type=? and prj.deleted != true and prj.deleted != true";

		long count = (Long) getByHQL(hql, orgId, projectGroupId, ProjectType.group);
		return count == 0;
	}

}
