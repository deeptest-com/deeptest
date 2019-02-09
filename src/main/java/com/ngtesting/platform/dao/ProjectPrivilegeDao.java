package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface ProjectPrivilegeDao {
    List<TstProjectPrivilegeDefine> listAllProjectPrivileges();

    List<Map<String, String>> listForUser(@Param("userId") Integer userId,
                                          @Param("prjId") Integer prjId,
                                          @Param("prjType") String prjType);
}
