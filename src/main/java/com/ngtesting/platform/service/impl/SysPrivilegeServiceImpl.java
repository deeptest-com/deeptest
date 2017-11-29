package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.SysPrivilege;
import com.ngtesting.platform.service.SysPrivilegeService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.SysPrivilegeVo;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class SysPrivilegeServiceImpl extends BaseServiceImpl implements SysPrivilegeService {

    @Override
    public Map<String, Boolean> listByUser(Long userId) {

        String hql = "select priv from SysPrivilege priv" +
                " join priv.sysRoleSet roles " +
                " join roles.userSet users " +
                " where users.id = ?" +
                " and priv.deleted != true and priv.disabled!= true " +
                " order by users.id asc";

        List<SysPrivilege> ls = getDao().getListByHQL(hql, userId);

        Map<String, Boolean> map = new HashMap();
        for (SysPrivilege po: ls) {
            map.put(po.getCode().toString(), true);
        }
        
		return map;
	}

    @Override
    public List<SysPrivilegeVo> genVos(List<SysPrivilege> pos) {
        List<SysPrivilegeVo> vos = new LinkedList();

        for (SysPrivilege po: pos) {
            SysPrivilegeVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public SysPrivilegeVo genVo(SysPrivilege po) {
        if (po == null) {
            return null;
        }
        SysPrivilegeVo vo = new SysPrivilegeVo();
        BeanUtilEx.copyProperties(vo, po);

        return vo;
    }

}
