package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuType;
import org.apache.ibatis.annotations.Param;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public interface IssueDao {
    IsuIssue get(@Param("id") Integer id,
                 @Param("userId") Integer userId,
                 @Param("prjId") Integer prjId);

    IsuIssue getDetail(@Param("id") Integer id,
                       @Param("userId") Integer userId,
                       @Param("prjId") Integer prjId);

    IsuIssue getData(@Param("id") Integer id,
                     @Param("userId") Integer userId,
                     @Param("prjId") Integer prjId);

    IsuIssue getByUuid(@Param("uuid") String uuid);

    List<Integer> getByIds(@Param("ids") LinkedList<Integer> ids);

    Integer save(@Param("elems") List<IsuPageElement> elems,
                 @Param("params") List<Object> params);
    Integer saveExt(@Param("elems") List<IsuPageElement> elems,
                    @Param("params") List<Object> params,
                    @Param("id") Integer id);

    void setDefaultVal(@Param("model") IsuIssue model);

    Integer update(@Param("elems") List<IsuPageElement> elems,
                   @Param("params") List<Object> params,
                   @Param("id") Integer id,
                   @Param("orgId") Integer orgId);
    Integer updateExt(@Param("elems") List<IsuPageElement> elems,
                   @Param("params") List<Object> params,
                   @Param("id") Integer id);

    Integer updateProp(@Param("id") Integer id,
                       @Param("code") String code,
                       @Param("value") String value,
                       @Param("projectId") Integer projectId);
    Integer updatePropExt(@Param("id") Integer id,
                       @Param("code") String code,
                       @Param("value") String value);



    IsuType getProjectDefaultType(@Param("orgId") Integer orgId,
                                  @Param("prjId") Integer prjId);

    List<Map<String, Object>> getProjectDefaultPages(@Param("orgId") Integer orgId,
                                      @Param("prjId") Integer prjId,
                                      @Param("typeId") Integer typeId);

    Integer delete(@Param("id") Integer id,
                   @Param("projectId") Integer projectId);
}
