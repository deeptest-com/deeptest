package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class ProjectServiceImpl extends BaseServiceImpl implements ProjectService {

	private static final Log log = LogFactory.getLog(ProjectServiceImpl.class);

    @Autowired
	HistoryService historyService;
	@Autowired
	private ProjectDao projectDao;
    @Autowired
    private CaseService caseService;
    @Autowired
	ProjectPrivilegeService projectPrivilegeService;

	@Override
	public List<TstProject> list(Integer orgId, Integer userId, String keywords, Boolean disabled) {
		Map<String, Map<String, Boolean>> privMap = new HashMap();
        List<Map<String, String>> projectPrivs = projectDao.getProjectPrivilegeByOrgForUser(userId, orgId);
        for (Map<String, String> map : projectPrivs) {
		    if (privMap.get(map.get("projectId")) == null) {
		        String prjId = map.get("projectId");
                privMap.put(prjId, new HashMap());
            }

			String str = map.get("code") + "-" + map.get("action");
            privMap.get(map.get("projectId").toString()).put(str, true);
		}

        List<TstProject> pos = projectDao.query(orgId, keywords, disabled);
		List<TstProject> vos = this.genVos(pos, privMap);

		return vos;
	}

	@Override
	public List<TstProject> list(Integer orgId, String keywords, Boolean disabled) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProject.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		if (StringUtil.isNotEmpty(disabled)) {
//			dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
//		}
//
//		dc.add(Restrictions.eq("type", ProjectType.group));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.addOrder(Order.asc("id"));
//
//		dc.setFetchMode("children", FetchMode.JOIN);
//		dc.setResultTransformer(CriteriaSpecification.DISTINCT_ROOT_ENTITY);
//
//		Filter filter = getDao().getSession().enableFilter("filter_project_deleted");
//		filter.setParameter("isDeleted", Boolean.valueOf(false));
//		List<TstProject> pos = findAllByCriteria(dc);
//		getDao().getSession().disableFilter("filter_project_deleted");
//
//		return pos;

		return null;
	}

	@Override
	public List<TstProject> listProjectGroups(Integer orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProject.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("type", ProjectType.group));
//		dc.add(Restrictions.eq("disabled", false));
//
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.addOrder(Order.asc("id"));
//
//		List<TstProject> pos = findAllByCriteria(dc);
//
//		List<TstProject> vos = this.genGroupVos(pos);
//
//		return vos;

		return null;
	}

	@Override
	public List<TstProjectAccessHistory> listRecentProject(Integer orgId, Integer userId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectAccessHistory.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("userId", userId));
//		dc.add(Restrictions.ne("deleted", true));
//		dc.add(Restrictions.ne("disabled", true));
//
//        dc.createAlias("project", "project");
////		dc.add(Restrictions.ne("project.deleted", true));
//        dc.add(Restrictions.ne("project.deleted", true));
//		dc.add(Restrictions.ne("project.disabled", true));
//
//		dc.addOrder(Order.desc("lastAccessTime"));
//
//		List<TstProjectAccessHistory> pos = findPage(dc, 0, 4).getItems();
//
//		return pos;

		return null;
	}
	@Override
	public List<TstProjectAccessHistory> listRecentProjectVo(Integer orgId, Integer userId) {
//		List<TstProjectAccessHistory> pos = listRecentProject(orgId, userId);
//		List<TstProjectAccessHistory> vos = genHistoryVos(pos);
//
//		return vos;

		return null;
	}

	@Override
	public TstProject getDetail(Integer id) {
//		if (id == null) {
//			return null;
//		}
//		TstProject po = (TstProject) get(TstProject.class, id);
//
//		return po;

		return null;
	}

	@Override
	public TstProject save(TstProject vo, Integer orgId, TstUser TstUser) {
//		if (vo == null) {
//			return null;
//		}
//
//		boolean isNew = vo.getId() == null;
//		TstProject po = new TstProject();
//		if (isNew) {
//            po.setOrgId(orgId);
//		} else {
//            po = (TstProject) get(TstProject.class, vo.getId());
//		}
//
//		boolean disableChanged = vo.getDisabled() != po.getDisabled();
//
//		po.setParentId(vo.getParentId());
//		po.setName(vo.getName());
//		po.setDescr(vo.getDescr());
//		po.setType(ProjectType.valueOf(vo.getType()));
//		po.setDisabled(vo.getDisabled());
//
//		saveOrUpdate(po);
//
//        if(isNew && ProjectType.project.equals(po.getType())) {
//            projectPrivilegeService.addUserAsProjectTestLeaderPers(orgId, po.getId(), "test_leader", TstUser.getId());
//            caseService.createRoot(po.getId(), TstUser);
//        }
//        if(ProjectType.project.equals(po.getType())) {
//            historyService.create(po.getId(), TstUser,
//                    isNew? Constant.MsgType.create.msg: Constant.MsgType.create.update.msg,
//                    TestHistory.TargetType.project, po.getId(), po.getName());
//        }
//
//		if (!disableChanged) {
//			return po;
//		}
//
//		// 项目被启用
//		if (!po.getDisabled()) {
//			if (po.getType().equals(ProjectType.project)) {
//				// 启用父
//				TstProject parent = po.getParent();
//				if (parent.getDisabled()) {
//					parent.setDisabled(false);
//					saveOrUpdate(parent);
//				}
//			} else {
//				// 启用子
//				for (TstProject child : po.getChildren()) {
//					if (child.getDisabled()) {
//						child.setDisabled(false);
//						saveOrUpdate(child);
//					}
//				}
//			}
//		}
//
//		// 项目组被归档，归档子项目
//		if (po.getDisabled() && po.getType().equals(ProjectType.group)) {
//			for (TstProject child : po.getChildren()) {
//				if (!child.getDisabled()) {
//					child.setDisabled(true);
//					saveOrUpdate(child);
//				}
//			}
//		}
//
//		// this.removeCache(user.getOrgId());
//
//		return po;

		return null;
	}

	@Override
	public Boolean delete(Integer id) {
//		if (id == null) {
//			return null;
//		}
//
//		TstProject po = (TstProject) get(TstProject.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);
//
//		// 项目组被删除，删除子项目
//		if (po.getType().equals(ProjectType.group)) {
//			for (TstProject child : po.getChildren()) {
//				child.setDeleted(true);
//				saveOrUpdate(child);
//			}
//
//		}

		return true;
	}

	@Override
	public List<TstProject> listBrothers(Integer projectId) {
//	    TstProject project = (TstProject)get(TstProject.class, projectId);
//        DetachedCriteria dc = DetachedCriteria.forClass(TstProject.class);
//
//        dc.add(Restrictions.eq("parentId", project.getParentId()));
//        dc.add(Restrictions.eq("type", ProjectType.project));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//
//        List<TstProject> pos = findAllByCriteria(dc);
//
//        List<TstProject> vos = this.genVos(pos, null);
//        return vos;

		return null;
	}
    @Override
    public List<Integer> listBrotherIds(Integer projectId) {
//        String hql = "select prj.id from TstProject prj where prj.parentId=" +
//                "(select p.parentId from TstProject p where p.id=?)"
//                + " and type=? and prj.deleted != true and prj.deleted != true";
//
//        List<Integer> ids = getListByHQL(hql, projectId, ProjectType.project);
//        return ids;

		return null;
    }

	@Override
	public List<TstProjectAccessHistory> genHistoryVos(List<TstProjectAccessHistory> pos) {
		List<TstProjectAccessHistory> voList = new LinkedList<TstProjectAccessHistory>();
//		for (TstProjectAccessHistory po : pos) {
//			TstProjectAccessHistory vo = genHistoryVo(po);
//			voList.add(vo);
//		}

		return voList;
	}

	@Override
	public TstProjectAccessHistory genHistoryVo(TstProjectAccessHistory po) {
//		if (po == null) {
//			return null;
//		}
//		TstProjectAccessHistory vo = new TstProjectAccessHistory();
//		BeanUtilEx.copyProperties(vo, po);
//		vo.setProjectName(po.getProjectName());
//
//		return vo;

		return null;
	}

	@Override
	public TstProject viewPers(Integer projectId, TstUser TstUser) {
		TstProject project = getDetail(projectId);

//        TestUser userPo = (TestUser)get(TestUser.class, TstUser.getId());
//        if (project.getType().equals(ProjectType.project)) {
//            genHistoryPers(project.getOrgId(), TstUser.getId(), projectId, project.getName());
//
//			userPo.setDefaultPrjId(projectId);
//			saveOrUpdate(userPo);
//
//			TstUser.setDefaultPrjId(projectId);
//			TstUser.setDefaultPrjName(project.getName());
//		}

		TstProject vo = genVo(project, null);
		return vo;
	}

    @Override
    public void updateNameInHisotyPers(Integer projectId, Integer userId) {
//        TstProject project = getDetail(projectId);
//        genHistoryPers(project.getOrgId(), userId, projectId, project.getName());
    }

    @Override
    public void genHistoryPers(Integer orgId, Integer userId, Integer projectId, String projectName) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectAccessHistory.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("projectId", projectId));
//		dc.add(Restrictions.eq("userId", userId));
//		dc.add(Restrictions.eq("deleted", false));
//		dc.add(Restrictions.eq("disabled", false));
//
//		TstProjectAccessHistory history;
//		List<TstProjectAccessHistory> pos = findAllByCriteria(dc);
//		if (pos.size() > 0) {
//			history = pos.get(0);
//			history.setProjectName(projectName);
//		} else {
//			history = new TstProjectAccessHistory(orgId, userId, projectId, projectName);
//		}
//        history.setLastAccessTime(new Date());
//        saveOrUpdate(history);
	}

	@Override
	public boolean isLastestProjectGroup(Integer orgId, Integer projectGroupId) {
//		String hql = "select count(prj.id) from TstProject prj where prj.id != ?"
//				+ " and orgId=? and type=? and prj.deleted != true and prj.deleted != true";
//
//		long count = (Integer) getByHQL(hql, orgId, projectGroupId, ProjectType.group);
//		return count == 0;

		return true;
	}

    @Override
    public List<TstProject> genGroupVos(List<TstProject> pos) {
        List<TstProject> voList = new LinkedList<TstProject>();
//        for (TstProject po : pos) {
//            TstProject vo = genVo(po, null);
//            voList.add(vo);
//        }

        return voList;
    }

    @Override
    public List<TstProject> genVos(List<TstProject> pos, Map<String, Map<String, Boolean>> privMap) {
        List<TstProject> voList = new LinkedList<>();
        for (TstProject po : pos) {
            voList.add(po);
            List<TstProject> children = po.getChildren();
            boolean childCanView = false;
            for (TstProject child : children) {
                child = genVo(child, privMap);

                if (child.getPrivs() != null
                        && child.getPrivs().get("project-view") != null
                        && child.getPrivs().get("project-view") ) {
                    childCanView = true;
                }
                voList.add(child);
            }
            po.setChildrenNumb(po.getChildren().size());

            if (childCanView) {
                po.getPrivs().put("project-view", true);
            }
        }

        return voList;
    }
    @Override
    public TstProject genVo(TstProject po, Map<String, Map<String, Boolean>> privMap) {
        if (po == null) {
            return null;
        }

        if (privMap != null && privMap.get(po.getId().toString()) != null) {
            po.setPrivs(privMap.get(po.getId().toString()));
        }

        return po;
    }

}
