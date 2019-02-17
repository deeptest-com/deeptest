package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuIssue;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueTqlDao {
    List<IsuIssue> query(@Param("conditions") String conditions,
                         @Param("columns") String columns,
                         @Param("orders") List<Map<String, String>> orderBy);
}
