package com.ngtesting.platform.dao;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuType;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueDao {
    IsuIssue get(@Param("id") Integer id, @Param("orgId") Integer orgId);
    IsuIssue getByUuid(@Param("uuid") String uuid,
                       @Param("orgId") Integer orgId);

    Integer save(@Param("elems") List<IsuPageElement> elems,
                 @Param("params") List<Object> params);

    Integer update(@Param("issue") JSONObject issue,
                   @Param("elems") List<IsuPageElement> elems);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    IsuType getProjectDefaultType(@Param("orgId") Integer orgId,
                                  @Param("prjId") Integer prjId);

    List<Map<String, Object>> getProjectDefaultPages(@Param("orgId") Integer orgId,
                                      @Param("prjId") Integer prjId,
                                      @Param("typeId") Integer typeId);
}
