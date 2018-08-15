package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.dao.OrgDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.OrgService;
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

		return vo;
	}

	@Override
	public Boolean delete(Integer id) {
		Integer count = orgDao.delete(id);

		return count > 0;
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
