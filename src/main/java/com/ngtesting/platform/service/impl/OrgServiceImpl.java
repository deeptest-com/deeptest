package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

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
    OrgPrivilegeService orgRolePrivilegeService;

	@Autowired
	private OrgDao orgDao;
    @Autowired
    private UserDao userDao;

	@Override
	public List<TstOrg> list(Integer userId, String keywords, String disabled) {
        List<TstOrg> ls = orgDao.query(userId, keywords, disabled);

        genVos(ls, userId);

		return ls;
	}

	@Override
	public List<TstOrg> listByUser(Integer userId) {
        List<TstOrg> pos = orgDao.queryByUser(userId);
		genVos(pos, userId);
		return pos;
	}

	@Override
	public TstOrg get(Integer id) {
		TstOrg po = orgDao.get(id);
		return po;
	}

	@Override
    @Transactional
	public TstOrg save(TstOrg vo, Integer userId) {
		boolean isNew = vo.getId() == null;
		if (isNew) {
            vo.setDeleted(false);
            orgDao.save(vo);
		} else {
            orgDao.update(vo);
        }

        if (isNew) {
            orgDao.initOrg(vo.getId(), userId);
        }

		return vo;
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
	public void genVos(List<TstOrg> pos, Integer userId) {
		TstUser user = userDao.get(userId);

		for (TstOrg po : pos) {
			if (po.getId().longValue() == user.getDefaultOrgId().longValue()) {
                po.setDefaultOrg(true);
			}
			Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userId, po.getId());
            po.setOrgPrivileges(orgPrivileges);
		}
	}

}
