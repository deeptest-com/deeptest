package com.ngtesting.platform.service.impl;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;
import java.util.UUID;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysOrgPrivilege;
import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysCustomField.FieldApplyTo;
import com.ngtesting.platform.entity.SysCustomField.FieldFormat;
import com.ngtesting.platform.entity.SysCustomField.FieldType;
import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.service.TestProjectService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CasePriorityVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;

@Service
public class CustomFieldServiceImpl extends BaseServiceImpl implements CustomFieldService {

	@Autowired
	TestProjectService projectService;
	
	@Override
	public List<SysCustomField> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysCustomField.class);
        
        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("displayOrder"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}
	@Override
	public List<CustomFieldVo> listVos(Long orgId) {
        List<SysCustomField> ls = list(orgId);
        
        List<CustomFieldVo> vos = genVos(ls);
		return vos;
	}

	@Override
	public SysCustomField save(CustomFieldVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysCustomField po;
		if (vo.getId() != null) {
			po = (SysCustomField) get(SysCustomField.class, vo.getId());
		} else {
			po = new SysCustomField();
		}
		this.initPo(po, vo);
		
		po.setApplyTo(SysCustomField.FieldApplyTo.getValue(vo.getApplyTo()));
		po.setType(SysCustomField.FieldType.getValue(vo.getType()));
		if (StringUtil.isNotEmpty(vo.getFormat())) {
			po.setFormat(SysCustomField.FieldFormat.getValue(vo.getFormat()));
		}
		
		po.setOrgId(orgId);
		
		if (vo.getId() == null) {
			po.setCode(UUID.randomUUID().toString());
			
			String hql = "select max(displayOrder) from SysCustomField";
			Integer maxOrder = (Integer) getByHQL(hql);
	        po.setDisplayOrder(maxOrder + 10);
		}
		if (!po.getType().equals(FieldType.text)) {
			po.setRows(0);
			po.setFormat(null);
		}
		if (po.getIsGlobal() && po.getProjectSet().size() > 0) {
			po.setProjectSet(new HashSet<TestProject>(0));
		}
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		SysCustomField po = (SysCustomField) get(SysCustomField.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
	
	@Override
	public List<String> listApplyTo() {
		List<String> ls = new LinkedList<String>();
		for (FieldApplyTo item: SysCustomField.FieldApplyTo.values()) {
			ls.add(item.toString());
		}
		return ls;
	}
	@Override
	public List<String> listType() {
		List<String> ls = new LinkedList<String>();
		for (FieldType item: SysCustomField.FieldType.values()) {
			ls.add(item.toString());
		}
		return ls;
	}
	@Override
	public List<String> listFormat() {
		List<String> ls = new LinkedList<String>();
		for (FieldFormat item: SysCustomField.FieldFormat.values()) {
			ls.add(item.toString());
		}
		return ls;
	}
	
	@Override
	public boolean changeOrderPers(Long id, String act) {
		SysCustomField type = (SysCustomField) get(SysCustomField.class, id);
		
        String hql = "from SysCustomField tp where tp.deleted = false and tp.disabled = false ";
        if ("up".equals(act)) {
        	hql += "and tp.displayOrder < ? order by displayOrder desc";
        } else if ("down".equals(act)) {
        	hql += "and tp.displayOrder > ? order by displayOrder asc";
        } else {
        	return false;
        }
        
        SysCustomField neighbor = (SysCustomField) getDao().findFirstByHQL(hql, type.getDisplayOrder());
		
        Integer order = type.getDisplayOrder();
        type.setDisplayOrder(neighbor.getDisplayOrder());
        neighbor.setDisplayOrder(order);
        
        saveOrUpdate(type);
        saveOrUpdate(neighbor);
		
		return true;
	}
	
	@Override
	public List<TestProjectVo> listProjectsForField(Long orgId, Long fieldId) {
		List<TestProject> allProjects = projectService.list(orgId ,null, null);
        
		Set<TestProject> projectsForField;
        if (fieldId == null) {
        	projectsForField = new HashSet<TestProject>();
        } else {
        	SysCustomField field = (SysCustomField) get(SysCustomField.class, fieldId);
        	projectsForField = field.getProjectSet();
        }
        
        List<TestProjectVo> vos = new LinkedList<TestProjectVo>();
        for (TestProject po1 : allProjects) {
        	TestProjectVo vo = projectService.genVo(po1);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestProject po2 : projectsForField) {
        		if (po1.getId() == po2.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}
	
	@Override
	public boolean saveRelationsProjects(Long fieldId, List<TestProjectVo> projects) {
		if (projects == null) {
			return false;
		}
		
		SysCustomField field = (SysCustomField) get(SysCustomField.class, fieldId);
		Set<TestProject> projectSet = field.getProjectSet();
		
		for (Object obj: projects) {
			TestProjectVo vo = JSON.parseObject(JSON.toJSONString(obj), TestProjectVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				TestProject project = (TestProject) get(TestProject.class, vo.getId());
				
    			if (vo.getSelecting() && !projectSet.contains(project)) { // 勾选
    				projectSet.add(project);
    			} else if (project != null) { // 取消
    				projectSet.remove(project);
    			}
			}
		}
		saveOrUpdate(field);
		
		return true;
	}
    
	@Override
	public CustomFieldVo genVo(SysCustomField po) {
		if (po == null) {
			return null;
		}
		CustomFieldVo vo = new CustomFieldVo();
		BeanUtilEx.copyProperties(vo, po);
		
		return vo;
	}
	@Override
	public List<CustomFieldVo> genVos(List<SysCustomField> pos) {
        List<CustomFieldVo> vos = new LinkedList<CustomFieldVo>();

        for (SysCustomField po: pos) {
        	CustomFieldVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
	
	@Override
	public void initPo(SysCustomField po, CustomFieldVo vo) {
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setRows(vo.getRows());
		po.setIsGlobal(vo.getIsGlobal());
		po.setIsRequired(vo.getIsRequired());
	}

}