package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface ProjectDao {
    List<TstProject> query(@Param("orgId") Integer orgId,
                           @Param("keywordsParam") String keywordsParam,
                           @Param("disabledParam") Boolean disabledParam);

    TstProject get(@Param("id") Integer id);
    void delete(@Param("id") Integer id, @Param("userId") Integer userId);

    void genHistory(@Param("orgId") Integer orgId, @Param("userId") Integer userId,
                    @Param("prjId") Integer prjId, @Param("prjName") String prjName);

    List<TstProjectAccessHistory> listRecent(@Param("orgId") Integer orgId, @Param("userId") Integer userId);

    Integer isLastestProjectGroup(@Param("orgId") Integer orgId, @Param("projectGroupId") Integer projectGroupId);

    List<TstProject> getProjectsByOrg(@Param("orgId") Integer orgId);
    List<TstProject> listProjectGroups(@Param("orgId") Integer orgId);

    void save(TstProject vo);
    void update(TstProject vo);

    void enable(@Param("id") Integer id);
    void enableChildren(@Param("id") Integer id);
    void disableChildren(@Param("id") Integer id);

    List<TstProject> listBrothers(@Param("id") Integer id);

    List<Integer> listBrotherIds(@Param("id") Integer projectId);
}
