package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgPrivilegeDao;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.service.OrgPrivilegeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class OrgPrivilegeServiceImpl extends BaseServiceImpl implements OrgPrivilegeService {
    @Autowired
    private OrgPrivilegeDao orgPrivilegeDao;

    @Override
    public Map<String, Boolean> listByUser(Integer userId, Integer orgId) {
        Map<String, Boolean> map = new HashMap();
        if (orgId == null) {
            return map;
        }

        List<TstOrgPrivilegeDefine> ls = orgPrivilegeDao.listByUser(orgId, userId);
        for (TstOrgPrivilegeDefine priv: ls) {
            map.put(priv.getCode().toString(), true);
        }

        return map;
    }

    @Override
    public List<TstOrgPrivilegeDefine> listAllOrgPrivileges() {
        List<TstOrgPrivilegeDefine> ls = orgPrivilegeDao.listAllOrgPrivileges();
        return ls;
    }

}
