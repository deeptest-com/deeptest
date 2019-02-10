package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgPrivilegeDao {
    List<TstOrgPrivilegeDefine> listAllOrgPrivileges();
    List<TstOrgPrivilegeDefine> listByUser(@Param("userId") Integer userId,
                                           @Param("orgId") Integer orgId);
}
