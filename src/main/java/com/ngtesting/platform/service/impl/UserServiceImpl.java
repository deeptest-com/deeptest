package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestRelationProjectRoleEntity;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleEntityService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;
import com.ngtesting.platform.vo.RelationProjectRoleEntityVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {
	
	@Autowired
	AccountService accountService;
	@Autowired
    RelationProjectRoleEntityService relationProjectRoleUserService;
    @Autowired
    RelationOrgGroupUserService relationOrgGroupUserService;
	@Autowired
	RelationProjectRoleEntityService relationProjectRoleEntityService;

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
        
        dc.createAlias("orgSet", "companies");
        dc.add(Restrictions.eq("companies.id", orgId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.or(Restrictions.like("name", "%" + keywords + "%"),
        		   Restrictions.like("email", "%" + keywords + "%"),
        		   Restrictions.like("phone", "%" + keywords + "%") ));
        }
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }
        
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
		
		return page;
	}

	@Override
	public List<RelationProjectRoleEntityVo> getProjectUsers(Long orgId, Long projectId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleEntity.class);

		dc.add(Restrictions.eq("projectId", projectId));

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.addOrder(Order.asc("id"));
		List<TestRelationProjectRoleEntity> ls = findAllByCriteria(dc);

        List<RelationProjectRoleEntityVo> vos = new LinkedList();

        List<Long> userIds = new LinkedList<>();
        for(TestRelationProjectRoleEntity r : ls) {
            if (r.getType().equals(TestRelationProjectRoleEntity.EntityType.user)) {
                RelationProjectRoleEntityVo vo = new RelationProjectRoleEntityVo();
                BeanUtilEx.copyProperties(vo, r);
                vos.add(vo);

                userIds.add(r.getEntityId());
            }
        }

		for(TestRelationProjectRoleEntity r : ls) {
		    if (r.getType().equals(TestRelationProjectRoleEntity.EntityType.group)) {
                Long groupId = r.getEntityId();
                List<RelationOrgGroupUserVo> rUsers  = relationOrgGroupUserService.listRelationsByGroup(orgId, groupId);

                for(RelationOrgGroupUserVo ru : rUsers) {
                    if (userIds.contains(ru.getUserId())) {
                        continue;
                    }
                    RelationProjectRoleEntityVo vo = new RelationProjectRoleEntityVo();
                    vo.setProjectId(r.getProjectId());
                    vo.setProjectRoleId(r.getProjectRoleId());
                    vo.setProjectRoleName(relationProjectRoleEntityService.getEntityName(r));
                    vo.setEntityId(ru.getUserId());
                    vo.setEntityName(ru.getUserName());
                    vos.add(vo);
                }
            }
        }

        Collections.sort(vos, new Comparator<RelationProjectRoleEntityVo>(){
            public int compare(RelationProjectRoleEntityVo o1, RelationProjectRoleEntityVo o2) {
                return o1.getEntityName().compareTo(o2.getEntityName());
            }
        });

		return vos;
	}

	@Override
	public List<TestUser> search(Long orgId, String keywords, JSONArray exceptIds) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);

		dc.createAlias("orgSet", "orgs");
		dc.add(Restrictions.eq("orgs.id", orgId));

		List<Long> ids = new ArrayList();
		for (Object json : exceptIds.toArray()) {
            ids.add(Long.valueOf(json.toString()));
        }

		if (exceptIds.size() > 0) {
            dc.add(Restrictions.not(Restrictions.in("id", ids)));
        }

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		if (StringUtil.isNotEmpty(keywords)) {
			dc.add(Restrictions.or(Restrictions.like("name", "%" + keywords + "%"),
					Restrictions.like("email", "%" + keywords + "%"),
					Restrictions.like("phone", "%" + keywords + "%") ));
		}

		dc.addOrder(Order.asc("id"));
		Page page = findPage(dc, 0, 20);

		return page.getItems();
	}

	@Override
	public TestUser save(UserVo userVo, Long orgId) {
		if (userVo == null) {
			return null;
		}
		
		TestUser temp = accountService.getByEmail(userVo.getEmail());
		if (temp != null && temp.getId() != userVo.getId()) {
			return null;
		}
		
		TestUser po;
		if (userVo.getId() != null) {
			po = (TestUser) get(TestUser.class, userVo.getId());
		} else {
			po = new TestUser();
			po.setDefaultOrgId(orgId);
		}
		
		po.setName(userVo.getName());
		po.setPhone(userVo.getPhone());
		po.setEmail(userVo.getEmail());
		po.setDisabled(userVo.getDisabled());
		if (userVo.getAvatar() == null) {
			po.setAvatar("upload/sample/user/avatar.png");
		}
		
		TestOrg org = (TestOrg)get(TestOrg.class, orgId);
		if (!contains(org.getUserSet(), userVo.getId())) {
			org.getUserSet().add(po);
			saveOrUpdate(org);
		}
		
		saveOrUpdate(po);
		return po;
	}
	
	@Override
	public boolean disable(Long userId, Long orgId) {
		TestUser po = (TestUser) get(TestUser.class, userId);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean remove(Long userId, Long orgId) {
		TestUser po = (TestUser) get(TestUser.class, userId);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	public boolean setSizePers(Long userId, Integer left, Integer right) {
        TestUser po = (TestUser) get(TestUser.class, userId);
        po.setCaseBoardLeftSize(left);
        po.setCaseBoardRightSize(right);
        saveOrUpdate(po);
		return true;
	}
    
	@Override
	public UserVo genVo(TestUser user) {
		if (user == null) {
			return null;
		}
		UserVo vo = new UserVo();
		BeanUtilEx.copyProperties(vo, user);
		
		return vo;
	}
	@Override
	public List<UserVo> genVos(List<TestUser> pos) {
        List<UserVo> vos = new LinkedList<UserVo>();

        for (TestUser po: pos) {
        	UserVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}
