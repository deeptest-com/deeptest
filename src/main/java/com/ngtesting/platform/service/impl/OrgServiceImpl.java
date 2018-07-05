package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service(value = "orgService")
public class OrgServiceImpl extends BaseServiceImpl implements OrgService {

	@Autowired
	ProjectRoleService projectRoleService;
	@Autowired
	ProjectService projectService;
    @Autowired
	OrgRoleService orgRoleService;
    @Autowired
    OrgGroupService orgGroupService;
	@Autowired
    CaseExeStatusService caseExeStatusService;
	@Autowired
    CasePriorityService casePriorityService;
	@Autowired
    CaseTypeService caseTypeService;
	@Autowired
	RelationProjectRoleEntityService relationProjectRoleEntityService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
	@Autowired
    OrgRolePrivilegeService orgRolePrivilegeService;

	@Override
	public List<TstOrg> list(String keywords, String disabled, Integer userId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstOrg.class);
//        dc.createAlias("userSet", "users");
//        dc.add(Restrictions.eq("users.id", userId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        if (StringUtil.isNotEmpty(keywords)) {
//			dc.add(Restrictions.like("name", "%" + keywords + "%"));
//		}
//        if (StringUtil.isNotEmpty(disabled)) {
//			dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
//		}
//
//        dc.addOrder(Order.asc("id"));
//        List<TstOrg> ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	@Override
	public List<TstOrg> listVo(String keywords, String disabled, Integer id) {
//		List ls = list(keywords, disabled, id);
//		List<TstOrg> vos = genVos(ls, id);
//		return vos;

		return null;
	}

	@Override
	public TstOrg getDetail(Integer id) {
//		if (id == null) {
//			return null;
//		}
//		TstOrg po = (TstOrg) get(TstOrg.class, id);
//
//		return po;

		return null;
	}

	@Override
	public TstOrg save(TstOrg vo, Integer userId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TestUser user = (TestUser)get(TestUser.class, userId);
//
//		boolean isNew = vo.getId() == null;
//		TstOrg po = new TstOrg();
//		if (!isNew) {
//			po = (TstOrg) get(TstOrg.class, vo.getId());
//		}
//
//		po.setName(vo.getName());
//		po.setWebsite(vo.getWebsite());
//		po.setDisabled(vo.getDisabled());
//
//		saveOrUpdate(po);
//
//        if (isNew) {
//            getDao().querySql("{call init_org(?,?)}", po.getId(), user.getId());
//        }
//
//		if (user.getDefaultOrgId() == null) {
//			user.setDefaultOrgId(po.getId());
//			saveOrUpdate(user);
//		}
//
//		return po;

		return null;
	}

	@Override
	public Boolean disable(Integer id) {
//		if (id == null) {
//			return false;
//		}
//
//		TstOrg po = (TstOrg) get(TstOrg.class, id);
//		po.setDisabled(true);
//		saveOrUpdate(po);

		return true;
	}

	@Override
	public Boolean delete(Integer id) {
//		if (id == null) {
//			return false;
//		}
//
//		TstOrg po = (TstOrg) get(TstOrg.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

	@Override
	public void setDefaultPers(Integer orgId, TstUser TstUser) {
//		TestUser user = (TestUser) get(TestUser.class, TstUser.getId());
//
//		user.setDefaultOrgId(orgId);
//
//		List<TstProjectAccessHistory> recentProjects = projectService.listRecentProjectVo(orgId, TstUser.getId());
//        user.setDefaultPrjId(recentProjects.size()>0?recentProjects.get(0).getProjectId(): null);
//        saveOrUpdate(user);
//
//		TstUser.setDefaultOrgId(user.getDefaultOrgId());
//		if (user.getDefaultOrgId()!=null) {
//			TstOrg org = (TstOrg)get(TstOrg.class, user.getDefaultOrgId());
//			TstUser.setDefaultOrgName(org.getName());
//		}
//
//        TstUser.setDefaultPrjId(recentProjects.size()>0?recentProjects.get(0).getProjectId(): null);
//		TstUser.setDefaultPrjName(recentProjects.size()>0?recentProjects.get(0).getProjectName(): "");
	}

	@Override
	public List<TstOrg> genVos(List<TstOrg> pos, Integer userId) {
//		TestUser user = (TestUser)get(TestUser.class, userId);
//
//		List<TstOrg> voList = new LinkedList<TstOrg>();
//		for (TstOrg po : pos) {
//			TstOrg vo = genVo(po);
//			if (po.getId().longValue() == user.getDefaultOrgId().longValue()) {
//				vo.setDefaultOrg(true);
//			}
//			Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userId, po.getId());
//			vo.setOrgPrivileges(orgPrivileges);
//			voList.add(vo);
//		}
//
//		return voList;

		return null;
	}

	@Override
	public TstOrg genVo(TstOrg po) {
//		if (po == null) {
//			return null;
//		}
//		TstOrg vo = new TstOrg();
//		BeanUtilEx.copyProperties(vo, po);
//
//		return vo;

		return null;
	}

}
