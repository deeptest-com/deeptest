package com.ngtesting.platform.service.impl;

import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

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
//	@Autowired
//	private UserDao userDao;
    @Autowired
    private CaseService caseService;

    @Autowired
    private UserService userService;
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
	public List<TstProject> listProjectGroups(Integer orgId) {
		List<TstProject> pos = projectDao.listProjectGroups(orgId);
		this.genGroupVos(pos);
		return pos;
	}

	@Override
	public List<TstProjectAccessHistory> listRecentProject(Integer orgId, Integer userId) {
        PageHelper.startPage(0, 5);
		List<TstProjectAccessHistory> pos = projectDao.listRecent(orgId, userId);

		return pos;
	}

	@Override
	public TstProject get(Integer id) {
		if (id == null) {
			return null;
		}
		TstProject po = projectDao.get(id);

		return po;
	}

	@Override
    @Transactional
	public TstProject save(TstProject vo, Integer orgId, TstUser TstUser) {
		if (vo == null) {
			return null;
		}

        vo.setOrgId(orgId);

        boolean disableStatusChanged = false;
		boolean isNew = vo.getId() == null;
		if (isNew) {
            projectDao.save(vo);
		} else {
            TstProject old = projectDao.get(vo.getId());
            disableStatusChanged = vo.getDisabled() != old.getDisabled();

            projectDao.update(vo);
		}

        if(isNew && TstProject.ProjectType.project.equals(vo.getType())) {
            projectPrivilegeService.addUserAsProjectTestLeaderPers(orgId, vo.getId(), "test_leader", TstUser.getId());
            caseService.createSample(vo.getId(), TstUser);
        }
        if(TstProject.ProjectType.project.equals(vo.getType())) {
            historyService.create(vo.getId(), TstUser,
                    isNew? Constant.MsgType.create.msg: Constant.MsgType.create.update.msg,
                    TstHistory.TargetType.project, vo.getId(), vo.getName());
        }

		if (!disableStatusChanged) {
			return vo;
		}

		// 项目被启用
		if (!vo.getDisabled()) {
			if (vo.getType().equals(TstProject.ProjectType.project)) {
				// 启用父
				projectDao.enable(vo.getParentId());
			} else {
				// 启用子
                projectDao.enableChildren(vo.getId());
			}
		}

		// 项目组被归档，归档子项目
		if (vo.getDisabled() && vo.getType().equals(TstProject.ProjectType.group)) {
            projectDao.disableChildren(vo.getId());
		}

		return vo;
	}

	@Override
	public Boolean delete(Integer id, Integer userId) {
        projectDao.delete(id, userId);

		return true;
	}

	@Override
    @Transactional
	public TstProject viewPers(Integer projectId, TstUser tstUser) {
		TstProject po = get(projectId);

        if (po.getType().equals(TstProject.ProjectType.project)) {
            projectDao.genHistory(po.getOrgId(), tstUser.getId(), projectId, po.getName());

            userService.setDefaultPrj(tstUser, projectId);
		}

		return po;
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
//			history = pos.getDetail(0);
//			history.setProjectName(projectName);
//		} else {
//			history = new TstProjectAccessHistory(orgId, userId, projectId, projectName);
//		}
//        history.setLastAccessTime(new Date());
//        saveOrUpdate(history);
	}

	@Override
	public boolean isLastestProjectGroup(Integer orgId, Integer projectGroupId) {
		Integer count = projectDao.isLastestProjectGroup(orgId, projectGroupId);
		return count > 0;
	}

    @Override
    public List<TstProject> genGroupVos(List<TstProject> pos) {
        for (TstProject po : pos) {
            genVo(po, null);
        }

        return pos;
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
