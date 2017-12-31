package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
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

	@Override
	public List<TestOrg> list(String keywords, String disabled, Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestOrg.class);
        dc.createAlias("userSet", "users");
        dc.add(Restrictions.eq("users.id", userId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        if (StringUtil.isNotEmpty(keywords)) {
			dc.add(Restrictions.like("name", "%" + keywords + "%"));
		}
        if (StringUtil.isNotEmpty(disabled)) {
			dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
		}

        dc.addOrder(Order.asc("id"));
        List<TestOrg> ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public List<OrgVo> listVo(String keywords, String disabled, Long id) {
		List ls = list(keywords, disabled, id);
		List<OrgVo> vos = genVos(ls, id);
		return vos;
	}

	@Override
	public TestOrg getDetail(Long id) {
		if (id == null) {
			return null;
		}
		TestOrg po = (TestOrg) get(TestOrg.class, id);

		return po;
	}

    @Override
    public void createDefaultBasicDataPers(TestUser user) {

        getDao().querySql("{call init_user(?)", user.getId());

//        TestOrg po = new TestOrg();
//
////      po.setAdminId(user.getId());
//        po.getUserSet().add(user);
//
//        po.setName("我的组织");
//        po.setWebsite("");
//
//        saveOrUpdate(po);
//
//        user.setDefaultOrgId(po.getId());
//        saveOrUpdate(user);
//
//        orgRoleService.initOrgRoleBasicDataPers(po.getId());
//        orgRoleService.addUserToOrgRolePers(user, po.getId(), TestOrgRole.OrgRoleCode.org_admin);
//        orgGroupService.initDefaultBasicDataPers(po);
//
//		caseExeStatusService.createDefaultBasicDataPers(po.getId());
//		casePriorityService.createDefaultBasicDataPers(po.getId());
//		caseTypeService.createDefaultBasicDataPers(po.getId());
//
//        TestProjectRoleForOrg defaultRole = projectRoleService.createDefaultBasicDataPers(po.getId());
//        Long projectId = projectService.initDefaultBasicDataPers(po.getId(), user.getId());
//		relationProjectRoleEntityService.addUserToProjectAsLeaderPers(user.getId(), defaultRole.getId(), projectId);

    }

	@Override
	public TestOrg save(OrgVo vo, Long userId) {
		if (vo == null) {
			return null;
		}

		TestUser user = (TestUser)get(TestUser.class, userId);

		boolean isNew = vo.getId() == null;
		TestOrg po = new TestOrg();
		if (!isNew) {
			po = (TestOrg) get(TestOrg.class, vo.getId());
		} else {
//			po.setAdminId(userId);
			po.getUserSet().add(user);
		}

		po.setName(vo.getName());
		po.setWebsite(vo.getWebsite());
		po.setDisabled(vo.getDisabled());

		saveOrUpdate(po);

        if (isNew) {
//            orgRoleService.initOrgRoleBasicDataPers(po.getId());
//            orgRoleService.addUserToOrgRolePers(user, po.getId(), TestOrgRole.OrgRoleCode.org_admin);
//            orgGroupService.initDefaultBasicDataPers(po);
//
//            caseExeStatusService.createDefaultBasicDataPers(po.getId());
//            casePriorityService.createDefaultBasicDataPers(po.getId());
//            caseTypeService.createDefaultBasicDataPers(po.getId());
//
//            projectService.initDefaultBasicDataPers(po.getId(), userId);
        }

		if (user.getDefaultOrgId() == null) {
			user.setDefaultOrgId(po.getId());
			saveOrUpdate(user);
		}

		return po;
	}

	@Override
	public Boolean disable(Long id) {
		if (id == null) {
			return false;
		}

		TestOrg po = (TestOrg) get(TestOrg.class, id);
		po.setDisabled(true);
		saveOrUpdate(po);

		return true;
	}

	@Override
	public Boolean delete(Long id) {
		if (id == null) {
			return false;
		}

		TestOrg po = (TestOrg) get(TestOrg.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);

		return true;
	}

	@Override
	public List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, UserVo userVo) {
		TestUser user = (TestUser) get(TestUser.class, userVo.getId());

		user.setDefaultOrgId(orgId);

		List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(orgId, userVo.getId());

		saveOrUpdate(user);

		userVo.setDefaultOrgId(user.getDefaultOrgId());

		return recentProjects;
	}

	@Override
	public List<OrgVo> genVos(List<TestOrg> pos, Long userId) {
		TestUser user = (TestUser)get(TestUser.class, userId);

		List<OrgVo> voList = new LinkedList<OrgVo>();
		for (TestOrg po : pos) {
			OrgVo vo = genVo(po);
			if (po.getId() == user.getDefaultOrgId()) {
				vo.setDefaultOrg(true);
			}

			voList.add(vo);
		}

		return voList;
	}

	@Override
	public OrgVo genVo(TestOrg po) {
		if (po == null) {
			return null;
		}
		OrgVo vo = new OrgVo();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

}
