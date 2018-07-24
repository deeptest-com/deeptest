package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.ProjectRoleDao;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRolePriviledgeRelationService;
import com.ngtesting.platform.service.ProjectRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProjectRoleServiceImpl extends BaseServiceImpl implements ProjectRoleService {
	@Autowired
	private ProjectRoleDao projectRoleDao;

    @Autowired
    private ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    private ProjectRolePriviledgeRelationService projectRolePriviledgeRelationService;

	@Override
	public List list(Integer orgId, String keywords, String disabled) {
		List<TstProjectRole> ls = projectRoleDao.list(orgId, keywords, disabled);

//        DetachedCriteria dc = DetachedCriteria.forClass(TestProjectRoleForOrg.class);
//        dc.add(Restrictions.eq("orgId", orgId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        if (StringUtil.isNotEmpty(keywords)) {
//        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//        if (StringUtil.isNotEmpty(disabled)) {
//        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
//        }
//
//        dc.addOrder(Order.asc("id"));
//        List ls = findAllByCriteria(dc);
//
//		List<TstProjectRole> vos = this.genVos(ls);
//
//		return vos;

		return ls;
	}

	@Override
	public TstProjectRole get(Integer roleId) {
		return null;
	}

	@Override
	public TstProjectRole save(TstProjectRole vo, Integer orgId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TestProjectRoleForOrg po = new TestProjectRoleForOrg();
//		if (vo.getId() != null) {
//			po = (TestProjectRoleForOrg) get(TestProjectRoleForOrg.class, vo.getId());
//		}
//
//		po.setName(vo.getName());
//		po.setDescr(vo.getDescr());
//		po.setOrgId(orgId);
//		po.setDisabled(vo.getDisabled());
//
//		saveOrUpdate(po);
//		return po;

		return null;
	}

	@Override
	public boolean delete(Integer id) {
//		TestProjectRoleForOrg po = (TestProjectRoleForOrg) get(TestProjectRoleForOrg.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

}
