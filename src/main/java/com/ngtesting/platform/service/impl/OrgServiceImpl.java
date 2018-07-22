package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;

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
    ProjectRoleEntityRelationService relationProjectRoleEntityService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
	@Autowired
    OrgRolePrivilegeService orgRolePrivilegeService;

	@Autowired
	private OrgDao orgDao;
    @Autowired
    private UserDao userDao;

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
	public List<TstOrg> listByUser(Integer userId) {
        List<TstOrg> pos = orgDao.queryByUser(userId);
		List<TstOrg> vos = genVos(pos, userId);
		return vos;
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
	public List<TstOrg> genVos(List<TstOrg> pos, Integer userId) {
		TstUser user = userDao.get(userId);

		for (TstOrg po : pos) {
			if (po.getId().longValue() == user.getDefaultOrgId().longValue()) {
                po.setDefaultOrg(true);
			}
			Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userId, po.getId());
            po.setOrgPrivileges(orgPrivileges);
		}

		return pos;
	}

}
