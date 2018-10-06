package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.dao.OrgDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.service.PushSettingsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.Map;

@Service(value = "orgService")
public class OrgServiceImpl extends BaseServiceImpl implements OrgService {
	@Autowired
    OrgPrivilegeService orgRolePrivilegeService;
    @Autowired
    PushSettingsService pushSettingsService;

    @Autowired
    private ProjectService projectService;
    @Autowired
    private ProjectDao projectDao;

	@Autowired
	private OrgDao orgDao;
    @Autowired
    private UserDao userDao;
    @Autowired
    private AuthDao authDao;

	@Override
	public List<TstOrg> list(Integer userId, String keywords, Boolean disabled) {
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
	public TstOrg save(TstOrg vo, TstUser user) {
		if (vo.getId() == null) {
            vo.setDeleted(false);
            orgDao.save(vo);

			orgDao.initOrg(vo.getId(), user.getId());
		} else {
            orgDao.update(vo);
        }

        pushSettingsService.pushMyOrgs(user);

		return vo;
	}

	@Override
	public Boolean delete(Integer id, TstUser user) {
        Integer currOrgId = user.getDefaultOrgId();

		Integer count = orgDao.delete(id);

		if (count > 0) {
            setUserDefaultOrgPrjToNullForDelete(id);

            if (currOrgId != null && id.intValue() == currOrgId.intValue()) { // 当前组织被删了
                changeDefaultOrg(user, null);
            }

            return true;
        }

        pushSettingsService.pushMyOrgs(user);

        return false;
	}

    @Override
    @Transactional
    public void changeDefaultOrg(TstUser user, Integer orgId) {
	    if (orgId == null) {
            user.setDefaultOrgId(null);
            user.setDefaultOrgName(null);

            pushSettingsService.pushOrgSettings(user);

            projectService.changeDefaultPrj(user, null);
            return;
        }

        TstOrg org = orgDao.get(orgId);
        orgDao.setDefault(user.getId(), orgId, org.getName());
        user.setDefaultOrgId(orgId);
        user.setDefaultOrgName(org.getName());

        projectService.changeToAnotherPrj(user);

        pushSettingsService.pushOrgSettings(user);
    }

    @Override
    @Transactional
    public void setUserDefaultOrgPrjToNullForDelete(Integer orgId) {
        orgDao.setDefaultOrgPrjToNullForDelete(orgId);
    }

	@Override
	public void genVos(List<TstOrg> pos, Integer userId) {
		TstUser user = userDao.get(userId);

		for (TstOrg po : pos) {
			if (user.getDefaultOrgId() != null
                    && user.getDefaultOrgId().longValue() == po.getId().longValue()) {
                po.setDefaultOrg(true);
			}
			Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userId, po.getId());
            po.setOrgPrivileges(orgPrivileges);
		}
	}

}
