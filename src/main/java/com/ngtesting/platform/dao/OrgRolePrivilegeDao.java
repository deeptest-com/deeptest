package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgRolePrivilegeDao {
    List<TstOrgPrivilegeDefine> queryByUser(@Param("userId") Integer userId, @Param("orgId") Integer orgId);
}
