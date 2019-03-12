package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface PermissionDao {
    List<String> getShiroStylePermissions(@Param("userId") Integer userId);

//    List<Map> listOrgPermission(@Param("userId") Integer userId);
//    List<Map> listPrjPermission(@Param("userId") Integer userId);
}
