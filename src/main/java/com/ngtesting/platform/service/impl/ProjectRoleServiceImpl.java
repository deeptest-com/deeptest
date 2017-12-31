package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestProjectPrivilegeDefine;
import com.ngtesting.platform.entity.TestProjectRoleForOrg;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRolePriviledgeRelationService;
import com.ngtesting.platform.service.ProjectRoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.ProjectRoleVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class ProjectRoleServiceImpl extends BaseServiceImpl implements ProjectRoleService {
    @Autowired
    private ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    private ProjectRolePriviledgeRelationService projectRolePriviledgeRelationService;

	@Override
	public List list(Long orgId, String keywords, String disabled) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestProjectRoleForOrg.class);
        dc.add(Restrictions.eq("orgId", orgId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);

		List<ProjectRoleVo> vos = this.genVos(ls);
		
		return vos;
	}

	@Override
	public TestProjectRoleForOrg save(ProjectRoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		TestProjectRoleForOrg po = new TestProjectRoleForOrg();
		if (vo.getId() != null) {
			po = (TestProjectRoleForOrg) get(TestProjectRoleForOrg.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setOrgId(orgId);
		po.setDisabled(vo.getDisabled());
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		TestProjectRoleForOrg po = (TestProjectRoleForOrg) get(TestProjectRoleForOrg.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public ProjectRoleVo genVo(TestProjectRoleForOrg role) {
		ProjectRoleVo vo = new ProjectRoleVo();
		BeanUtilEx.copyProperties(vo, role);
		
		return vo;
	}

//	@Override
//	public TestProjectRoleForOrg createDefaultBasicDataPers(Long orgId) {
//        List<TestProjectPrivilegeDefine> allProjectPrivileges =
//                projectPrivilegeService.listAllProjectPrivileges();
//
//		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectRoleForOrg.class);
//		dc.add(Restrictions.eq("isBuildIn", true));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("id"));
//		List<TestProjectRoleForOrg> ls = findAllByCriteria(dc);
//
//		TestProjectRoleForOrg defaultRole = null;
//        String sql = "";
//		for (TestProjectRoleForOrg p : ls) {
//			TestProjectRoleForOrg temp = new TestProjectRoleForOrg();
//			BeanUtilEx.copyProperties(temp, p);
//			temp.setId(null);
//			temp.setOrgId(orgId);
//			temp.setBuildIn(false);
//			saveOrUpdate(temp);
//
//			if ("test_leader".equals(temp.getCode())) {
//				defaultRole = temp;
//                sql += projectRolePriviledgeRelationService.addPriviledgeForLeaderPers(allProjectPrivileges, temp.getId());
//			} else if ("test_designer".equals(temp.getCode())) {
//                sql += projectRolePriviledgeRelationService.addPriviledgeForDesignerPers(allProjectPrivileges, temp.getId());
//			} else if ("tester".equals(temp.getCode())) {
//                sql += projectRolePriviledgeRelationService.addPriviledgeForTesterPers(allProjectPrivileges, temp.getId());
//			}
//		}
//        getDao().querySql(sql);
//		return defaultRole;
//	}

	@Override
	public List<ProjectRoleVo> genVos(List<TestProjectRoleForOrg> pos) {
        List<ProjectRoleVo> vos = new LinkedList<ProjectRoleVo>();

        for (TestProjectRoleForOrg po: pos) {
        	ProjectRoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
