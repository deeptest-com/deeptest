package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.SysPrivilege;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface SysPrivilegeDao {

    List<SysPrivilege> queryByUser(@Param("userId") Integer userId);

}
