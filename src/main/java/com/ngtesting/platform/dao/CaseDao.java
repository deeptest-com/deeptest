package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCase;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseDao {
    void create(TstCase testCase);
    List<TstCase> query(@Param("projectId") Integer projectId);

    TstCase get(@Param("id") Integer id, @Param("prjId") Integer prjId);
    TstCase getDetail(@Param("id") Integer id, @Param("prjId") Integer prjId);

    void renameNew(TstCase testCasePo);
    void renameUpdate(TstCase testCasePo);

    void moveCopy(TstCase testCase);
    void moveUpdate(TstCase testCase);

    Integer update(
            @Param("obj") TstCase testCasePo,
            @Param("props") List<String> props,
            @Param("projectId") Integer projectId);

    Integer delete(@Param("pId") Integer pId,
                @Param("projectId") Integer projectId);


    Integer changeContentType(@Param("id") Integer id,
                           @Param("contentType") String contentType,
                           @Param("projectId") Integer projectId);

    Integer reviewResult(@Param("id") Integer id,
                      @Param("result") Boolean result,
                      @Param("projectId") Integer projectId);

    Integer updateProp(@Param("id") Integer id,
                    @Param("prop") String prop,
                    @Param("value") String value,
                    @Param("projectId") Integer projectId);

    void updateParentIfNeeded(@Param("pId") Integer pId);
    List<TstCase> getChildren(@Param("id") Integer id);

    Integer getChildMaxOrderNumb(@Param("pId") Integer pId);

    void addOrderForTargetAndNextCases(@Param("srcId") Integer srcId,
                                       @Param("targetOrdr") Integer targetOrdr,
                                       @Param("targetPid") Integer targetPid);

    void addOrderForNextCases(@Param("srcId") Integer srcId,
                              @Param("targetOrdr") Integer targetOrdr,
                              @Param("targetPid") Integer targetPid);
}
