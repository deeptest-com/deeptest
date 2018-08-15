package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface TestTaskDao {
    List<TstMsg> query(@Param("userId") Integer userId);
    List<TstTask> listByPlan(@Param("planId") Integer planId);

    TstTask get(@Param("id") Integer id);
    TstTask getDetail(@Param("id") Integer id,
                      @Param("projectId") Integer projectId);

    List<Integer> listCaseIds(@Param("id") Integer id);

    Integer save(TstTask vo);
    Integer update(TstTask vo);
    Integer delete(@Param("id") Integer id, @Param("projectId") Integer projectId);
    Integer close(@Param("id") Integer id, @Param("projectId") Integer projectId);

    void removeAssignees(@Param("id") Integer id);
    void saveAssignees(@Param("id") Integer id, @Param("list")  List<TstUser> assignees);

    List<Map> countStatus(@Param("id") Integer id);

    void updateCaseProject(@Param("id") Integer id, @Param("caseProjectId") Integer caseProjectId);

    void addCasesBySuites(@Param("taskId") Integer taskId, @Param("suiteIds") String suiteIds);

    void addCases(@Param("taskId") Integer taskId,
                  @Param("caseIds") String caseIds,
                  @Param("append") Boolean append);

    void start(@Param("id") Integer id);

    List<Integer> listAssigneeIds(@Param("id") Integer id);
}
