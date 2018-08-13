package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCase;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseDao {
    void create(TstCase testCase);
    List<TstCase> query(@Param("projectId") Integer projectId);

    TstCase get(@Param("id") Integer id);
    TstCase getDetail(@Param("id") Integer id);

    void renameNew(TstCase testCasePo);
    void renameUpdate(TstCase testCasePo);

    void moveCopy(TstCase testCase);
    void moveUpdate(TstCase testCase);

    void update(
            @Param("obj") TstCase testCasePo,
            @Param("props") List<String> props);

    void delete(@Param("pId") Integer id);

    Integer getChildMaxOrderNumb(@Param("pId") Integer pId);

    void addOrderForTargetAndNextCases(@Param("srcId") Integer srcId,
                                       @Param("targetOrdr") Integer targetOrdr,
                                       @Param("targetPid") Integer targetPid);

    void addOrderForNextCases(@Param("srcId") Integer srcId,
                              @Param("targetOrdr") Integer targetOrdr,
                              @Param("targetPid") Integer targetPid);

    void updateProp(@Param("id") Integer id,
                    @Param("prop") String prop,
                    @Param("value") String value);

    List<TstCase> getChildren(@Param("id") Integer id);

    void changeContentTypePers(@Param("id") Integer id,
                               @Param("contentType") String contentType);

    void reviewResult(@Param("id") Integer id,
                      @Param("result") Boolean result);

    void updateParentIfNeeded(@Param("pId") Integer pId);
}
