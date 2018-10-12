package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface ProjectDao {
    List<TstProject> query(@Param("orgId") Integer orgId,
                           @Param("keywordsParam") String keywordsParam,
                           @Param("disabledParam") Boolean disabledParam);

    List<Map<String, String>> queryIdAndName(@Param("orgId") Integer orgId);

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

    void setDefault(@Param("id") Integer id,
                       @Param("prjId") Integer prjId,
                       @Param("prjName") String prjName);
    void setUserDefaultPrjToNullForDelete(@Param("prjId") Integer prjId);

    void enable(@Param("id") Integer id);
    void enableChildren(@Param("id") Integer id);
    void disableChildren(@Param("id") Integer id);

    List<TstProject> listBrothers(@Param("id") Integer id);

    List<Integer> listBrotherIds(@Param("id") Integer projectId);
}
