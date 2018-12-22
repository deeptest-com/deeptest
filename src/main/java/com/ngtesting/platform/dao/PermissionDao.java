package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface PermissionDao {

    List<Map> listOrgPermission(@Param("userId") Integer userId);

    List<Map> listPrjPermission(@Param("userId") Integer userId);

}
