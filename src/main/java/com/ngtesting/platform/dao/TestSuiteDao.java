package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInSuite;
import com.ngtesting.platform.model.TstSuite;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestSuiteDao {
    List<TstSuite> query(@Param("projectId") Integer projectId,
                         @Param("keywords") String keywords,
                         @Param("disabled") Boolean disabled);

    List<TstSuite> listForImport(@Param("projectIds") List<Integer> projectIds);

    TstSuite get(@Param("id") Integer id,
                 @Param("projectId") Integer projectId);
    TstSuite getWithCases(@Param("id") Integer id);

    Integer save(TstSuite vo);
    Integer update(TstSuite vo);
    Integer delete(@Param("id") Integer id,
                @Param("projectId") Integer projectId);

    List<TstCaseInSuite> listCases(@Param("id") Integer id);

    List<Integer> listCaseIds(@Param("id") Integer id);

    void updateSuiteProject(@Param("id") Integer id,
                            @Param("projectId") Integer projectId,
                            @Param("caseProjectId") Integer caseProjectId,
                            @Param("userId") Integer userId);

    void addCases(@Param("suiteId") Integer suiteId,
                  @Param("caseIds") String caseIds);
}
