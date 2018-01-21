package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.ngtesting.platform.config.PropertyConfig;
import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;
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
    RelationOrgGroupUserService orgGroupUserService;
    @Autowired
    ProjectService projectService;

    @Autowired
    MailService mailService;

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
	public List<Map> getProjectUsers(Long orgId, Long projectId) {

		List<Map> users = new LinkedList<>();

        List<Object[]> ls = getDao().getListBySQL("{call get_project_users(?)}", projectId);
        for (Object[] arr : ls) {
            Map<String, Object> map = new HashMap();
            map.put("id", arr[0].toString());
            map.put("name", arr[1].toString());

            users.add(map);
        }

		return users;
	}

	@Override
	public List<TestUser> search(Long orgId, String keywords, JSONArray exceptIds) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);

		dc.createAlias("orgSet", "orgs");
		dc.add(Restrictions.eq("orgs.id", orgId));

		if (exceptIds != null && exceptIds.size() > 0) {
			List<Long> ids = new ArrayList();
			for (Object json : exceptIds.toArray()) {
				ids.add(Long.valueOf(json.toString()));
			}

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
        saveOrUpdate(po);
		
		TestOrg org = (TestOrg)get(TestOrg.class, orgId);
		if (!contains(org.getUserSet(), po.getId())) {
			org.getUserSet().add(po);
			saveOrUpdate(org);
		}

		return po;
	}

    @Override
    public TestUser invitePers(UserVo userVo, UserVo newUser, List<RelationOrgGroupUserVo> relations) {
	    Long orgId = userVo.getDefaultOrgId();
        Long prjId = userVo.getDefaultPrjId();
        String prjName = userVo.getDefaultPrjName();

        TestUser existUser  = accountService.getByEmail(newUser.getEmail());
        boolean isNew;

        TestUser po;
        if (existUser != null) {
            isNew = false;
            po = existUser;
        } else {
            isNew = true;
            po = new TestUser();
            po.setDefaultOrgId(orgId);

            po.setName(newUser.getName());
            po.setPhone(newUser.getPhone());
            po.setEmail(newUser.getEmail());
            po.setDisabled(newUser.getDisabled());
            po.setAvatar("upload/sample/user/avatar.png");

            saveOrUpdate(po);
        }

        TestOrg org = (TestOrg)get(TestOrg.class, orgId);
        if (!contains(org.getUserSet(), po.getId())) {
            org.getUserSet().add(po);
            saveOrUpdate(org);

            projectService.getHistoryPers(orgId, po.getId(), prjId, prjName);

            orgGroupUserService.saveRelations(relations);

            TestVerifyCode verifyCode = accountService.genVerifyCodePers(po.getId());
            String sys = PropertyConfig.getConfig("sys.name");
            Map<String, String> map = new HashMap<String, String>();
            map.put("user", po.getName() + "(" + po.getEmail() + ")");
            map.put("name", po.getName());
            map.put("sys", sys);

            String url;
            if (isNew) {
                url = PropertyConfig.getConfig("url.reset.password") + "/" + verifyCode.getCode();
            } else {
                url = PropertyConfig.getConfig("url.login");
            }
            map.put("url", url);
            mailService.sendTemplateMail("来自[" + sys + "]的邀请", "invite-user.ftl", newUser.getEmail(), map);

            return po;
        } else {
            return null;
        }
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
