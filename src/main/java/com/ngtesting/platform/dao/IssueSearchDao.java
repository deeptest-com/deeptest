package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueSearchDao {
    List<Map> idSearch(@Param("text") String text,
                       @Param("exceptIds") List exceptIds,
                       @Param("projectId") Integer projectId);

    List<Map> titleSearch(@Param("text") String text,
                          @Param("exceptIds") List exceptIds,
                          @Param("projectId") Integer projectId);
}
