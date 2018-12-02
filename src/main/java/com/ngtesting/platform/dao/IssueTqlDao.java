package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuIssue;
import org.apache.ibatis.annotations.Param;

import java.util.LinkedHashMap;
import java.util.List;

public interface IssueTqlDao {
    List<IsuIssue> query(@Param("conditions") String conditions,
                         @Param("columns") String columns,
                         @Param("orderBy") LinkedHashMap<String, String> orderBy);
}
