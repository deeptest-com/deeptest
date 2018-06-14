package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.entity.TestHistory;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.entity.TestProject.ProjectType;
import com.ngtesting.platform.entity.TestProjectAccessHistory;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.ProjectPrivilegeService;
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

import java.util.*;

@Service
public class ProjectServiceImpl extends BaseServiceImpl implements
		ProjectService {

	private static final Log log = LogFactory
			.getLog(ProjectService.class);

    @Autowired
    HistoryService historyService;
	@Autowired
	private ProjectDao projectDao;
    @Autowired
    private CaseService caseService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;

	@Override
	public List<TestProjectVo> listVos(Long orgId, Long userId, String keywords, String disabled) {
		Map<String, Map<String, Boolean>> privMap = new HashMap();
		List<Object[]> ls = getDao().getListBySQL("{call get_project_privilege_by_org_for_user(?,?)}",
				userId, orgId);
		for (Object[] arr : ls) {
		    if (privMap.get(arr[0].toString()) == null) {
                privMap.put(arr[0].toString(), new HashMap());
            }

			String str = arr[1].toString() + "-" + arr[2].toString();
            privMap.get(arr[0].toString()).put(str, true);
		}

		List<TestProject> pos = list(orgId, keywords, disabled);
		List<TestProjectVo> vos = this.genVos(pos, keywords, disabled, privMap);

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
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));

        dc.createAlias("project", "project");
//		dc.add(Restrictions.ne("project.deleted", true));
        dc.add(Restrictions.ne("project.deleted", true));
		dc.add(Restrictions.ne("project.disabled", true));

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
	public TestProject save(TestProjectVo vo, Long orgId, UserVo userVo) {
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

        if(isNew && ProjectType.project.equals(po.getType())) {
            projectPrivilegeService.addUserAsProjectTestLeaderPers(orgId, po.getId(), "test_leader", userVo.getId());
            caseService.createRoot(po.getId(), userVo);
        }
        if(ProjectType.project.equals(po.getType())) {
            historyService.create(po.getId(), userVo,
                    isNew?Constant.MsgType.create.msg:Constant.MsgType.create.update.msg,
                    TestHistory.TargetType.project, po.getId(), po.getName());
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
	public List<TestProjectVo> listBrothers(Long projectId) {
	    TestProject project = (TestProject)get(TestProject.class, projectId);
        DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);

        dc.add(Restrictions.eq("parentId", project.getParentId()));
        dc.add(Restrictions.eq("type", ProjectType.project));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));

        List<TestProject> pos = findAllByCriteria(dc);

        List<TestProjectVo> vos = this.genVos(pos, null);
        return vos;
	}
    @Override
    public List<Long> listBrotherIds(Long projectId) {
        String hql = "select prj.id from TestProject prj where prj.parentId=" +
                "(select p.parentId from TestProject p where p.id=?)"
                + " and type=? and prj.deleted != true and prj.deleted != true";

        List<Long> ids = getListByHQL(hql, projectId, ProjectType.project);
        return ids;
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
	public TestProjectVo viewPers(Long projectId, UserVo userVo) {
		TestProject project = getDetail(projectId);

        TestUser userPo = (TestUser)get(TestUser.class, userVo.getId());
        if (project.getType().equals(ProjectType.project)) {
            genHistoryPers(project.getOrgId(), userVo.getId(), projectId, project.getName());

			userPo.setDefaultPrjId(projectId);
			saveOrUpdate(userPo);

			userVo.setDefaultPrjId(projectId);
			userVo.setDefaultPrjName(project.getName());
		}

		TestProjectVo vo = genVo(project, null);
		return vo;
	}

    @Override
    public void updateNameInHisotyPers(Long projectId, Long userId) {
        TestProject project = getDetail(projectId);
        genHistoryPers(project.getOrgId(), userId, projectId, project.getName());
    }

    @Override
    public void genHistoryPers(Long orgId, Long userId, Long projectId, String projectName) {
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
			history.setProjectName(projectName);
		} else {
			history = new TestProjectAccessHistory(orgId, userId, projectId, projectName);
		}
        history.setLastAccessTime(new Date());
        saveOrUpdate(history);
	}

	@Override
	public boolean isLastestProjectGroup(Long orgId, Long projectGroupId) {
		String hql = "select count(prj.id) from TestProject prj where prj.id != ?"
				+ " and orgId=? and type=? and prj.deleted != true and prj.deleted != true";

		long count = (Long) getByHQL(hql, orgId, projectGroupId, ProjectType.group);
		return count == 0;
	}

    @Override
    public List<TestProjectVo> genVos(List<TestProject> pos, Map<String, Map<String, Boolean>> privMap) {
        return this.genVos(pos, null, null, privMap);
    }
    @Override
    public List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled, Map<String, Map<String, Boolean>> privMap) {
        List<TestProjectVo> voList = new LinkedList<TestProjectVo>();
        for (TestProject po : pos) {
            TestProjectVo vo = genVo(po, privMap);
            voList.add(vo);

            List<TestProjectVo> voList2 = new LinkedList<TestProjectVo>();
            List<TestProject> children = po.getChildren();
            boolean childCanView = false;
            for (TestProject child : children) {
                if ( (StringUtil.IsEmpty(keywords) || child.getName().toLowerCase().indexOf(keywords.toLowerCase()) > -1)
                        && ( StringUtil.IsEmpty(disabled) || child.getDisabled() == Boolean.valueOf(disabled)) ) {
                    TestProjectVo childVo = genVo(child, privMap);
                    voList2.add(childVo);

                    if (childVo.getPrivs().get("project-view") ) {
                        childCanView = true;
                    }
                }
            }
            vo.setChildrenNumb(voList2.size());
            voList.addAll(voList2);

            if (childCanView) {
                vo.getPrivs().put("project-view", true);
            }
        }

        return voList;
    }
    @Override
    public List<TestProjectVo> genGroupVos(List<TestProject> pos) {
        List<TestProjectVo> voList = new LinkedList<TestProjectVo>();
        for (TestProject po : pos) {
            TestProjectVo vo = genVo(po, null);
            voList.add(vo);
        }

        return voList;
    }

    @Override
    public TestProjectVo genVo(TestProject po, Map<String, Map<String, Boolean>> privMap) {
        if (po == null) {
            return null;
        }
        TestProjectVo vo = new TestProjectVo();
        BeanUtilEx.copyProperties(vo, po);
        if (po.getParentId() == null) {
            vo.setParentId(null);
        }

        if (privMap != null && privMap.get(po.getId().toString()) != null) {
            vo.setPrivs(privMap.get(po.getId().toString()));
        }

        return vo;
    }

}
