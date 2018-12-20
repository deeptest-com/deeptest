package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface ReportIssueDao {
    List<TstMsg> query(@Param("userId") Integer userId);

    List<Map> chartIssueTrendCreate(@Param("_projectIds") String projectIds,
                                    @Param("_numb") Integer numb);
    List<Map> chartIssueTrendFinal(@Param("_projectIds") String projectIds,
                                   @Param("_numb") Integer numb);

    List<Map> chartIssueAge(@Param("_projectIds") String projectIds,
                            @Param("_numb") Integer numb);

    List<Map> chartIssueDistribByPriority(@Param("_projectId") Integer projectId,
                                          @Param("_projectType") String projectType);
    List<Map> chartIssueDistribByStatus(@Param("_projectId") Integer projectId,
                                        @Param("_projectType") String projectType);
}
