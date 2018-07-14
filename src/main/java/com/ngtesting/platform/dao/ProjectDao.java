package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface ProjectDao {
    List<TstProject> query(@Param("orgId") Integer orgId, @Param("keywords") String keywords,
                           @Param("disabled") Boolean disabled);

    TstProject getDetail(@Param("id") Integer id);
    List<Map<String, String>> getProjectPrivilegeByOrgForUser(@Param("userId") Integer userId, @Param("orgId") Integer orgId);

    void genHistory(@Param("orgId") Integer orgId, @Param("userId") Integer userId,
                    @Param("prjId") Integer prjId, @Param("prjName") String prjName);

    List<TstProjectAccessHistory> listRecent(@Param("orgId") Integer orgId, @Param("userId") Integer userId);
}
