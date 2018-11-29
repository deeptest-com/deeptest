package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuType;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueDao {
    IsuIssue get(@Param("id") Integer id, @Param("orgId") Integer orgId);

    IsuType getProjectDefaultType(@Param("orgId") Integer orgId,
                                  @Param("prjId") Integer prjId);

    List<Map<String, Object>> getProjectDefaultPages(@Param("orgId") Integer orgId,
                                      @Param("prjId") Integer prjId,
                                      @Param("typeId") Integer typeId);

    IsuIssue save(@Param("issue") IsuIssue issue,
                  @Param("elems") List<IsuPageElement> elems);
    Integer update(@Param("issue") IsuIssue issue,
                   @Param("elems") List<IsuPageElement> elems);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);
}
